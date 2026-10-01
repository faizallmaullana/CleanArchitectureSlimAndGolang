FROM golang:1.25-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGET=api
RUN go build -o /out/app ./cmd/${TARGET}

FROM alpine:3.20

RUN addgroup -S app && adduser -S -G app app

COPY --from=builder /out/app /app

USER app

ENTRYPOINT ["/app"]