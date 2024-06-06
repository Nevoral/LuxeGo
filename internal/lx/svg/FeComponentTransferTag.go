package svg

import (
	"LuxeGo/internal/lx"
)

// FeComponentTransfer -
func FeComponentTransfer() *FeComponentTransferTag {
	return &FeComponentTransferTag{ComponentSvgTag: &ComponentSvgTag{Name: "fecomponenttransfer", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeComponentTransferTag struct {
	*ComponentSvgTag
}

// In -
func (f *FeComponentTransferTag) In(value string) *FeComponentTransferTag {
	f.AddAttribute("in", value)
	return f
}
