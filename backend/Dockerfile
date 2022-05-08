FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN cd cmd/ && go build -o ../backend

EXPOSE 8000

CMD ["./backend" ]