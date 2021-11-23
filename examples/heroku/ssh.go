package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"time"
	"unsafe"

	"github.com/creack/pty"
	"github.com/gliderlabs/ssh"

	gossh "golang.org/x/crypto/ssh"
)

var (
	shellUrl = os.Args[1]
)

func setWinsize(f *os.File, w, h int) {
	syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), uintptr(syscall.TIOCSWINSZ),
		uintptr(unsafe.Pointer(&struct{ h, w, x, y uint16 }{uint16(h), uint16(w), 0, 0})))
}

func casedShellIsUserAuthority(providedPubKey gossh.PublicKey) bool {
	resp, err := http.Get(fmt.Sprintf("%s/ca.pub", shellUrl))
	if err != nil || resp.StatusCode != 200 {
		log.Println("error contacting shell")
		return false
	}
	authorizedKeyString, _ := io.ReadAll(resp.Body)
	authorizedKey, _, _, _, err := gossh.ParseAuthorizedKey([]byte(authorizedKeyString))
	if err != nil {
		log.Println(fmt.Sprintf("Failed to parse ca.pub: %s %v", authorizedKeyString, err))
		return false
	}
	return bytes.Equal(providedPubKey.Marshal(), authorizedKey.Marshal())
}

func casedShellPublicKeyHandler(ctx ssh.Context, pubKey ssh.PublicKey) bool {
	cert, ok := pubKey.(*gossh.Certificate)
	if !ok {
		log.Println("normal key pairs not accepted")
		return false
	}
	if cert.CertType != gossh.UserCert {
		log.Printf("cert has type %d\n", cert.CertType)
		return false
	}
	c := &gossh.CertChecker{
		IsUserAuthority: casedShellIsUserAuthority,
	}

	if !c.IsUserAuthority(cert.SignatureKey) {
		log.Println("certificate signed by unrecognized authority")
		return false
	}

	resp, err := http.Get(fmt.Sprintf("%s/principal.txt", shellUrl))
	if err != nil || resp.StatusCode != 200 {
		log.Println("error contacting shell")
		return false
	}
	principal, _ := io.ReadAll(resp.Body)

	if err := c.CheckCert(string(principal), cert); err != nil {
		log.Printf("%s not in list of valid principals %v\n", string(principal), cert.ValidPrincipals)
		return false
	}

	if cert.Permissions.CriticalOptions["force-command"] != "" {
		log.Printf("invalid force-command: %s\n", cert.Permissions.CriticalOptions["force-command"])
		return false
	}

	log.Printf("accepted SSH Certificate from %s\n", cert.ValidPrincipals[0])

	return true
}

func casedShellSessionHandler(tokens *sessionOAuthTokens, command []string) ssh.Handler {
	return func(s ssh.Session) {
		log.Printf("accepted connection for user %s\n", s.User())
		if len(s.Command()) > 0 {
			log.Println("command execution not supported")
			io.WriteString(s, "command execution not supported\n")
			s.Exit(1)
			return
		}

		log.Println("starting pty")
		ptyReq, winCh, isPty := s.Pty()

		if isPty {
			cert, ok := s.PublicKey().(*gossh.Certificate)
			if !ok {
				io.WriteString(s, "error parsing SSH Certificate\n")
				s.Exit(1)
				return
			}
			io.WriteString(s, fmt.Sprintf("Login to Heroku: https://%s/oauth/auth?stateToken=%s\n", os.Getenv("CASED_SHELL_HOSTNAME"), cert.KeyId))
			io.WriteString(s, "\nWaiting for token...")
			var token string
			for range time.Tick(time.Second * 1) {
				if tokens.Get(cert.KeyId) == "" || tokens.Get(cert.KeyId) == "pending" {
					io.WriteString(s, ".")
					continue
				} else {
					token = tokens.Get(cert.KeyId)
					io.WriteString(s, "done!\n")
					break
				}
			}
			cmd := exec.Command(command[0], command[1:]...)
			cmd.Env = append(cmd.Env, fmt.Sprintf("TERM=%s", ptyReq.Term))
			cmd.Env = append(cmd.Env, "SHELL=/bin/bash")
			cmd.Env = append(cmd.Env, fmt.Sprintf("HOME=%s", os.Getenv("HOME")))
			cmd.Env = append(cmd.Env, fmt.Sprintf("HEROKU_API_KEY=%s", token))
			f, err := pty.Start(cmd)
			if err != nil {
				io.WriteString(s, "error starting PTY\n")
				s.Exit(1)
				return
			}
			go func() {
				for win := range winCh {
					setWinsize(f, win.Width, win.Height)
				}
			}()
			go func() {
				io.Copy(f, s)
			}()
			io.Copy(s, f)
			cmd.Wait()
		} else {
			io.WriteString(s, "No PTY requested.\n")
			s.Exit(1)
		}
	}
}
