package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gliderlabs/ssh"
)

type sessionOAuthTokens struct {
	Tokens map[string]string
}

func (s *sessionOAuthTokens) Get(sessionID string) string {
	return s.Tokens[sessionID]
}

func (s *sessionOAuthTokens) Set(sessionID, value string) {
	s.Tokens[sessionID] = value
}

func main() {
	tokens := &sessionOAuthTokens{
		Tokens: make(map[string]string),
	}
	http.HandleFunc("/oauth/auth", handleAuth(tokens))
	http.HandleFunc("/oauth/auth/callback", handleAuthCallback(tokens))
	http.HandleFunc("/oauth/user", handleUser)
	go http.ListenAndServe(":2225", nil)

	s := &ssh.Server{
		Addr:             ":2224",
		PublicKeyHandler: casedShellPublicKeyHandler,
		IdleTimeout:      60 * time.Second,
		Version:          "Cased Shell SSH",
	}

	s.Handle(casedShellSessionHandler(tokens, []string{"bash", "-i"}))
	log.Println("starting ssh server on port 2224...")
	log.Fatal(s.ListenAndServe())

}
