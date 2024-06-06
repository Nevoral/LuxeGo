package svg

import (
	"LuxeGo/internal/lx"
)

// AnimateTransform -
func AnimateTransform() *AnimateTransformTag {
	return &AnimateTransformTag{ComponentSvgTag: &ComponentSvgTag{Name: "animatetransform", Attributes: &lx.Attributes{}, Children: nil}}
}

type AnimateTransformTag struct {
	*ComponentSvgTag
}

// Accumulate -
func (a *AnimateTransformTag) Accumulate(value string) *AnimateTransformTag {
	a.AddAttribute("accumulate", value)
	return a
}

// Additive -
func (a *AnimateTransformTag) Additive(value string) *AnimateTransformTag {
	a.AddAttribute("additive", value)
	return a
}

// AttributeName -
func (a *AnimateTransformTag) AttributeName(value string) *AnimateTransformTag {
	a.AddAttribute("attributename", value)
	return a
}

// AttributeType -
func (a *AnimateTransformTag) AttributeType(value string) *AnimateTransformTag {
	a.AddAttribute("attributetype", value)
	return a
}

// Begin -
func (a *AnimateTransformTag) Begin(value string) *AnimateTransformTag {
	a.AddAttribute("begin", value)
	return a
}

// By -
func (a *AnimateTransformTag) By(value string) *AnimateTransformTag {
	a.AddAttribute("by", value)
	return a
}

// CalcMode -
func (a *AnimateTransformTag) CalcMode(value string) *AnimateTransformTag {
	a.AddAttribute("calcmode", value)
	return a
}

// Dur -
func (a *AnimateTransformTag) Dur(value string) *AnimateTransformTag {
	a.AddAttribute("dur", value)
	return a
}

// End -
func (a *AnimateTransformTag) End(value string) *AnimateTransformTag {
	a.AddAttribute("end", value)
	return a
}

// Fill -
func (a *AnimateTransformTag) Fill(value string) *AnimateTransformTag {
	a.AddAttribute("fill", value)
	return a
}

// From -
func (a *AnimateTransformTag) From(value string) *AnimateTransformTag {
	a.AddAttribute("from", value)
	return a
}

// KeySplines -
func (a *AnimateTransformTag) KeySplines(value string) *AnimateTransformTag {
	a.AddAttribute("keysplines", value)
	return a
}

// KeyTimes -
func (a *AnimateTransformTag) KeyTimes(value string) *AnimateTransformTag {
	a.AddAttribute("keytimes", value)
	return a
}

// Max -
func (a *AnimateTransformTag) Max(value string) *AnimateTransformTag {
	a.AddAttribute("max", value)
	return a
}

// Min -
func (a *AnimateTransformTag) Min(value string) *AnimateTransformTag {
	a.AddAttribute("min", value)
	return a
}

// RepeatCount -
func (a *AnimateTransformTag) RepeatCount(value string) *AnimateTransformTag {
	a.AddAttribute("repeatcount", value)
	return a
}

// RepeatDur -
func (a *AnimateTransformTag) RepeatDur(value string) *AnimateTransformTag {
	a.AddAttribute("repeatdur", value)
	return a
}

// Restart -
func (a *AnimateTransformTag) Restart(value string) *AnimateTransformTag {
	a.AddAttribute("restart", value)
	return a
}

// To -
func (a *AnimateTransformTag) To(value string) *AnimateTransformTag {
	a.AddAttribute("to", value)
	return a
}

// Type -
func (a *AnimateTransformTag) Type(value string) *AnimateTransformTag {
	a.AddAttribute("type", value)
	return a
}

// Values -
func (a *AnimateTransformTag) Values(value string) *AnimateTransformTag {
	a.AddAttribute("values", value)
	return a
}
