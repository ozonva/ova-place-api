FROM golang:1.16 as build

WORKDIR /app

RUN apt-get update && \
  apt install -y protobuf-compiler && \
  go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0 \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0 && \
  go get google.golang.org/grpc && \
  go get github.com/golang/protobuf/protoc-gen-go

COPY . .

RUN make bin-deps && \
  make

EXPOSE 9000 7002 8080

WORKDIR /app/bin

FROM scratch

COPY --from=build /app/bin .
COPY --from=build /app/.env .

CMD ["./main"]
