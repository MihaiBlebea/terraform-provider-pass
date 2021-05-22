package pass

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	c "github.com/MihaiBlebea/go-pass-client"
)

func dataSourceCatalog() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCatalogRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"coffee_id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"coffee_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"coffee_teaser": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"coffee_description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"coffee_price": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"coffee_image": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"quantity": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCatalogRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := c.New(d.Get("token").(string))

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	catalogID := d.Get("id").(int)

	_, err := client.GetCatalog(catalogID)
	if err != nil {
		return diag.FromErr(err)
	}

	// orderItems := flattenOrderItemsData(&order.Items)
	// if err := d.Set("items", orderItems); err != nil {
	// 	return diag.FromErr(err)
	// }

	d.SetId(strconv.Itoa(catalogID))

	return diags
}
