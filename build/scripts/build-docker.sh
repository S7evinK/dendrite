#!/usr/bin/env sh

GOOS="$(go env GOOS)"
GOARCH="$(go env GOARCH)"
if [ "${GOARCH}" = "arm" ]; then
    GOARM="$(go env GOARM)"
fi

TARGET="${GOOS}-${GOARCH}"
if [ "${GOARCH}" = "arm" ] && [ -n "${GOARM}" ]; then
    TARGET="${TARGET}-v${GOARM}"
fi
if [ "${GOOS}" = "windows" ]; then
    TARGET="${TARGET}.exe"
fi

for f in /out/*
do
  echo ${f}
  mv "$f" "$f-${TARGET}"
done