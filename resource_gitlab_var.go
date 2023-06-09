package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGitlabVar() *schema.Resource {
	return &schema.Resource{
		Create: resourceGitlabVarCreate,
		Read:   resourceGitlabVarRead,
		Update: resourceGitlabVarUpdate,
		Delete: resourceGitlabVarDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceGitlabVarCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	value := d.Get("value").(string)
	// TODO: implement the logic to create a new GitLab variable

	// Set the ID and data of the resource.
	// Here we simply set the ID to the variable name and store the value.
	// In a real provider, you'd likely set the ID to a unique identifier returned by the GitLab API,
	// and store additional data if needed.
	d.SetId(name)
	d.Set("value", value)
	return nil
}

func resourceGitlabVarRead(d *schema.ResourceData, m interface{}) error {
	// TODO: implement the logic to read a GitLab variable
	return nil
}

func resourceGitlabVarUpdate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	value := d.Get("value").(string)
	// TODO: implement the logic to update a GitLab variable

	// Update the data of the resource.
	d.Set("name", name)
	d.Set("value", value)
	return nil
}

func resourceGitlabVarDelete(d *schema.ResourceData, m interface{}) error {
	// TODO: implement the logic to delete a GitLab variable
	d.SetId("")
	return nil
}
