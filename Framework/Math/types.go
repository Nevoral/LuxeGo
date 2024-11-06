package math

import (
	"strings"

	base "github.com/Nevoral/LuxeGo/internal/BaseLayerDefinition"
	"github.com/Nevoral/LuxeGo/pkg/SpecGenzo/model"
)

func stringValue(value string) base.Content {
	if strings.HasPrefix(value, "<!--") {
		return base.NewTextNode(
			model.CommentType,
			model.MATH,
			value[4:],
		)
	}
	return base.NewTextNode(
		model.TextContentType,
		model.MATH,
		value,
	)
}
