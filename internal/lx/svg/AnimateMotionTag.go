package svg

import (
	"LuxeGo/internal/lx"
)

// AnimateMotion -
func AnimateMotion() *AnimateMotionTag {
	return &AnimateMotionTag{ComponentSvgTag: &ComponentSvgTag{Name: "animatemotion", Attributes: &lx.Attributes{}, Children: nil}}
}

type AnimateMotionTag struct {
	*ComponentSvgTag
}

// Accumulate -
func (a *AnimateMotionTag) Accumulate(value string) *AnimateMotionTag {
	a.AddAttribute("accumulate", value)
	return a
}

// Additive -
func (a *AnimateMotionTag) Additive(value string) *AnimateMotionTag {
	a.AddAttribute("additive", value)
	return a
}

// Begin -
func (a *AnimateMotionTag) Begin(value string) *AnimateMotionTag {
	a.AddAttribute("begin", value)
	return a
}

// By -
func (a *AnimateMotionTag) By(value string) *AnimateMotionTag {
	a.AddAttribute("by", value)
	return a
}

// CalcMode -
func (a *AnimateMotionTag) CalcMode(value string) *AnimateMotionTag {
	a.AddAttribute("calcmode", value)
	return a
}

// Dur -
func (a *AnimateMotionTag) Dur(value string) *AnimateMotionTag {
	a.AddAttribute("dur", value)
	return a
}

// End -
func (a *AnimateMotionTag) End(value string) *AnimateMotionTag {
	a.AddAttribute("end", value)
	return a
}

// Fill -
func (a *AnimateMotionTag) Fill(value string) *AnimateMotionTag {
	a.AddAttribute("fill", value)
	return a
}

// From -
func (a *AnimateMotionTag) From(value string) *AnimateMotionTag {
	a.AddAttribute("from", value)
	return a
}

// KeyPoints -
func (a *AnimateMotionTag) KeyPoints(value string) *AnimateMotionTag {
	a.AddAttribute("keypoints", value)
	return a
}

// KeySplines -
func (a *AnimateMotionTag) KeySplines(value string) *AnimateMotionTag {
	a.AddAttribute("keysplines", value)
	return a
}

// KeyTimes -
func (a *AnimateMotionTag) KeyTimes(value string) *AnimateMotionTag {
	a.AddAttribute("keytimes", value)
	return a
}

// Max -
func (a *AnimateMotionTag) Max(value string) *AnimateMotionTag {
	a.AddAttribute("max", value)
	return a
}

// Min -
func (a *AnimateMotionTag) Min(value string) *AnimateMotionTag {
	a.AddAttribute("min", value)
	return a
}

// Path -
func (a *AnimateMotionTag) Path(value string) *AnimateMotionTag {
	a.AddAttribute("path", value)
	return a
}

// RepeatCount -
func (a *AnimateMotionTag) RepeatCount(value string) *AnimateMotionTag {
	a.AddAttribute("repeatcount", value)
	return a
}

// RepeatDur -
func (a *AnimateMotionTag) RepeatDur(value string) *AnimateMotionTag {
	a.AddAttribute("repeatdur", value)
	return a
}

// Restart -
func (a *AnimateMotionTag) Restart(value string) *AnimateMotionTag {
	a.AddAttribute("restart", value)
	return a
}

// Rotate -
func (a *AnimateMotionTag) Rotate(value string) *AnimateMotionTag {
	a.AddAttribute("rotate", value)
	return a
}

// To -
func (a *AnimateMotionTag) To(value string) *AnimateMotionTag {
	a.AddAttribute("to", value)
	return a
}

// Values -
func (a *AnimateMotionTag) Values(value string) *AnimateMotionTag {
	a.AddAttribute("values", value)
	return a
}
