FROM golang:1.13.5
WORKDIR /app
COPY . .
RUN go build -o rest-app
ENTRYPOINT ["./rest-app"]