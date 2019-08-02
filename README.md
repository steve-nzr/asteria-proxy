# asteria-proxy
[![forthebadge](https://forthebadge.com/images/badges/built-by-developers.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![Build Status](https://travis-ci.org/steve-nzr/asteria-proxy.svg?branch=master)](https://travis-ci.org/steve-nzr/asteria-proxy)
[![Docker Pulls](https://img.shields.io/docker/pulls/asteriatools/proxy.svg)](https://hub.docker.com/r/asteriatools/proxy/)
[![GitHub issues](https://img.shields.io/github/issues/steve-nzr/asteria-proxy.svg)](https://github.com/steve-nzr/asteria-proxy/issues)
[![GitHub stars](https://img.shields.io/github/stars/steve-nzr/asteria-proxy.svg?style=social&label=Star)](https://github.com/steve-nzr/asteria-proxy)
[![GitHub watchers](https://img.shields.io/github/watchers/steve-nzr/asteria-proxy.svg?style=social&label=Watchers)](https://github.com/steve-nzr/asteria-proxy)

<p align="center">
  <img src="https://raw.githubusercontent.com/steve-nzr/asteria-proxy/master/assets/banner.png" />
</p>

**asteria-proxy** is an open-source project which integrates in the **asteria** tools suite.
It aims to provide an open messaging API through which you can develop your game servers without worrying about the networking layer.

We provide you a **scalable proxy** that relays client's events like connections & messages through a message queue backed by a stable infrastructure powered by docker & kubernetes.

## Installation
---
Make sure you have these tools installed on your machine
```
Docker (18.09+)
Docker Compose (1.24+)
```

### docker and/or compose
- **Sample with working receiver** [Here](https://github.com/steve-nzr/asteria-sample-receiver)
- **Compose proxy+rabbitmq only** Go to [the compose example folder](https://github.com/steve-nzr/asteria-proxy/tree/master/examples/compose)
- **Docker proxy only** `docker run --it -p 23000:23000 --env-file=.env asteriatools/proxy:latest`


### Kubernetes
- soon...

### From scratch
```
Install rabbitmq & set your environnement variables reflecting those that are in configs/env/base.env.
```

Il will pull all the required images & build the project if necessary.

## Running test suite
- Unit tests
```
$ docker-compose -f docker-compose.test.yml run --rm unit
```
- Functional tests
```
$ not yet...
```

## Documentation
Advanced documentation can be found in the **docs** folder.

## Maintainers
- [@steve-nzr](https://github.com/steve-nzr)

## Supporters
- [@zetsumi](https://github.com/zetsumi)
- [@ImNotAVirus](https://github.com/ImNotAVirus)
