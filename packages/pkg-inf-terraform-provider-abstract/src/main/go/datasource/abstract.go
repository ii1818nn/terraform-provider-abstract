package datasource

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/ii1818nn/pkg-inf-terraform-provider-abstract/src/main/go/helpers"
	"github.com/ii1818nn/pkg-inf-terraform-provider-abstract/src/main/go/runner"
)

func Abstract() *schema.Resource {
	return &schema.Resource{
		Read: read,

		Schema: map[string]*schema.Schema{
			"read_command": {
				Type:     schema.TypeString,
				Required: true,
			},
			"variables": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"secrets": {
				Type:      schema.TypeMap,
				Optional:  true,
				Sensitive: true,
				Elem:      &schema.Schema{Type: schema.TypeString},
			},
			"triggers": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"output": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exit_code": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func read(d *schema.ResourceData, _ interface{}) error {
	res, err := runner.Shell(d.Get("read_command").(string), helpers.GetEnvironment(d))
	if err != nil {
		return fmt.Errorf("read command failed: %w", err)
	}

	d.SetId("abstract_datasource")
	_ = d.Set("output", res.Output)
	_ = d.Set("exit_code", res.ExitCode)

	return nil
}
