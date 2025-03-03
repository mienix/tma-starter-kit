FROM golang:alpine as build 
RUN apk add --no-cache ca-certificates

WORKDIR /go/src/github.com/devflex-pro/tma-starter-kit/backend
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -ldflags '-extldflags "-static"' -o app ./cmd/user-api/main.go 

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt \
  /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/src/github.com/devflex-pro/tma-starter-kit/backend/app /app

ENTRYPOINT ["/app"]
