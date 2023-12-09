##########################
# STAGE 1
##########################
FROM golang:1.20-alpine as builder
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
RUN mkdir -p /go/src/github.com/grippenet/webhook-service
ENV GO111MODULE=on
ADD . /go/src/github.com/grippenet/webhook-service/
WORKDIR /go/src/github.com/grippenet/webhook-service
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
WORKDIR /go/src/github.com/grippenet/webhook-service/cmd/webhook-service
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o app .

##########################
# STAGE 2
##########################
FROM scratch
# we need those on scratch for connecting to mongoDB:
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/grippenet/webhook-service/cmd/webhook-service/app /app/
WORKDIR /app
EXPOSE 3253:3253
CMD ["./app"]
