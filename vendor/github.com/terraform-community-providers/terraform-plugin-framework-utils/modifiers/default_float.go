package modifiers

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DefaultFloat(def float64) planmodifier.Float64 {
	return defaultFloatModifier{Default: &def}
}

func NullableFloat() planmodifier.Float64 {
	return defaultFloatModifier{}
}

type defaultFloatModifier struct {
	Default *float64
}

func (m defaultFloatModifier) String() string {
	if m.Default == nil {
		return "null"
	}

	return fmt.Sprintf("%f", *m.Default)
}

func (m defaultFloatModifier) Description(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%s`", m)
}

func (m defaultFloatModifier) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%s`", m)
}

func (m defaultFloatModifier) PlanModifyFloat64(ctx context.Context, req planmodifier.Float64Request, resp *planmodifier.Float64Response) {
	if !req.ConfigValue.IsNull() {
		return
	}

	if m.Default == nil {
		resp.PlanValue = types.Float64Null()
	} else {
		resp.PlanValue = types.Float64Value(*m.Default)
	}
}
