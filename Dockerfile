FROM golang:1.18-alpine as base

WORKDIR /builder
RUN apk add upx

COPY go.mod go.sum /builder/
RUN go mod download

COPY . .

RUN go build \
    -ldflags "-s -w" \
    -o /builder/main /builder/main.go
RUN upx -9 /builder/main


# runner image
FROM gcr.io/distroless/static:latest
WORKDIR /app
COPY --from=base /builder/main main

ENV PORT=8080

EXPOSE 8080

CMD ["./main"]


