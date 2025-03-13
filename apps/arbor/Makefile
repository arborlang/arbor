export GO111MODULE=on
ARBOR_VERSION?=0.0.0-rc0

all:
	go build -o arbor ./cmd/arbor/main.go

publish:
	go build -ldflags "-X main.Version=$(ARBOR_VERSION)" -o arbor ./cmd/arbor/main.go

test_file: all
	./arbor build -wast -o test.wast test.ab
	
test_debug: build_debug
	./arbor build -wast -o test.wast test.ab

build_debug:
	go build -gcflags='all=-N -l' -tags nopkcs11 -ldflags='-linkmode internal' -o arbor ./cmd/arbor/main.go 

wast_test:
	wat2wasm -o test.wasm test.wast
	./arbor run --wasm --entrypoint main test.wasm

debug_run: test_debug
	wat2wasm -o test.wasm test.wast
	dlv exec ./arbor -- run --wasm --entrypoint main test.wasm

test_run: refresh test_file wast_test

refresh:
	rm -rf vendor/
	go mod vendor