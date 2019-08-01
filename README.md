# asteria-proxy
[![forthebadge](https://forthebadge.com/images/badges/built-by-developers.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com) 

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

Then run
```sh
$ docker-compose up -d
```

or ...
```
Install rabbitmq & set your environnement variables reflecting those that are in configs/env/base.env.
```

Il will pull all the required images & build the project if necessary.

## Documentation
Advanced documentation can be found in the **docs** folder.

## Maintainers
- [@steve-nzr](https://github.com/steve-nzr)

## Supporters
- [@zetsumi](https://github.com/zetsumi)
- [@ImNotAVirus](https://github.com/ImNotAVirus)
