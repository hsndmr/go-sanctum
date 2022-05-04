FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN apk add build-base

RUN cp .env.example .env

RUN go build -o go-sanctum

EXPOSE 3000

CMD [ "./go-sanctum" ]