FROM golang:1.13.5-alpine AS builder
WORKDIR /app
COPY . .
ENV GOPATH="$PWD"
RUN CGO_ENABLED=0 go build -o rest-app

FROM scratch
COPY --from=builder /app/rest-app /app/rest-app
EXPOSE 8080
ENTRYPOINT ["/app/rest-app"]