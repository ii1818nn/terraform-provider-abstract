.PHONY: \
	package-lint \
	package-lint-fix \
	package-build \
	package-dev \
	package-test \
	package-check \
	package-go-lint \
	package-go-lint-fix \
	package-go-build \
	package-go-dev \
	package-go-test \
	package-go-check \

package-dev: package-go-dev

package-lint: package-go-lint

package-lint-fix: package-go-lint-fix

package-build: package-go-build

package-test: package-go-test

package-check: package-go-check

package-go-dev:

package-go-lint:
	golangci-lint run ./...

package-go-lint-fix:
	golangci-lint run --fix ./...

package-go-build:
	mkdir -p bin
	go build -o bin/terraform-provider-abstract ./src/main/go/

package-go-test:
	TF_ACC=1 go test ./src/test/go/... -v

package-go-check: package-go-lint package-go-test
