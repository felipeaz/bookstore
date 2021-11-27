# Bookstore

## Generate proto file
protoc --go_out=plugins=grpc:build/ build/server/inventory/grpc/proto/inventory.proto