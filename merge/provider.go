package merge

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func NewProvider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"merge_merge": dataSourceMerge(),
		},
	}
}
