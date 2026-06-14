package runner_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/ii1818nn/terraform-provider-abstract/src/main/go/provider"
)

var providerFactories = map[string]func() (*schema.Provider, error){
	"abstract": func() (*schema.Provider, error) {
		return provider.Provider(), nil
	},
}

func TestAbstractResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "abstract_resource" "test" {
						lifecycle_commands {
							create = "echo created"
							read   = "exit 0"
							update = "exit 0"
							delete = "echo deleted"
						}
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("abstract_resource.test", "exit_code", "0"),
					resource.TestCheckResourceAttr("abstract_resource.test", "output", "created"),
				),
			},
		},
	})
}

func TestAbstractResource_withVariables(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "abstract_resource" "test" {
						lifecycle_commands {
							create = "echo $GREETING"
							read   = "exit 0"
							update = "exit 0"
							delete = "exit 0"
						}
						variables = {
							GREETING = "hello"
						}
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("abstract_resource.test", "output", "hello"),
				),
			},
		},
	})
}

func TestAbstractResource_withSecrets(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "abstract_resource" "test" {
						lifecycle_commands {
							create = "echo $SECRET_TOKEN"
							read   = "exit 0"
							update = "exit 0"
							delete = "exit 0"
						}
						secrets = {
							SECRET_TOKEN = "s3cr3t"
						}
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("abstract_resource.test", "output", "s3cr3t"),
				),
			},
		},
	})
}

func TestAbstractResource_secretOverridesVariable(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "abstract_resource" "test" {
						lifecycle_commands {
							create = "echo $VAR_KEY $SECRET_KEY"
							read   = "exit 0"
							update = "exit 0"
							delete = "exit 0"
						}
						variables = {
							VAR_KEY = "from_variable"
						}
						secrets = {
							SECRET_KEY = "from_secret"
						}
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("abstract_resource.test", "output", "from_variable from_secret"),
				),
			},
		},
	})
}

func TestAbstractResource_withTriggers(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "abstract_resource" "test" {
						lifecycle_commands {
							create = "echo v1"
							read   = "exit 0"
							update = "exit 0"
							delete = "exit 0"
						}
						triggers = {
							version = "1"
						}
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("abstract_resource.test", "output", "v1"),
					resource.TestCheckResourceAttr("abstract_resource.test", "triggers.version", "1"),
				),
			},
			{
				Config: `
					resource "abstract_resource" "test" {
						lifecycle_commands {
							create = "echo v2"
							read   = "exit 0"
							update = "exit 0"
							delete = "exit 0"
						}
						triggers = {
							version = "2"
						}
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("abstract_resource.test", "output", "v2"),
					resource.TestCheckResourceAttr("abstract_resource.test", "triggers.version", "2"),
				),
			},
		},
	})
}
