package postman

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	postmanSDK "github.com/jonnydgreen/terraform-provider-postman/client/postman"
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
				"api_key": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("POSTMAN_API_KEY", nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"postman_workspace": dataSourceWorkspace(),
				"postman_coffees":   dataSourceCoffees(),
				"hashicups_order":   dataSourceOrder(),
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
		apiKey := d.Get("api_key").(string)

		// Warning or errors can be collected in a slice type
		var diags diag.Diagnostics

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Some debug info",
			Detail:   "Some debug info",
		})

		if apiKey == "" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create Postman API client",
				Detail:   "No API Key specified for Postman API client. Please provider via environment variable (POSTMAN_API_KEY) or Provider variable (api_key).",
			})
			return nil, diags
		}

		// // Example ApiKey provider
		// // See: https://swagger.io/docs/specification/authentication/api-keys/
		// apiKeyProvider, apiKeyProviderErr := securityprovider.NewSecurityProviderApiKey("header", "x-api-key", apiKey)
		// if apiKeyProviderErr != nil {
		// 	diags = append(diags, diag.Diagnostic{
		// 		Severity: diag.Error,
		// 		Summary:  "Unable to setup the Postman API Key Provider",
		// 		Detail:   "Unable to auth API Key for authenticated Postman API client.",
		// 	})
		// 	return nil, diags
		// }

		configuration := postmanSDK.NewConfiguration()
		// TODO: This works for now
		configuration.AddDefaultHeader("x-api-key", apiKey)
		apiClient := postmanSDK.NewAPIClient(configuration)
		return apiClient, diags
	}
}
