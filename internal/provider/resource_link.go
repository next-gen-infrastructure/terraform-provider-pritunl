package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/next-gen-infrastructure/terraform-provider-pritunl/internal/pritunl"
)

func resourceLink() *schema.Resource {
	return &schema.Resource{
		Description: "The link resource allows managing information about a particular Pritunl link.",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resource, also acts as it's unique ID",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "site_to_site",
				Description: "",
			},
			"ipv6": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "",
			},
			"host_check": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "restart",
				Description: "",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"preferred_ike": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "aes128-sha256-x25519",
				Description: "",
			},
			"preferred_esp": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "aes128gcm128-x25519",
				Description: "",
			},
			"force_preferred": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "",
			},
		},
		CreateContext: resourceCreateLink,
		ReadContext:   resourceReadLink,
		UpdateContext: resourceUpdateLink,
		DeleteContext: resourceDeleteLink,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// Uses for importing
func resourceReadLink(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(pritunl.Client)

	link, err := apiClient.GetLink(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", link.Name)

	return nil
}

func resourceDeleteLink(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(pritunl.Client)

	err := apiClient.DeleteLink(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func resourceUpdateLink(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(pritunl.Client)

	link, err := apiClient.GetLink(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange("name") {
		link.Name = d.Get("name").(string)

		err = apiClient.UpdateLink(d.Id(), link)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceCreateLink(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(pritunl.Client)

	linkData := pritunl.Link{
		Name:           d.Get("name").(string),
		Type:           d.Get("type").(string),
		Action:         d.Get("action").(string),
		PreferredIKE:   d.Get("preferred_ike").(string),
		PreferredESP:   d.Get("preferred_esp").(string),
		HostCheck:      d.Get("host_check").(bool),
		IPv6:           d.Get("ipv6").(bool),
		ForcePreferred: d.Get("force_preferred").(bool),
	}

	link, err := apiClient.CreateLink(linkData)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(link.ID)

	return nil
}
