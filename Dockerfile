FROM golang:1.22 AS builder

WORKDIR /src

COPY ./ /src

#RUN go build -o /bin/obstwiese ./server.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/obstwiese ./server.go



FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /bin/obstwiese .
COPY ./assets ./assets
CMD ["./obstwiese"]