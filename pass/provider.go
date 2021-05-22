package pass

import (
	"context"

	c "github.com/MihaiBlebea/go-pass-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: false,
				// DefaultFunc: schema.EnvDefaultFunc("PASS_HOST", nil),
			},
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: false,
				// DefaultFunc: schema.EnvDefaultFunc("PASS_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"pass_catalog": resourceCatalog(),
		},
		// DataSourcesMap: map[string]*schema.Resource{
		// 	"pass_catalog": dataSourceCatalog(),
		// },
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	host := d.Get("host").(string)
	token := d.Get("token").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if host == "" || token == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Pass client",
			Detail:   "Unable to create Pass client, missing token or host",
		})

		return nil, diags
	}

	return c.New(token), diags
}

// func configureFunc() func(*schema.ResourceData) (interface{}, error) {
// 	return func(d *schema.ResourceData) (interface{}, error) {
// 		client := todoistRest.NewClient(d.Get("api_key").(string))
// 		return client, nil
// 	}
// }

// func providerConfigure(d *schema.ResourceData) (interface{}, error) {
// 	token := d.Get("token").(string)
// 	// host := d.Get("host").(int)

// 	return c.New(token), nil
// }
