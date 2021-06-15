FROM golang:alpine

RUN mkdir pokeprac/

WORKDIR /pokeprac

COPY . .

RUN go mod download

CMD ["go","run","main.go"]