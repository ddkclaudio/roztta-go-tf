package main

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGitlabVar() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGitlabVarRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGitlabVarRead(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	// TODO: implement the logic to read a GitLab variable

	// Fetch the value of the GitLab variable. Here, you'd typically make an API call to GitLab to get the variable's value.
	value := "some value" // replace this with the actual value

	// Set the ID and data of the data source.
	d.SetId(time.Now().UTC().String())
	d.Set("name", name)
	d.Set("value", value)

	return nil
}
