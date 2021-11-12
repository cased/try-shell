# Cased Shell Kubernetes Demo

> If you have any questions, please email us at support@cased.com!

### Prerequisites

- `kubectl` access to the Kubernetes cluster you'd like to deploy Cased Shell to.
- Access to the settings tab of Cased Shell instance with a 'hostname' of `localhost:NNNN` where `NNNN` is a free local port like `8899` (e.g. https://app.cased.com/shell/programs/shell_EXAMPLE/settings).

## Summary

The example in this directory deploys Cased Shell and an internally-accessible SSH server to your Kubernetes cluster. Access to the internal SSH service requires a SSH certificate signed by your Cased Organzation (which Cased Shell facilitates). Once authenticated, users have access to `kubectl` with admin access to the Kubernetes cluster. Access to the web interface is facilitated using `kubectl port-forward`.

## Deploying

### Clone this repo

```
git clone https://github.com/cased/try-shell
cd try-shell
```
### Update configuration

Obtain the value of `CASED_SHELL_SECRET` from your Shell Instance's settings page (e.g. https://app.cased.com/shell/programs/shell_EXAMPLE/settings) and set it in `.env`.

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

While you're there, also make sure to set the value of CASED_SHELL_HOSTNAME to `localhost:NNNN` where `NNNN` is the port of your Shell Instance.

### Preview configuration

```
kubectl kustomize examples/kubernetes | less
```

### Create a namespace

```
kubectl create ns shell
```

### Apply configuration to the shell namespace

```
kubectl -n shell apply -k examples/kubernetes
```

> Example output:

```
serviceaccount/cased-shell created
serviceaccount/kubectl-sshd created
clusterrolebinding.rbac.authorization.k8s.io/kubectl-sshd unchanged
configmap/cased-shell-2m4ggff27h created
configmap/jump-k722fttfc4 created
configmap/kubectl-sshd-gd2m2mbtf4 created
secret/cased-shell-b5gc526665 created
service/cased-shell created
service/kubectl-sshd created
deployment.apps/cased-shell created
deployment.apps/kubectl-sshd created
```

### Inspect status

```
kubectl -n shell get all
```

> Example output:

```
NAME                               READY   STATUS    RESTARTS   AGE
pod/cased-shell-b976bf64f-ps52l    1/2     Running   0          26s
pod/kubectl-sshd-c99b55f9d-bqp2b   1/1     Running   0          26s

NAME                   TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)   AGE
service/cased-shell    ClusterIP   10.32.0.122   <none>        443/TCP   26s
service/kubectl-sshd   ClusterIP   10.32.0.105   <none>        22/TCP    26s

NAME                           READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/cased-shell    0/1     1            0           26s
deployment.apps/kubectl-sshd   1/1     1            1           26s

NAME                                     DESIRED   CURRENT   READY   AGE
replicaset.apps/cased-shell-b976bf64f    1         1         0       26s
replicaset.apps/kubectl-sshd-c99b55f9d   1         1         1       26s
```

## Port forward

Run the following, making sure to replace `NNNN` with the port of your Shell Instance:

```
kubectl -n shell port-forward service/cased-shell NNNN:http
```
### Visit your shell

Next, visit your shell instance at http://localhost:NNNN. After logging in, click the `kubectl` prompt to run interactive commands using `kubectl`, or use the `event log` link to see a stream of all Kubernetes events. Edit `jump.yaml` to add additional entries, making sure to run `kubectl apply` again to apply the change.
