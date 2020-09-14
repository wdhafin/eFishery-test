FROM golang:1.14.2-alpine3.11

WORKDIR /app

COPY . ./

# Download project dependencies.
RUN go mod download

RUN go build -o main .

CMD [ "/app/main" ]