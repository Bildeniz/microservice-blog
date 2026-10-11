FROM golang:1.27.1-alpine

WORKDIR /app

COPY . /app

RUN go get .

CMD [ "go",  "run", "main.go" ]
