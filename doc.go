//go:generate protoc --go_out=plugins=grpc:./../../../ -Isrc/main/protobuf src/main/protobuf/devim/protobuf/page.proto
//go:generate protoc --go_out=plugins=grpc:./../../../ -Isrc/main/protobuf src/main/protobuf/devim/protobuf/time.proto
//go:generate protoc --go_out=plugins=grpc:./../../../ -Isrc/main/protobuf src/main/protobuf/devim/protobuf/relay.proto
package proto_utils
