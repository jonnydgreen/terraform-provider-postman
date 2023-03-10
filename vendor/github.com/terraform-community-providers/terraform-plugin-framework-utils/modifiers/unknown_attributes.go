package modifiers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func UnknownAttributesOnUnknown() planmodifier.Object {
	return unknownAttributesOnUnknownModifier{}
}

type unknownAttributesOnUnknownModifier struct{}

func (m unknownAttributesOnUnknownModifier) Description(ctx context.Context) string {
	return "If value is unknown, defaults to an object with all attributes set to unknown."
}

func (m unknownAttributesOnUnknownModifier) MarkdownDescription(ctx context.Context) string {
	return "If value is unknown, defaults to an object with all attributes set to unknown."
}

func (m unknownAttributesOnUnknownModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	if !req.PlanValue.IsUnknown() {
		return
	}

	var object types.Object

	diags := tfsdk.ValueAs(ctx, req.PlanValue, &object)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	attrs := make(map[string]attr.Value, len(object.AttributeTypes(ctx)))

	for name, attrType := range object.AttributeTypes(ctx) {
		attrValue, err := attrType.ValueFromTerraform(ctx, tftypes.NewValue(attrType.TerraformType(ctx), tftypes.UnknownValue))

		if err != nil {
			resp.Diagnostics.AddAttributeError(
				req.Path.AtName(name),
				"Attribute Plan Modification Error",
				"While creating unknown values for object attributes, an unexpected error occurred. Please report the following to the provider developers.\n\n"+
					"Error: "+err.Error(),
			)
		}

		attrs[name] = attrValue
	}

	resp.PlanValue = types.ObjectValueMust(object.AttributeTypes(ctx), attrs)
}
