A container with an SSH server and a kubectl client, designed to be used as a 'jump host' for kubectl actions alongside Cased Shell. 

# Configuration

Expects a `kubectl-sshd` configmap with `KUBECTL_VERSION` and `PUBLIC_KEY` values.