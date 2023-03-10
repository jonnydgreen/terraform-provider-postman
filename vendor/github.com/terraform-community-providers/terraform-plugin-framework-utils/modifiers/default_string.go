package modifiers

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DefaultString(def string) planmodifier.String {
	return defaultStringModifier{Default: &def}
}

func NullableString() planmodifier.String {
	return defaultStringModifier{}
}

type defaultStringModifier struct {
	Default *string
}

func (m defaultStringModifier) String() string {
	if m.Default == nil {
		return "null"
	}

	return *m.Default
}

func (m defaultStringModifier) Description(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%s`", m)
}

func (m defaultStringModifier) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%s`", m)
}

func (m defaultStringModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if !req.ConfigValue.IsNull() {
		return
	}

	if m.Default == nil {
		resp.PlanValue = types.StringNull()
	} else {
		resp.PlanValue = types.StringValue(*m.Default)
	}
}
