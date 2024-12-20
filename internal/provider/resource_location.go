package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/next-gen-infrastructure/terraform-provider-pritunl/internal/pritunl"
)

func resourceLocation() *schema.Resource {
	return &schema.Resource{
		Description: "The location resource allows managing information about a particular Pritunl location.",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resource, also acts as it's unique ID",
			},
			"link_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the link to create location at",
			},
		},
		CreateContext: resourceCreateLocation,
		ReadContext:   resourceReadLocation,
		UpdateContext: resourceUpdateLocation,
		DeleteContext: resourceDeleteLocation,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// Uses for importing
func resourceReadLocation(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(pritunl.Client)

	location, err := apiClient.GetLocation(d.Id(), d.Get("link_id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	_ = d.Set("name", location.Name)
	_ = d.Set("link_id", location.LinkId)

	return nil
}

func resourceDeleteLocation(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(pritunl.Client)

	err := apiClient.DeleteLocation(d.Id(), d.Get("link_id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func resourceUpdateLocation(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(pritunl.Client)

	location, err := apiClient.GetLocation(d.Id(), d.Get("link_id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange("name") {
		location.Name = d.Get("name").(string)

		err = apiClient.UpdateLocation(d.Id(), location)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceReadLocation(ctx, d, meta)
}

func resourceCreateLocation(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(pritunl.Client)

	locationData := pritunl.Location{
		Name:   d.Get("name").(string),
		LinkId: d.Get("link_id").(string),
	}

	location, err := apiClient.CreateLocation(locationData)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(location.ID)

	return nil
}
