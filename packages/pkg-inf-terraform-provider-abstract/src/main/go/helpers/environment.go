package helpers

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GetEnvironment(d *schema.ResourceData) map[string]string {
	env := map[string]string{}
	if v, ok := d.GetOk("variables"); ok {
		for k, val := range v.(map[string]interface{}) {
			env[k] = val.(string)
		}
	}
	if v, ok := d.GetOk("secrets"); ok {
		for k, val := range v.(map[string]interface{}) {
			env[k] = val.(string)
		}
	}
	return env
}
