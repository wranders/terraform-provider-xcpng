package xcpng

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Default:     "",
				Description: "The XenAPI endpoint",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Default:     "",
				Description: "The user that will be authenticated",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Default:     "",
				Description: "The password to authenticate the user",
			},
			"insecureSkipVerify": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Skip TLS verification when connecting to server",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap:   map[string]*schema.Resource{},
		ConfigureFunc:  providerConfigureFunc,
	}
}

func providerConfigureFunc(rd *schema.ResourceData) (interface{}, error) {
	return nil, nil
}
