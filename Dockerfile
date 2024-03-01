FROM golang:1.22

WORKDIR /src

COPY ./ /

RUN go build -o /bin/obstwiese ./server.go



FROM alpine
WORKDIR /root/
COPY --from=builder /bin/obstwiese ./app
CMD ["./app/obstwiese"]