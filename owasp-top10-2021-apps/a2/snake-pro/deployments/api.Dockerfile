FROM golang

WORKDIR /go/src/github.com/globocom/secDevLabs/owasp-top10-2021-apps/a2/snake-pro/app

COPY app/go.mod app/go.sum ./
RUN go mod download

RUN openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 \
    -subj "/CN=localhost" \
    -keyout key.pem  -out cert.pem

COPY app/ ./
