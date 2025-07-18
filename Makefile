build:
    go build -o bin/tanto ./cmd/tanto

test:
    go fmt ./...
    go vet ./...
    go test ./...

run-up:
    ./bin/tanto up topo.yml

run-down:
    ./bin/tanto destroy <lab-id>
