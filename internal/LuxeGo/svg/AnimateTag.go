package svg

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Animate -
func Animate() *AnimateTag {
	return &AnimateTag{ComponentSvgTag: &ComponentSvgTag{Name: "animate", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type AnimateTag struct {
	*ComponentSvgTag
}

// Accumulate -
func (a *AnimateTag) Accumulate(value string) *AnimateTag {
	a.AddAttribute("accumulate", value)
	return a
}

// Remove -
func (a *AnimateTag) Remove(value string) *AnimateTag {
	a.AddAttribute("remove", value)
	return a
}

// Additive -
func (a *AnimateTag) Additive(value string) *AnimateTag {
	a.AddAttribute("additive", value)
	return a
}

// AttributeName -
func (a *AnimateTag) AttributeName(value string) *AnimateTag {
	a.AddAttribute("attributename", value)
	return a
}

// AttributeType -
func (a *AnimateTag) AttributeType(value string) *AnimateTag {
	a.AddAttribute("attributetype", value)
	return a
}

// Begin -
func (a *AnimateTag) Begin(value string) *AnimateTag {
	a.AddAttribute("begin", value)
	return a
}

// By -
func (a *AnimateTag) By(value string) *AnimateTag {
	a.AddAttribute("by", value)
	return a
}

// CalcMode -
func (a *AnimateTag) CalcMode(value string) *AnimateTag {
	a.AddAttribute("calcmode", value)
	return a
}

// Dur -
func (a *AnimateTag) Dur(value string) *AnimateTag {
	a.AddAttribute("dur", value)
	return a
}

// End -
func (a *AnimateTag) End(value string) *AnimateTag {
	a.AddAttribute("end", value)
	return a
}

// Fill -
func (a *AnimateTag) Fill(value string) *AnimateTag {
	a.AddAttribute("fill", value)
	return a
}

// From -
func (a *AnimateTag) From(value string) *AnimateTag {
	a.AddAttribute("from", value)
	return a
}

// KeySplines -
func (a *AnimateTag) KeySplines(value string) *AnimateTag {
	a.AddAttribute("keysplines", value)
	return a
}

// KeyTimes -
func (a *AnimateTag) KeyTimes(value string) *AnimateTag {
	a.AddAttribute("keytimes", value)
	return a
}

// Max -
func (a *AnimateTag) Max(value string) *AnimateTag {
	a.AddAttribute("max", value)
	return a
}

// Min -
func (a *AnimateTag) Min(value string) *AnimateTag {
	a.AddAttribute("min", value)
	return a
}

// RepeatCount -
func (a *AnimateTag) RepeatCount(value string) *AnimateTag {
	a.AddAttribute("repeatcount", value)
	return a
}

// RepeatDur -
func (a *AnimateTag) RepeatDur(value string) *AnimateTag {
	a.AddAttribute("repeatdur", value)
	return a
}

// Restart -
func (a *AnimateTag) Restart(value string) *AnimateTag {
	a.AddAttribute("restart", value)
	return a
}

// To -
func (a *AnimateTag) To(value string) *AnimateTag {
	a.AddAttribute("to", value)
	return a
}

// Values -
func (a *AnimateTag) Values(value string) *AnimateTag {
	a.AddAttribute("values", value)
	return a
}
