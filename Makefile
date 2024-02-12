BUILD_DIR=build

run:
	go run cmd/automata/main.go

build: 
	go build -o="$(BUILD_DIR)/automata" ./cmd/automata
