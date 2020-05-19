# tbls-cacoo [![Build Status](https://github.com/k1LoW/tbls-cacoo/workflows/build/badge.svg)](https://github.com/k1LoW/tbls-cacoo/actions) [![GitHub release](https://img.shields.io/github/release/k1LoW/tbls-cacoo.svg)](https://github.com/k1LoW/tbls-cacoo/releases)

`tbls-cacoo` is an external subcommand of tbls for [Cacoo](https://cacoo.com).

## Usage

tbls-cacoo is provided as an external subcommand of [tbls](https://github.com/k1LoW/tbls).

`tbls cacoo csv` generate CSV for [Cacoo's Database Schema Importer](https://support.cacoo.com/hc/en-us/articles/360045672494).

``` console
$ tbls cacoo csv
```

``` console
$ tbls cacoo csv --out cacoo.csv
```

## Install

**deb:**

Use [dpkg-i-from-url](https://github.com/k1LoW/dpkg-i-from-url)

``` console
$ export TBLS_CACOO_VERSION=X.X.X
$ curl -L https://git.io/dpkg-i-from-url | bash -s -- https://github.com/k1LoW/tbls-cacoo/releases/download/v$TBLS_CACOO_VERSION/tbls-cacoo_$TBLS_CACOO_VERSION-1_amd64.deb
```

**RPM:**

``` console
$ export TBLS_CACOO_VERSION=X.X.X
$ yum install https://github.com/k1LoW/tbls-cacoo/releases/download/v$TBLS_CACOO_VERSION/tbls-cacoo_$TBLS_CACOO_VERSION-1_amd64.rpm
```

**homebrew tap:**

```console
$ brew install k1LoW/tap/tbls-cacoo
```

**manually:**

Download binary from [releases page](https://github.com/k1LoW/tbls-cacoo/releases)

**go get:**

```console
$ go get github.com/k1LoW/tbls-cacoo
```

## Requirements

- [tbls](https://github.com/k1LoW/tbls) > 1.38.2
