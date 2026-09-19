FROM golang:1.27-alpine3.23 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY . .
RUN go build -o /link-cutter ./cmd/main.go

FROM alpine:3.23 as run

ENV DSN="host=localhost user=postgres password=postgres dbname=postgres port=4321 sslmode=disable"
ENV TOKEN="a8a370e71a28619bfe381b00d5c3f5a5473f973a7e892c7206b4282a283dfba1"
ENV DB_USER="postgres"
ENV DB_PASSWORD="postgres"
ENV DB_NAME="link"

COPY --from=build /link-cutter /link-cutter

EXPOSE 8081
CMD ["/link-cutter"]