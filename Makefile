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
	pkg-inf-terraform-provider-abstract-lint \
	pkg-inf-terraform-provider-abstract-lint-fix \
	pkg-inf-terraform-provider-abstract-build \
	pkg-inf-terraform-provider-abstract-dev \
	pkg-inf-terraform-provider-abstract-test \
	pkg-inf-terraform-provider-abstract-check \

package-dev: package-go-dev

package-lint: package-go-lint

package-lint-fix: package-go-lint-fix

package-build: package-go-build

package-test: package-go-test

package-check: package-go-check

package-go-dev:
	$(MAKE) -C packages/pkg-inf-terraform-provider-abstract package-go-dev

package-go-lint:
	$(MAKE) -C packages/pkg-inf-terraform-provider-abstract package-go-lint

package-go-lint-fix:
	$(MAKE) -C packages/pkg-inf-terraform-provider-abstract package-go-lint-fix

package-go-build:
	$(MAKE) -C packages/pkg-inf-terraform-provider-abstract package-go-build

package-go-test:
	$(MAKE) -C packages/pkg-inf-terraform-provider-abstract package-go-test

package-go-check:
	$(MAKE) -C packages/pkg-inf-terraform-provider-abstract package-go-check

pkg-inf-terraform-provider-abstract-dev:
	$(MAKE) -C packages/pkg-inf-terraform-provider-abstract package-dev

pkg-inf-terraform-provider-abstract-lint:
	$(MAKE) -C packages/pkg-inf-terraform-provider-abstract package-lint

pkg-inf-terraform-provider-abstract-lint-fix:
	$(MAKE) -C packages/pkg-inf-terraform-provider-abstract package-lint-fix

pkg-inf-terraform-provider-abstract-build:
	$(MAKE) -C packages/pkg-inf-terraform-provider-abstract package-build

pkg-inf-terraform-provider-abstract-test:
	$(MAKE) -C packages/pkg-inf-terraform-provider-abstract package-test

pkg-inf-terraform-provider-abstract-check:
	$(MAKE) -C packages/pkg-inf-terraform-provider-abstract package-check
