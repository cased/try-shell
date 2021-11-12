# Cased Shell Kubernetes Demo

> If you have any questions, please email us at support@cased.com!

### Prerequisites

- `kubectl` access to a Kubernetes cluster that can either provide public access to a service (via a Service Load Balancer or Ingress Controller) or private access to a Service with self-managed TLS.
- Access to the settings tab of Cased Shell instance with a 'hostname' of `localhost:8888`. (e.g. https://app.cased.com/shell/programs/shell_EXAMPLE/settings) with Certificate Authentication enabled.

## Summary

The example in this directory deploys Cased Shell and an internally-accessible SSH server to your Kubernetes cluster. The container running the SSH server is granted admin access to the Kubernetes cluster, and is configured to allow access from any connection using the Certificate Authority member of your SSO organization. Access to the demo install of Cased Shell is facilitated using `kubectl port-forward`.

## Deploying

### Clone this repo

```
git clone https://github.com/cased/try-shell
cd try-shell
```

### Create a namespace

```
kubectl create ns shell
```

### Update configuration

Obtain the value of `CASED_SHELL_SECRET` Shell Instance's settings page (e.g. https://app.cased.com/shell/programs/shell_EXAMPLE/settings) and set it in `.env`.

```
vi .env
```

Obtain the `~/.ssh/authorized_keys` entry referenced on your Shell Instance's settings page. The value should look something like:

```
cert-authority,principals="noreply+org_EXAMPLE@cased.com" ecdsa-sha2-nistp256 AAAAEXAMPLE=
```

Set this as the value of the `PUBLIC_KEY` in the example `kustomization.yaml`:

```
vi examples/kubernetes/kustomization.yaml
```

### Preview configuration

```
kubectl kustomize examples/kubernetes | less
```

### Apply configuration

```
kubectl -n shell apply -k examples/kubernetes
```

### Inspect status

TODO

## Port forward

TODO

### Visit your shell

Next, visit your shell instance at http://localhost:8888. After logging in, click the `kubectl` prompt to run interactive commands using `kubectl`, or use the `event log` link to see a stream of all Kubernetes events. Edit `jump.yaml` to add additional entries, making sure to run `kubectl apply` again to apply the change.
