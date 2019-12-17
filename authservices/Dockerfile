# FROM golang:1.8

# WORKDIR /go/src/app
# COPY . .

# RUN go get -d -v ./...
# RUN go install -v ./...

# RUN go build -o main .

FROM golang:1.9

ADD . /go/src/auth
WORKDIR /go/src/auth
RUN go get auth
RUN go install
ENTRYPOINT ["/go/bin/auth"]






# FROM golang:alpine as builder

# RUN apk --no-cache add git

# WORKDIR /app/mini-auth-service

# COPY . .

# RUN go mod download

# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mini-auth-service

# FROM alpine:latest

# RUN apk --no-cache add ca-certificates

# RUN mkdir /app
# WORKDIR /app

# COPY --from=builder /app/mini-auth-service/mini-auth-service .

# CMD ["./mini-auth-service"]