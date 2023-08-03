# Build Stage
FROM golang:1.20.7-alpine AS Build

WORKDIR /opt/go-grow
COPY . .

RUN go mod download
RUN go build -o /opt/go-grow/app cmd/app/main.go


# Deploy Image
FROM alpine:3.18

WORKDIR /opt/go-grow
COPY --from=Build /opt/go-grow/app /opt/go-grow/app
COPY --from=Build /opt/go-grow/env.env /opt/go-grow/env.env

EXPOSE 8080

ENTRYPOINT ["/opt/go-grow/app"]
