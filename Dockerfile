FROM golang:1.20-alpine as build

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/app cmd/main.go

FROM alpine as final

COPY --from=build /usr/local/src/bin/app /

COPY config/config.yml /config/config.yml

CMD ["/app"]