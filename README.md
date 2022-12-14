# Netgear Nighthawk Metrics

Prometheus Exporter Netgear Routers.

Tested models
- `Nighthawk MR60 Mesh`
- `Nighthawk R8000P`

## Intro
Parses HTML from http://www.routerlogin.com/RST_stattbl.htm from two router models.

Note: This script ends any other sessions on the router admin page.

Netgear products have bad customer support and extensibility. If you're reading this, you're probably not going to replace your $300 router right away so I hope this fills a need.

`Nighthawk R8000P` 90s interface

![R8000P](./refrence/R8000P.png)

`Nighthawk MR60 Mesh` new UI

![MR60](./refrence/MR60.png)


## Environment Variables
```
ROUTER_URL="http://www.routerlogin.com/"
ROUTER_PASSWORD=""
ROUTER_USERNAME="admin"
PORT="8080"
```

Flags
```
export ROUTER_PASSWORD=""
go run . --url --username --port
```

## Helm Chart Usage
- Add Repository
```
helm repo add connorsapps-netgear https://connorsapps.github.io/netgear-nighthawk-metrics
```

- Install Chart
```
helm install netgear-metrics connorsapps-netgear/netgear-metrics
```

- Delete Chart ☹️
```
helm delete netgear-metrics
```