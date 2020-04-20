![cocktail](logo.png)
# Kname
 Combine several passions and this is what comes out of it.

I am a big fan of mixology, fitness andd the great outdoors. One day I decided I hated coming up with names for my [Kubernetes](https://kubernetes.io/) clusters as I test stuff, and so Kname was born.

# What is this?
Kname is a simple CLI that you can use to spit out a random name based on the theme you choose. You can use to name whatever, but I use it for Kubernetes cluster names as I have a difficult time coming up with them.


# Installing

Like any other `go` package.

```bash
go get github.com/sharepointoscar/kname
```

## Using Jenkins X CLI

I use Jenkins X on GKE and EKS (soon AKS), but here is how I quickly create my cluster passing the random name.
```bash
jx create cluster gke --skip-installation -n $(kname generate-name --theme cocktails)
```
## GCloud CLI usage

If you are using the `gcloud` CLI, you can just execute the following (add whatever other flags you want)

```bash
gcloud container clusters create $(kname generate-name --theme yoga)
```

## eksctl CLI usage

```bash
eksctl create cluster --name $(kname generate-name --theme national-parks)
```
