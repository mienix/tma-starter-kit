FROM golang:1.22-alpine as build
RUN apk add --no-cache ca-certificates

WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w -extldflags "-static"' -o app ./cmd/user-api/main.go

FROM gcr.io/distroless/static:nonroot
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /src/app /app
USER 65532:65532
ENTRYPOINT ["/app"]
