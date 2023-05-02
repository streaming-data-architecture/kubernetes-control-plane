## Installation (Linux)

Follow from official guide or,

```bash
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.18.0/kind-linux-amd64
chmod +x ./kind
sudo mv ./kind /usr/local/bin/kind
```

`kind get clusters`

`kind create cluster --name mf-cluster`

`kind delete  cluster --name mf-cluster`

`kind create cluster --name mf --config kind-ha-config.yaml`
