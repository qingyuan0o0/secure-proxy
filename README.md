# Secure Proxy

Secure Proxy for GAE/Go

- Support Edge Cache (Cloud CDN)

## Required

- [Cloud SDK](https://cloud.google.com/sdk/)
- [App Engine SDK for Go](https://cloud.google.com/appengine/docs/go/download)

## Recommended

- [direnv](https://github.com/direnv/direnv)

```sh
$ curl https://sdk.cloud.google.com | bash
$ gcloud auth login
$ open https://cloud.google.com/appengine/docs/go/download
$ export PATH="$PATH:$HOME/go_appengine"
$ git clone git@github.com:s-aska/secure-proxy.git
$ cd secure-proxy/src
```

## direnv

```sh
$ sh env.sh
or
$ cp env.sh .envrc
```

## Format

```sh
$ gofmt -d -w ./lib
```

## Run

```sh
$ goapp serve
```

## Deploy

```sh
$ goapp deploy -application apple-secure-proxy -version 1
```
