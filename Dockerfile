FROM golang:1.23 AS development

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN go build -o /app/tmp/main -buildvcs=false .

CMD ["air"]
