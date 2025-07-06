# Create new build from base image
FROM golang:latest

# Change work directory
WORKDIR /base

# Install air
RUN go install github.com/air-verse/air@latest

COPY go.* ./
RUN go mod download

# Copy files to /base
COPY . .

EXPOSE 8002

# Run command
CMD ["air", "-c", ".air.toml"]