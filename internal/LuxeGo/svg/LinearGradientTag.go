package svg

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// LinearGradient -
func LinearGradient(tags ...LuxeGo.Content) *LinearGradientTag {
	return &LinearGradientTag{ComponentSvgTag: &ComponentSvgTag{Name: "lineargradient", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type LinearGradientTag struct {
	*ComponentSvgTag
}

// X1 -
func (l *LinearGradientTag) X1(value string) *LinearGradientTag {
	l.AddAttribute("x1", value)
	return l
}

// X2 -
func (l *LinearGradientTag) X2(value string) *LinearGradientTag {
	l.AddAttribute("x2", value)
	return l
}

// Y1 -
func (l *LinearGradientTag) Y1(value string) *LinearGradientTag {
	l.AddAttribute("y1", value)
	return l
}

// Y2 -
func (l *LinearGradientTag) Y2(value string) *LinearGradientTag {
	l.AddAttribute("y2", value)
	return l
}

// GradientUnits -
func (l *LinearGradientTag) GradientUnits(value string) *LinearGradientTag {
	l.AddAttribute("gradientunits", value)
	return l
}

// GradientTransform -
func (l *LinearGradientTag) GradientTransform(value string) *LinearGradientTag {
	l.AddAttribute("gradienttransform", value)
	return l
}

// SpreadMethod -
func (l *LinearGradientTag) SpreadMethod(value string) *LinearGradientTag {
	l.AddAttribute("spreadmethod", value)
	return l
}
