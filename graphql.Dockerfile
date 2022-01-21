FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

WORKDIR /app/cmd/graphql
RUN go build

WORKDIR /app

EXPOSE 8080

CMD ["./cmd/graphql/graphql", "start", "--port=8080", "--env=.auth.env"]
