FROM golang:1.27-alpine3.23 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY . .
RUN go build -o /link-cutter ./cmd/main.go

FROM alpine:3.23 as run

COPY --from=build /link-cutter /link-cutter

EXPOSE 8081
CMD ["/link-cutter"]