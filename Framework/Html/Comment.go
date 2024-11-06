package html

import (
	"fmt"
	"strings"

	base "github.com/Nevoral/LuxeGo/internal/BaseLayerDefinition"
	"github.com/Nevoral/LuxeGo/pkg/SpecGenzo/model"
)

func Comment(comment string) string {
	return fmt.Sprintf("<!--%s", strings.TrimSpace(comment))
}

func stringValue(value string) base.Content {
	if strings.HasPrefix(value, "<!--") {
		return base.NewTextNode(
			model.CommentType,
			model.HTML,
			value[4:],
		)
	}
	return base.NewTextNode(
		model.TextContentType,
		model.HTML,
		value,
	)
}
