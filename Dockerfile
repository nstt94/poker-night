FROM golang:latest

RUN mkdir /app

ADD . /app
# Set destination for COPY
WORKDIR /app

# Build
RUN go build -o pokernight

# Run
CMD ["/app/pokernight"]