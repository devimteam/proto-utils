//go:generate protoc --go_out=plugins=grpc:./../../../ -Isrc/main/protobuf src/main/protobuf/devim/protobuf/page.proto
//go:generate protoc --go_out=plugins=grpc:./../../../ -Isrc/main/protobuf src/main/protobuf/devim/protobuf/time.proto
//go:generate protoc --go_out=plugins=grpc:./../../../ -Isrc/main/protobuf src/main/protobuf/devim/protobuf/relay.proto
//go:generate protoc --go_out=plugins=grpc:./../../../ -Isrc/main/protobuf src/main/protobuf/devim/protobuf/dadata.proto
//go:generate protoc --go_out=plugins=grpc:./../../../ -Isrc/main/protobuf src/main/protobuf/devim/protobuf/collect.proto
//go:generate protoc --go_out=plugins=grpc:./../../../ -Isrc/main/protobuf src/main/protobuf/devim/protobuf/person.proto
package proto_utils
