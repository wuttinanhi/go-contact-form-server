#build stage
FROM golang:1.23.2 AS builder
RUN apt-get install -y git 
WORKDIR /go/src/app

ENV CGO_ENABLED=1

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /out/app

FROM gcr.io/distroless/base-debian12 AS final
COPY ./pages /pages
COPY --from=builder /out/app /app
CMD ["/app"]
