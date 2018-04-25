FROM golang:1.9

WORKDIR /go/src/github.com/timkellogg/store/identification

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/timkellogg/store/identification/identification-service .

CMD ["./identification-service"]
