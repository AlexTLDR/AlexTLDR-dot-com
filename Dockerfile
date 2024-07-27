FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -o alextldr-dot-com .
CMD ./alextldr-dot-com

FROM alpine:latest
WORKDIR /app
COPY templates templates
COPY --from=builder /app/alextldr-dot-com .
CMD ./alextldr-dot-com