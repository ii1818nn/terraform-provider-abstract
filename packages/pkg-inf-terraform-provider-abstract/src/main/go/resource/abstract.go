package resource

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/ii1818nn/pkg-inf-terraform-provider-abstract/src/main/go/helpers"
	"github.com/ii1818nn/pkg-inf-terraform-provider-abstract/src/main/go/runner"
)

func Abstract() *schema.Resource {
	return &schema.Resource{
		Create: create,
		Read:   read,
		Update: update,
		Delete: delete,

		Schema: map[string]*schema.Schema{
			"lifecycle_commands": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"create": {
							Type:     schema.TypeString,
							Required: true,
						},
						"read": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "exit 0",
						},
						"update": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "exit 0",
						},
						"delete": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
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

func getCommand(d *schema.ResourceData, key string) string {
	cmds := d.Get("lifecycle_commands").([]interface{})
	return cmds[0].(map[string]interface{})[key].(string)
}

func create(d *schema.ResourceData, _ interface{}) error {
	res, err := runner.Shell(getCommand(d, "create"), helpers.GetEnvironment(d))
	if err != nil {
		return fmt.Errorf("create command failed: %w", err)
	}
	if res.ExitCode != 0 {
		return fmt.Errorf("create command exited with code %d: %s", res.ExitCode, res.Output)
	}

	d.SetId("abstract_resource")
	_ = d.Set("output", res.Output)
	_ = d.Set("exit_code", res.ExitCode)

	return nil
}

func read(d *schema.ResourceData, _ interface{}) error {
	res, err := runner.Shell(getCommand(d, "read"), helpers.GetEnvironment(d))
	if err != nil {
		return fmt.Errorf("read command failed: %w", err)
	}
	if res.ExitCode != 0 {
		d.SetId("")
		return nil
	}

	_ = d.Set("output", res.Output)
	_ = d.Set("exit_code", res.ExitCode)

	return nil
}

func update(d *schema.ResourceData, _ interface{}) error {
	res, err := runner.Shell(getCommand(d, "update"), helpers.GetEnvironment(d))
	if err != nil {
		return fmt.Errorf("update command failed: %w", err)
	}
	if res.ExitCode != 0 {
		return fmt.Errorf("update command exited with code %d: %s", res.ExitCode, res.Output)
	}

	_ = d.Set("output", res.Output)
	_ = d.Set("exit_code", res.ExitCode)

	return nil
}

func delete(d *schema.ResourceData, _ interface{}) error {
	res, err := runner.Shell(getCommand(d, "delete"), helpers.GetEnvironment(d))
	if err != nil {
		return fmt.Errorf("delete command failed: %w", err)
	}
	if res.ExitCode != 0 {
		return fmt.Errorf("delete command exited with code %d: %s", res.ExitCode, res.Output)
	}

	return nil
}
