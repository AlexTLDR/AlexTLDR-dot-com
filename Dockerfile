# Multi-stage Dockerfile for Go app with Tailwind CSS and DaisyUI

# Stage 1: Build CSS with Node.js
FROM node:18-alpine AS css-builder

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install Node.js dependencies (including dev dependencies for CSS build)
RUN npm ci

# Copy source files needed for CSS build
COPY src/ ./src/
COPY tailwind.config.js ./
COPY templates/ ./templates/
COPY static/ ./static/

# Create CSS directory and build CSS
RUN mkdir -p static/css
RUN npm run build-css-prod

# Stage 2: Build Go application
FROM golang:latest AS go-builder

WORKDIR /app

# Install templ CLI
RUN go install github.com/a-h/templ/cmd/templ@latest

# Copy Go module files
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy source code
COPY . .

# Generate templ files
RUN templ generate

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Stage 3: Runtime
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy the built Go binary
COPY --from=go-builder /app/main .

# Copy static assets
COPY --from=go-builder /app/static ./static

# Copy the built CSS from css-builder stage
COPY --from=css-builder /app/static/css ./static/css

# Expose port
EXPOSE 8080

# Set environment variables
ENV PORT=8080
ENV GIN_MODE=release

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/ || exit 1

# Run the application
CMD ["./main"]
