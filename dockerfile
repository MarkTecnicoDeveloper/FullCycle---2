FROM golang:1.14.6-alpine as builder

WORKDIR /go/src/app

COPY . .

# flagas para gerar um binario reduzido
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w"

FROM scratch

WORKDIR /go/src/app

COPY --from=builder /go/src/app/app .
COPY --from=builder /go/src/app/index.html .

CMD ["./app"]
#CMD ["go","run","."]