
FROM golang:1.23-alpine

WORKDIR /app
RUN go install github.com/air-verse/air@latest


COPY go.* ./

RUN go mod download

COPY . .


EXPOSE 8019

RUN go build -o main main.go

CMD ["air",  "-c", ".air.toml"]

