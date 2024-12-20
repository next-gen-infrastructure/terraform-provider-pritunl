package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/next-gen-infrastructure/terraform-provider-pritunl/internal/pritunl"
)

func resourceRoute() *schema.Resource {
	return &schema.Resource{
		Description: "The route resource allows managing information about a particular Pritunl location route.",
		Schema: map[string]*schema.Schema{
			"network": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The network of the resource",
			},
			"link_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the link to create route at",
			},
			"location_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the location to create route at",
			},
		},
		CreateContext: resourceCreateRoute,
		ReadContext:   resourceReadRoute,
		UpdateContext: resourceUpdateRoute,
		DeleteContext: resourceDeleteRoute,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// Uses for importing
func resourceReadRoute(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(pritunl.Client)

	route, err := apiClient.GetRoute(d.Id(), d.Get("link_id").(string), d.Get("location_id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	_ = d.Set("network", route.Network)
	_ = d.Set("link_id", route.LinkId)
	_ = d.Set("location_id", route.LocationId)

	return nil
}

func resourceDeleteRoute(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(pritunl.Client)

	err := apiClient.DeleteRoute(d.Id(), d.Get("link_id").(string), d.Get("location_id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func resourceUpdateRoute(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(pritunl.Client)

	route, err := apiClient.GetRoute(d.Id(), d.Get("link_id").(string), d.Get("location_id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange("network") {
		route.Network = d.Get("network").(string)

		err = apiClient.UpdateRoute(d.Id(), route)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceReadRoute(ctx, d, meta)
}

func resourceCreateRoute(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(pritunl.Client)

	routeData := pritunl.LocationRoute{
		Network:    d.Get("network").(string),
		LinkId:     d.Get("link_id").(string),
		LocationId: d.Get("location_id").(string),
	}

	route, err := apiClient.CreateRoute(routeData)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(route.ID)

	return nil
}
