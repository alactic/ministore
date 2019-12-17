# FROM golang:1.8

# WORKDIR /go/src/app
# COPY . .

# RUN go get -d -v ./...
# RUN go install -v ./...

# RUN go build -o main .

FROM golang:1.9

ADD . /go/src/user
WORKDIR /go/src/user
RUN go get user
RUN go install
ENTRYPOINT ["/go/bin/user"]