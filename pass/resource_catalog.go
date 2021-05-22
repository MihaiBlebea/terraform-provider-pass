package pass

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	cat "github.com/MihaiBlebea/go-pass-client/resource/catalog"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCatalog() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCatalogCreate,
		ReadContext:   resourceCatalogRead,
		UpdateContext: resourceCatalogUpdate,
		DeleteContext: resourceCatalogDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Optional: false,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Optional: false,
			},
		},
	}
}

func resourceCatalogCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client, ok := m.(cat.Service)
	if ok == false {
		return diag.FromErr(errors.New("Could not cast client"))
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req := cat.CreateCatalogRequest{
		Name:     d.Get("name").(string),
		Category: d.Get("catgory").(string),
	}
	id, err := client.CreateCatalog(req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(id))

	// items := d.Get("items").([]interface{})
	// ois := []hc.OrderItem{}

	// for _, item := range items {
	// 	i := item.(map[string]interface{})

	// 	co := i["coffee"].([]interface{})[0]
	// 	coffee := co.(map[string]interface{})

	// 	oi := hc.OrderItem{
	// 		Coffee: hc.Coffee{
	// 			ID: coffee["id"].(int),
	// 		},
	// 		Quantity: i["quantity"].(int),
	// 	}

	// 	ois = append(ois, oi)
	// }

	// o, err := c.CreateOrder(ois)
	// if err != nil {
	// 	return diag.FromErr(err)
	// }

	// d.SetId(strconv.Itoa(o.ID))

	resourceCatalogRead(ctx, d, m)

	return diags
}

func resourceCatalogRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client, ok := m.(cat.Service)
	if ok == false {
		return diag.FromErr(errors.New("Could not cast client"))
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	catalogID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	c, err := client.GetCatalog(catalogID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("catalog", c)

	return diags
}

func resourceCatalogUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// client, ok := m.(cat.Service)
	// if ok == false {
	// 	return diag.FromErr(errors.New("Could not cast client"))
	// }

	// catalogID, err := strconv.Atoi(d.Id())
	// if err != nil {
	// 	return diag.FromErr(err)
	// }

	if d.HasChange("catalog") {
		c := d.Get("catalog")
		fmt.Println(c)
		// ois := []hc.OrderItem{}

		// for _, item := range items {
		// 	i := item.(map[string]interface{})

		// 	co := i["coffee"].([]interface{})[0]
		// 	coffee := co.(map[string]interface{})

		// 	oi := hc.OrderItem{
		// 		Coffee: hc.Coffee{
		// 			ID: coffee["id"].(int),
		// 		},
		// 		Quantity: i["quantity"].(int),
		// 	}
		// 	ois = append(ois, oi)
		// }

		// req := cat.UpdateCatalogRequest{Name: c.Name, Category: c.Category}
		// _, err := client.UpdateCatalog(catalogID, ois)
		// if err != nil {
		// 	return diag.FromErr(err)
		// }

		// d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceCatalogRead(ctx, d, m)
}

func resourceCatalogDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client, ok := m.(cat.Service)
	if ok == false {
		return diag.FromErr(errors.New("Could not cast client"))
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	catalogID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.DeleteCatalog(catalogID)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
