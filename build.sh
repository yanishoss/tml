#!/usr/bin/env sh

GOOS=js GOARCH=wasm go build -o ./bin/main.wasm ./cmd/wasm/wasm.go
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./bin/wasm_exec.js