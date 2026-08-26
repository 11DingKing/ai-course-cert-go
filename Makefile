test:
	GOCACHE=/tmp/gocache-gochuti go test ./... -count=1
race:
	GOCACHE=/tmp/gocache-gochuti go test -race ./... -count=1
run:
	go run ./cmd/server
