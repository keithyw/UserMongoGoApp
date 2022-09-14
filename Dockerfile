FROM golang:alpine AS builder

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

WORKDIR /app

COPY src/ ./

RUN go mod download
RUN go mod verify
RUN go build -o /message_app

FROM golang:alpine

RUN apk --no-cache add ca-certificates git

WORKDIR /

COPY --from=builder /message_app /message_app
EXPOSE 8081
CMD ["/message_app"]
