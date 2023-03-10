package modifiers

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DefaultBool(def bool) planmodifier.Bool {
	return defaultBoolModifier{Default: &def}
}

func NullableBool() planmodifier.Bool {
	return defaultBoolModifier{}
}

type defaultBoolModifier struct {
	Default *bool
}

func (m defaultBoolModifier) String() string {
	if m.Default == nil {
		return "null"
	}

	return fmt.Sprintf("%t", *m.Default)
}

func (m defaultBoolModifier) Description(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%s`", m)
}

func (m defaultBoolModifier) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%s`", m)
}

func (m defaultBoolModifier) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	if !req.ConfigValue.IsNull() {
		return
	}

	if m.Default == nil {
		resp.PlanValue = types.BoolNull()
	} else {
		resp.PlanValue = types.BoolValue(*m.Default)
	}
}
