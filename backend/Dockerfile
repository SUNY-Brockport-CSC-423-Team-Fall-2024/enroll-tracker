FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./src .

RUN CGO_ENABLED=0 GOOS=linux go build -o /enroll-tracker-api

CMD ["/enroll-tracker-api"]

