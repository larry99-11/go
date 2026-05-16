FROM golang:1.26 AS build_stage

WORKDIR /app
# Copy everything from web-app dir to our img
COPY . .
# Need to build a statically linked binary. Update your Go build, makes the binary statically linked (means it runs anywhere)
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine:latest
WORKDIR /app
# Install security certificates (often needed for HTTPS/APIs)
RUN apk --no-cache add ca-certificates
COPY --from=build_stage /app/server .
COPY --from=build_stage /app/templates/ ./templates/

# Ensure our server binary is executable
RUN chmod +x ./server
CMD ["./server"]
