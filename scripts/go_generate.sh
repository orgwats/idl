#!/bin/bash
# generate.sh - proto 파일을 Go 코드로 생성하는 스크립트

set -e

PROTO_DIR="../protos"
OUT_DIR="../gen/go"

mkdir -p "${OUT_DIR}"

protoc \
  -I="${PROTO_DIR}" \
  --go_out="${OUT_DIR}" \
  --go-grpc_out="${OUT_DIR}" \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  $(find "${PROTO_DIR}" -name "*.proto")