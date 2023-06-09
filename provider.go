package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Provider function returns terraform provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"roztta_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ROZTTA_TOKEN", nil),
				Description: "The Roztta token for API operations. You can retrieve this from the Roztta Console.",
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringLenBetween(10, 20),
				),
			},
			"gitlab_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ROZTTA_GITLAB_TOKEN", nil),
				Description: "The GitLab token for API operations. You can retrieve this from the GitLab Console.",
			},
			"bitbucket_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ROZTTA_BITBUCKET_TOKEN", nil),
				Description: "The Bitbucket token for API operations. You can retrieve this from the Bitbucket Console.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"roztta_hello_world": resourceHelloWorld(),
			"roztta_gitlab_var":  resourceGitlabVar(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"roztta_gitlab_var": dataSourceGitlabVar(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	return &ProviderConfig{
		RozttaToken: d.Get("roztta_token").(string),
		GitlabToken: d.Get("gitlab_token").(string),
	}, nil
}
