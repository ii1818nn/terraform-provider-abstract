package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/ii1818nn/pkg-inf-terraform-provider-abstract/src/main/go/datasource"
	"github.com/ii1818nn/pkg-inf-terraform-provider-abstract/src/main/go/resource"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"abstract_resource": resource.Abstract(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"abstract_datasource": datasource.Abstract(),
		},
	}
}
