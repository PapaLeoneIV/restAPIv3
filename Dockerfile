FROM golang:alpine

WORKDIR /students_api
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/api ./cmd/api 

CMD ["/students_api/bin/api"]
EXPOSE 8081
