package svg

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// FeComponentTransfer -
func FeComponentTransfer() *FeComponentTransferTag {
	return &FeComponentTransferTag{ComponentSvgTag: &ComponentSvgTag{Name: "fecomponenttransfer", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type FeComponentTransferTag struct {
	*ComponentSvgTag
}

// In -
func (f *FeComponentTransferTag) In(value string) *FeComponentTransferTag {
	f.AddAttribute("in", value)
	return f
}
