package provider

import (
	"context"
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/next-gen-infrastructure/terraform-provider-pritunl/internal/pritunl"
)

func dataSourceLocation() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get information about the Pritunl location.",
		ReadContext: dataSourceLocationRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"link_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceLocationRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	linkId := d.Get("link_id").(string)

	filterFunction := func(location pritunl.Location) bool {
		return location.Name == name
	}

	location, err := filterLocations(meta, filterFunction, linkId)
	if err != nil {
		return diag.Errorf("could not find location with a name %+v. Previous error message: %v", location, err)
	}

	d.SetId(location.ID)
	_ = d.Set("name", location.Name)
	_ = d.Set("link_id", location.LinkId)

	return nil
}

func filterLocations(meta interface{}, test func(location pritunl.Location) bool, linkId string) (pritunl.Location, error) {
	apiClient := meta.(pritunl.Client)

	locations, err := apiClient.GetLocations(linkId)

	if err != nil {
		return pritunl.Location{}, err
	}

	for _, dir := range locations {
		if test(dir) {
			return dir, nil
		}
	}

	return pritunl.Location{}, errors.New("could not find a link with specified parameters")
}
