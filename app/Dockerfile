FROM golang:1.18-alpine
RUN apk update && apk add git

COPY . /src/app
WORKDIR /src/app

RUN go install github.com/cosmtrek/air@latest && go mod tidy

CMD sh -c "air -c .air.toml"