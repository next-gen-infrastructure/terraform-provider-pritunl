package provider

import (
	"context"
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/next-gen-infrastructure/terraform-provider-pritunl/internal/pritunl"
)

func dataSourceLink() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get information about the Pritunl link.",
		ReadContext: dataSourceLinkRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"host_check": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"preferred_ike": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"preferred_esp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"force_preferred": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceLinkRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name")
	filterFunction := func(link pritunl.Link) bool {
		return link.Name == name
	}

	link, err := filterLinks(meta, filterFunction)
	if err != nil {
		return diag.Errorf("could not find link with a name %+v. Previous error message: %v", link, err)
	}

	d.SetId(link.ID)
	_ = d.Set("name", link.Name)
	_ = d.Set("type", link.Type)
	_ = d.Set("ipv6", link.IPv6)
	_ = d.Set("host_check", link.Name)
	_ = d.Set("action", link.Action)
	_ = d.Set("preferred_ike", link.PreferredIKE)
	_ = d.Set("preferred_esp", link.PreferredESP)
	_ = d.Set("force_preferred", link.ForcePreferred)

	return nil
}

func filterLinks(meta interface{}, test func(link pritunl.Link) bool) (pritunl.Link, error) {
	apiClient := meta.(pritunl.Client)

	links, err := apiClient.GetLinks()

	if err != nil {
		return pritunl.Link{}, err
	}

	for _, dir := range links {
		if test(dir) {
			return dir, nil
		}
	}

	return pritunl.Link{}, errors.New("could not find a link with specified parameters")
}
