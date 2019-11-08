# angago

[![CircleCI](https://circleci.com/gh/manigandand/angago/tree/master.svg?style=shield)](https://circleci.com/gh/manigandand/angago/tree/master)
[![Coverage Status](https://img.shields.io/codecov/c/gh/manigandand/angago.svg?logo=codecov&style=for-the-badge)](https://codecov.io/gh/manigandand/angago)

<!-- [![](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/manigandand/angago) -->

[![License](https://img.shields.io/badge/license-MIT%20License-blue.svg)](https://github.com/manigandand/angago/blob/master/LICENSE)

angago(anga po/அங்க போ) means Go there.
Localhost Proxy Tunnel(reverse proxy).
[Demo Video](https://www.youtube.com/watch?v=pcRgChXpU94)

angago is small localhost proxy tunnel(reverse proxy) which routes to different
application servers based on **request** `Host/Hostname` from [local DNS](#local-dns).

angago determines the **target url** to proxy based on the simple [yaml configuration](#angago-configuration).
It's a simple `map[string]string`.

```go
map[host]targethost
```

![Image of Yaktocat](/asset/demo.gif)

> build and run from source

```shell
source .env
go build
./angago
```

or

```shell
./deployment.sh
```

> run from docker image

```shell
docker run -v /Users/manigandand/.angago/config.yaml:/mnt/Users/manigandand/.angago/config.yaml -e 'ANGAGO_CONFIG_PATH=/mnt/Users/manigandand/.angago/config.yaml' -it manigandanjeff/angago:latest
```

or

```shell
./docker_run.sh /Users/manigandand/.angago/config.yaml
```

> angago architecture

![Image of Yaktocat](/asset/angago_diagram.png)

## Local DNS

> /etc/hosts

```
##
# Host Database
#
# localhost is used to configure the loopback interface
# when the system is booting.  Do not change this entry.
##
127.0.0.1	localhost
255.255.255.255	broadcasthost
::1             localhost
127.0.0.1	status.gopherhut.com
127.0.0.1	app.gopherhut.com
127.0.0.1	api.gopherhut.com
127.0.0.1	nope.gopherhut.com
# Added by Docker Desktop
# To allow the same kube context to work on the host and the container:
127.0.0.1 kubernetes.docker.internal
# End of section
```

## angago configuration

`$HOME/.angago/config.yaml`

```yaml
# apiVersion: api/v1
# apiNamespace: angago.gopherhut.com
kind: Domains
meta:
    shortName: domains
tls: false
domains:
    app.gopherhut.com: localhost:8081
    status.gopherhut.com: localhost:8082
    admin.gopherhut.com: localhost:8083
```

## TODO

-   Add TLS support

## Licence

MIT
