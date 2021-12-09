FROM golang:1.16 AS builder

WORKDIR /builder
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod tidy
RUN go mod download

COPY ./ ./

RUN GOOS=linux GOARCH=amd64 go build -o /builder/ppp .

FROM alpine:latest

COPY --from=builder /builder/ppp  /
COPY --from=builder /builder/.env /
CMD /ppp