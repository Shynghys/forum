FROM golang:1.15.7
# Set the Current Working Directory inside the container
WORKDIR /app

# Populate the module cache based on the go.{mod,sum} files.
COPY go.mod .       
COPY go.sum .
RUN go mod download
COPY . .


# Build the Go app
RUN go build -o forum .

EXPOSE 8000

# Run the binary program produced by go install
CMD ["./forum"]