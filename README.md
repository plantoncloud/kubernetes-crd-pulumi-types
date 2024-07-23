# kubernetes-crd-pulumi-types

Using [crd2pulumi](https://github.com/pulumi/crd2pulumi) makes is possible to generate strong types for 
Kubernetes Custom Resource Definitions, which makes it easy to create those custom resources in pulumi modules.

1. Install `crd2pulumi` cli.

```shell
brew install pulumi/tap/crd2pulumi
```

2. Generate Pulumi Types for CRDs ex: Istio CRDs

```bash
crd2pulumi --goPath=pkg/istio https://raw.githubusercontent.com/istio/istio/master/manifests/charts/base/crds/crd-all.gen.yaml
```
