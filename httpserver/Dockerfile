FROM golang:1.14 as builder
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build .
RUN mkdir publish && cp ./ publish
FROM scratch
WORKDIR /app

COPY --from=builder /app/publish .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/cert
ENV PORT=80
EXPOSE 80
ENTRYPOINT ["./httpserver"]
