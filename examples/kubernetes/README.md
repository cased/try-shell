# Running Cased Shell on Kubernetes

> If you have any questions, please email us at support@cased.com!
### Prerequisites

- A Kubernetes cluster that can either provide public access to a service (via a Service Load Balancer or Ingress Controller) or private access to a Service with self-managed TLS.
- Access to the settings tab of Cased Shell instance created using a hostname that can access the Kubernetes cluster as mentioned above (e.g. https://app.cased.com/shell/programs/shell_EXAMPLE/settings), with Certificate Authentication enabled.
- Local `kubectl` access to the cluster.

## Summary

The example in this directory deploys Cased Shell and an internally-accessible SSH server to your Kubernetes cluster. The container running the SSH server is granted admin access to the Kubernetes cluster, and is configured to allow access from any member of your SSO organization.
## Deploying

### Create a namespace

```
kubectl create ns shell
```

### Update configuration

```
vi .env
```

Update the Kustomize file, paying special attention to the ingress strategy your cluster uses. You will need to either edit `ingresses/cased-shell.yaml` or [patch](https://kubernetes.io/docs/tasks/manage-kubernetes-objects/kustomization/#customizing) the included Service to change its type to LoadBalancer to allow HTTPS traffic to reach Cased Shell using `$CASED_SHELL_HOSTNAME`.

```
vi ingresses/cased-shell.yaml
vi kustomization.yaml
```

### Preview configuration

```
kubectl kustomize examples/kubernetes | less
```

### Apply configuration


```
kubectl -n shell apply -k examples/kubernetes
```

### Visit your shell

Next, visit your shell instance at https://$CASED_SHELL_HOSTNAME. After logging in, click the `kubectl` prompt to run interactive commands using `kubectl`, or use the `event log` link to see a stream of all Kubernetes events.
