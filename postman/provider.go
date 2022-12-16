package postman

import (
	"context"

	"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"username": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("HASHICUPS_USERNAME", nil),
				},
				"password": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					Sensitive:   true,
					DefaultFunc: schema.EnvDefaultFunc("HASHICUPS_PASSWORD", nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"postman_coffees": dataSourceCoffees(),
				"hashicups_order": dataSourceOrder(),
			},
			ResourcesMap: map[string]*schema.Resource{},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

// TODO: we probably want this
// type apiClient struct {
// 	// Add whatever fields, client or connection info, etc. here
// 	// you would need to setup to communicate with the upstream
// 	// API.
// }

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
	// return func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
	// 	// Setup a User-Agent for your API client (replace the provider name for yours):
	// 	// userAgent := p.UserAgent("terraform-provider-postman", version)
	// 	// TODO: myClient.UserAgent = userAgent

	// 	return &apiClient{}, nil
	// }
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		username := d.Get("username").(string)
		password := d.Get("password").(string)

		// Warning or errors can be collected in a slice type
		var diags diag.Diagnostics

		if (username != "") && (password != "") {
			c, err := hashicups.NewClient(nil, &username, &password)
			if err != nil {
				return nil, diag.FromErr(err)
			}

			return c, diags
		}

		c, err := hashicups.NewClient(nil, nil, nil)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return c, diags
	}
}
