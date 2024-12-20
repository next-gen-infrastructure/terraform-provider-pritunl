package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/next-gen-infrastructure/terraform-provider-pritunl/internal/pritunl"
)

func resourceHost() *schema.Resource {
	return &schema.Resource{
		Description: "The host resource allows managing information about a particular Pritunl location host.",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resource",
			},
			"link_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the link to create host at",
			},
			"location_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the location to create host at",
			},
			"uri": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "URI of the host",
			},
		},
		CreateContext: resourceCreateHost,
		ReadContext:   resourceReadHost,
		UpdateContext: resourceUpdateHost,
		DeleteContext: resourceDeleteHost,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// Uses for importing
func resourceReadHost(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(pritunl.Client)

	host, err := apiClient.GetHost(
		d.Id(),
		d.Get("link_id").(string),
		d.Get("location_id").(string),
		d.Get("uri").(string),
	)
	if err != nil {
		return diag.FromErr(err)
	}

	_ = d.Set("name", host.Name)
	_ = d.Set("link_id", host.LinkID)
	_ = d.Set("location_id", host.LocationID)
	_ = d.Set("uri", host.URI)
	return nil
}

func resourceDeleteHost(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(pritunl.Client)

	err := apiClient.DeleteHost(d.Id(), d.Get("link_id").(string), d.Get("location_id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func resourceUpdateHost(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(pritunl.Client)

	host, err := apiClient.GetHost(
		d.Id(),
		d.Get("link_id").(string),
		d.Get("location_id").(string),
		d.Get("uri").(string),
	)
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange("name") {
		host.Name = d.Get("name").(string)

		err = apiClient.UpdateHost(d.Id(), host)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("uri") {
		host.URI = d.Get("uri").(string)

		err = apiClient.UpdateHost(d.Id(), host)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return resourceReadHost(ctx, d, meta)
}

func resourceCreateHost(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(pritunl.Client)

	hostData := pritunl.LocationHost{
		Name:       d.Get("name").(string),
		LinkID:     d.Get("link_id").(string),
		LocationID: d.Get("location_id").(string),
		URI:        d.Get("uri").(string),
	}

	host, err := apiClient.CreateHost(hostData)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(host.ID)

	return nil
}
