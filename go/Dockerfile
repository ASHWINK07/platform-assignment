FROM golang:latest

ENV GO111MODULE=on
ENV PORT=8080
WORKDIR /app/server
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
EXPOSE 8080
RUN go build -o httpserver
CMD ["./httpserver"]