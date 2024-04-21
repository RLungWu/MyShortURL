FROM golang:alpine as builder

WORKDIR /app

ADD . /app/

WORKDIR /app/cmd/MyShortURL

RUN go mod download

RUN go build -o /build/main .

#stage 2

FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . /app

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 3000

CMD ["./main"]
