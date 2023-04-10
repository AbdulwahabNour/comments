FROM golang:1.20 As builder

RUN mkdir /app
ADD . /app
WORKDIR /app

Run CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

FROM alpine:latest As production
COPY --from=builder /app .
CMD [ "./app" ]

