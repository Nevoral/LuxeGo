package svg

import (
	"strings"

	base "github.com/Nevoral/LuxeGo/internal/BaseLayerDefinition"
	"github.com/Nevoral/LuxeGo/pkg/SpecGenzo/model"
)

func stringValue(value string) base.Content {
	if strings.HasPrefix(value, "<!--") {
		return base.NewTextNode(
			model.CommentType,
			model.SVG,
			value[4:],
		)
	}
	return base.NewTextNode(
		model.TextContentType,
		model.SVG,
		value,
	)
}
