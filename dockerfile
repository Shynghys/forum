FROM golang:latest
# Set the Current Working Directory inside the container
WORKDIR /app

ENV GOPATH=/go/src/forum
# Populate the module cache based on the go.{mod,sum} files.
COPY go.mod .       
COPY go.sum .
RUN go mod download
COPY . .
# RUN pwd
# RUN ls

# Build the Go app
# RUN go get ./
# RUN go build
# RUN go install
RUN go build -o forum .

# Run the binary program produced by go install
CMD ["./forum"]