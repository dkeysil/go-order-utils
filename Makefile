#######################################################
############## formats, lint, and tests ###############
#######################################################

.PHONY: fmt
fmt:
	@echo "----------------------------------------------------------------"
	@echo " ⚙️  Formatting code..."
	@echo "----------------------------------------------------------------"
	gofmt -s -w ./.

.PHONY: lint
lint:
	@echo "----------------------------------------------------------------"
	@echo " ⚙️  Linting code..."
	@echo "----------------------------------------------------------------"
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.11.4
	golangci-lint run ./... --config=.golangci.yaml
	golangci-lint fmt --diff ./...
	go mod tidy
	@echo "Linting complete!"

.PHONY: test
test:
	@echo "----------------------------------------------------------------"
	@echo " ⚙️  Testign the code..."
	@echo "----------------------------------------------------------------"
	GOPRIVATE=${PRIVATE_REPOS} go test ./... -v
	@echo "Tests complete!"

.PHONE: build
build:
	@echo
