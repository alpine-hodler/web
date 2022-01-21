FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

WORKDIR /app/cmd/graphql
RUN go build

WORKDIR /app

EXPOSE $PORT

CMD ./cmd/graphql/graphql start --port=$PORT --env=$AUTH_FILE
