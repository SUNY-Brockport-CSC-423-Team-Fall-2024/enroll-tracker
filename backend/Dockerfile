FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /enroll-tracker-api ./cmd/app

CMD ["/enroll-tracker-api"]

