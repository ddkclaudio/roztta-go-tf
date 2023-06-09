package main

import (
	"fmt"

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

	config := m.(*ProviderConfig)

	fmt.Printf("Roztta Token: %s\n", config.RozttaToken)
	fmt.Printf("Gitlab Token: %s\n", config.GitlabToken)

	// Define the URL
	url := "https://cms.dev.mytaverse.com/auth/signIn"

	// Define the payload
	payload := map[string]string{
		"email":    "claudio@mytaverse.com",
		"password": "senha-segura",
	}

	// Instantiate the HTTP client
	httpClient := NewHTTPClient()

	// Call the Post method
	response, err := httpClient.Post(url, payload)
	if err != nil {
		return fmt.Errorf("error making POST request: %s", err)
	}

	fmt.Println("Response: ", response)

	// Set the ID and data of the resource.
	// Here we simply set the ID to the variable name and store the value.
	// In a real provider, you'd likely set the ID to a unique identifier returned by the GitLab API,
	// and store additional data if needed.
	d.SetId(name)
	d.Set("value", value)
	return nil
}

func resourceGitlabVarRead(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	value := d.Get("value").(string)

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Value: %s\n", value)

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
