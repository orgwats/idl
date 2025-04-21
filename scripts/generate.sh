#!/bin/bash
set -e

PROTO_DIR="./protos"
GO_OUT_DIR="./gen/go"
TS_OUT_DIR="./gen/ts"

mkdir -p "${GO_OUT_DIR}"
mkdir -p "${TS_OUT_DIR}"

while IFS= read -r -d '' PROTO_SUBDIR; do
  PROTO_FILES=$(find "${PROTO_SUBDIR}" -maxdepth 1 -name "*.proto")
  if [ -n "${PROTO_FILES}" ]; then
    SUBDIR_NAME=$(basename "${PROTO_SUBDIR}")

    GO_OUT_SUBDIR="${GO_OUT_DIR}/${SUBDIR_NAME}"
    mkdir -p "${GO_OUT_SUBDIR}"
    protoc \
      -I="${PROTO_SUBDIR}" \
      --go_out="${GO_OUT_SUBDIR}" \
      --go-grpc_out="${GO_OUT_SUBDIR}" \
      --go_opt=paths=source_relative \
      --go-grpc_opt=paths=source_relative \
      ${PROTO_FILES}

    TS_OUT_SUBDIR="${TS_OUT_DIR}/${SUBDIR_NAME}"
    mkdir -p "${TS_OUT_SUBDIR}"
    protoc \
      --plugin=protoc-gen-ts_proto=./node_modules/.bin/protoc-gen-ts_proto \
      --ts_proto_out="${TS_OUT_SUBDIR}" \
      --ts_proto_opt=nestJs=true,useObservable=true,outputClientImpl=grpc-js \
      -I "${PROTO_SUBDIR}" \
      ${PROTO_FILES}
  fi
done < <(find "${PROTO_DIR}" -type d -print0)