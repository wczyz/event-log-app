FROM golang:1.18
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY src src/
RUN cd src && go build -o main .
CMD ["/app/src/main"]
