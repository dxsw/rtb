FROM golang:1.25-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o bidder ./cmd/bidder

FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=build /app/bidder /app/bidder

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/bidder"]
