FROM golang:1.23

# Set app workdir
WORKDIR /go/src/app

# Copy dependencies list
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy application sources
COPY . .

# Build app
RUN go build -o app github.com/Russia9/Muskrat/cmd/main

# Run app
CMD ["./app"]
