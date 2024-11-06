package html

import (
	base "github.com/Nevoral/LuxeGo/internal/BaseLayerDefinition"
	"github.com/Nevoral/LuxeGo/pkg/SpecGenzo/model"
)

type Namespace int

const (
	HTML Namespace = iota
	SVG
	MATH
)

func (n Namespace) String() string {
	return [...]string{"html", "svg", "math"}[n]
}

// TODO: Public and System can be enum or string with own location also there can be specified Entities for prefixes atc. Not handled yet
type DoctypeAttributes struct {
	Namespace Namespace
	Public    string
	System    string
}

func DOCTYPE(attributes *DoctypeAttributes) base.Content {
	return base.NewDoctypeNode(
		model.Namespace(attributes.Namespace),
	)
}
