package runner_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAbstractDatasource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					data "abstract_datasource" "test" {
						read_command = "echo hello"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.abstract_datasource.test", "exit_code", "0"),
					resource.TestCheckResourceAttr("data.abstract_datasource.test", "output", "hello"),
				),
			},
		},
	})
}

func TestAbstractDatasource_withVariables(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					data "abstract_datasource" "test" {
						read_command = "echo $GREETING"
						variables = {
							GREETING = "hello"
						}
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.abstract_datasource.test", "output", "hello"),
				),
			},
		},
	})
}

func TestAbstractDatasource_withSecrets(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					data "abstract_datasource" "test" {
						read_command = "echo $SECRET_TOKEN"
						secrets = {
							SECRET_TOKEN = "s3cr3t"
						}
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.abstract_datasource.test", "output", "s3cr3t"),
				),
			},
		},
	})
}

func TestAbstractDatasource_variablesAndSecrets(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					data "abstract_datasource" "test" {
						read_command = "echo $VAR_KEY $SECRET_KEY"
						variables = {
							VAR_KEY = "from_variable"
						}
						secrets = {
							SECRET_KEY = "from_secret"
						}
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.abstract_datasource.test", "output", "from_variable from_secret"),
				),
			},
		},
	})
}
