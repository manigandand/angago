# angago

[![CircleCI](https://circleci.com/gh/manigandand/angago/tree/master.svg?style=svg)](https://circleci.com/gh/manigandand/angago/tree/master)
[![Coverage Status](https://img.shields.io/codecov/c/gh/manigandand/angago.svg?logo=codecov&style=for-the-badge)](https://codecov.io/gh/manigandand/angago)
[![](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/manigandand/angago)

angago(anga po/அங்க போ) means Go there. Localhost Proxy Tunnel

![Image of Yaktocat](/asset/demo.gif)

```shell
go build
./angago
```

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

### TODO

-   Add TLS support
