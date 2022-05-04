schema:
	go generate ./ent && go run ./cobra/cobra.go migrate

test:
	go test ./...