#!/bin/bash
# generate.sh - proto 파일을 Go 코드로 생성하는 스크립트

set -e

PROTO_DIR="../protos"
OUT_DIR="../gen/go"

mkdir -p "${OUT_DIR}"

while IFS= read -r -d '' PROTO_SUBDIR; do
  PROTO_FILES=$(find "${PROTO_SUBDIR}" -maxdepth 1 -name "*.proto")
  if [ -n "${PROTO_FILES}" ]; then
    SUBDIR_NAME=$(basename "${PROTO_SUBDIR}")
    OUT_SUBDIR="${OUT_DIR}/${SUBDIR_NAME}"
    mkdir -p "${OUT_SUBDIR}"
    protoc \
      -I="${PROTO_SUBDIR}" \
      --go_out="${OUT_SUBDIR}" \
      --go-grpc_out="${OUT_SUBDIR}" \
      --go_opt=paths=source_relative \
      --go-grpc_opt=paths=source_relative \
      ${PROTO_FILES}
  fi
done < <(find "${PROTO_DIR}" -type d -print0)