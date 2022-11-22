# Netgear Nighthawk Metrics

Prometheus Exporter for MR60 Mesh & R8000P Netgear Routers

Parses HTML from http://www.routerlogin.com/RST_stattbl.htm from two router models.

Netgear products have bad customer support and extensibility. If you're reading this then you might not want to buy a new router.

`Nighthawk R8000P` 90s interface

![R8000P](./refrence/R8000P.png)

`Nighthawk MR60 Mesh` new UI

![MR60](./refrence/MR60.png)


Environment Variables
```
NETGEAR_URL="http://www.routerlogin.com/"
NETGEAR_PASSWORD="admin"
NETGEAR_USERNAME="2"
PORT="8080"
```

Flags
```
--url
--username
--port
```
