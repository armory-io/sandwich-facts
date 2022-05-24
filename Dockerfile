FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY sandwich.go ./

RUN go build -o /sandwich

CMD ["/sandwich"]