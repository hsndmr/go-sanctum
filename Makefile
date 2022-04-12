schema:
	go generate ./ent
migrate: 
	go run ./cobra/cobra.go migrate