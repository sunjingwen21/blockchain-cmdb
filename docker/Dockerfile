FROM golang:1.21-alpine AS backend-builder

WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./
RUN go build -o cmdb-api main.go

# Frontend build stage
FROM node:18-alpine AS frontend-builder

WORKDIR /app
COPY frontend/package*.json ./
RUN npm ci

COPY frontend/ ./
RUN npm run build

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy backend binary
COPY --from=backend-builder /app/cmdb-api .

# Copy frontend build
COPY --from=frontend-builder /app/build ./frontend

EXPOSE 8080

CMD ["./cmdb-api"]
