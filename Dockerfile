FROM golang:1.19-alpine as build

# Multi stage build from https://stackoverflow.com/questions/47028597/choosing-golang-docker-base-image

RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates

RUN addgroup -S netgear
RUN adduser -S -u 10000 -g netgear netgear

WORKDIR /src

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./main.go .
COPY ./utils ./utils

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app .

FROM scratch as final

COPY --from=build /app /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd

USER netgear

ENTRYPOINT [ "/app" ]