package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider function returns terraform provider
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"roztta_hello_world": resourceHelloWorld(),
			"roztta_gitlab_var":  resourceGitlabVar(),
		},
	}
}
