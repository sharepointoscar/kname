# Cocktail
With all the madness around me, I decided to do something I want todo, vs what I have to do.  Combine two passions and this is what comes out of it.

I am a big fan of mixology, and well I work in tech.  One day I decided I hated coming up with names for my Kubernetes clusters as I test stuff, and so Cocktail was born.

# What is this?
Cocktail is a simple CLI that you can use to spit out a random cocktail name which you can use to name whatever, but I use it for Kubernetes cluster names as I have a difficult time coming up with them.


# Installing

## Using Jenkins X CLI

I use Jenkins X on GKE and EKS (soon AKS), but here is how I quickly create my cluster passing the random name.
```bash
jx create cluster gke --skip-installation -n $(cocktail generate-name)
```
## GCloud CLI usage

If you are using the `gcloud` CLI, you can just execute the following (add whatever other flags you want)

```bash
gcloud container clusters create $(cocktail generate-name)
```