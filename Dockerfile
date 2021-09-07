FROM golang:1.16

WORKDIR /app

RUN apt-get update
RUN apt install -y protobuf-compiler

RUN go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0 \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

RUN go get google.golang.org/grpc
RUN go get github.com/golang/protobuf/protoc-gen-go

COPY . .

RUN make bin-deps
RUN make

EXPOSE 9000 7002 8080

WORKDIR /app/bin

CMD ["./main"]
