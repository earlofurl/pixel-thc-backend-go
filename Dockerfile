## Build stage
#FROM golang:1.19-alpine3.16 AS builder
#WORKDIR /app
#COPY . .
#RUN go build -o main main.go
## RUN apk add curl
## RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
#
## Run stage
#FROM alpine:3.16
#WORKDIR /app
#COPY --from=builder /app/main .
## COPY --from=builder /app/migrate.linux-amd64 ./migrate
## COPY app.env .
##COPY --from=builder /app/start.sh .
#COPY start.sh .
#RUN chmod +x start.sh
## RUN chmod +x /app/start.sh
#COPY db/migration ./db/migration
#
#EXPOSE 8080
#CMD [ "/app/main" ]
#ENTRYPOINT [ "start.sh" ]

# Build stage
FROM golang:1.19-alpine3.16 AS builder

WORKDIR /app

COPY . .

RUN go build -o /pixelthc

# Run stage
FROM alpine:3.16

WORKDIR /

COPY --from=builder /pixelthc /pixelthc
COPY app.env .
COPY internal/db/migration ./internal/db/migration

EXPOSE 8080

ENTRYPOINT [ "/pixelthc/cmd/pixelthc-entrypoint" ]
