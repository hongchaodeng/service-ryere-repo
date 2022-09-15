##
## Build
##
FROM golang:1.19 AS builder

WORKDIR /workdir

COPY ./ ./

RUN go build -o build/app main.go

##
## Run
##
FROM golang:1.19

WORKDIR /workdir

COPY --from=builder /workdir/build/ ./

EXPOSE 8000

CMD ["sh", "-c", "./app"]
