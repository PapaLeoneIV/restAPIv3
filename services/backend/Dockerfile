FROM golang:alpine

WORKDIR /students_api
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build .

CMD ["./students"]
EXPOSE 8443
