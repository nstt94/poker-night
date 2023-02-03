FROM golang:latest

# Set destination for COPY
WORKDIR /app

# Download Go module
COPY go.mod .
RUN go mod download

# Build
RUN go build -o /pokernight

# Run
CMD ["/pokernight"]