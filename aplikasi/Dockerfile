FROM golang:1.17.1-buster

RUN apt update & apt install

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o binary

CMD ["/app/binary"]