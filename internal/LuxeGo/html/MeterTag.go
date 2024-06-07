package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Meter -
func Meter(tags ...LuxeGo.Content) *MeterTag {
	return &MeterTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "meter", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type MeterTag struct {
	*ComponentHtmlTag
}

// Value - Specifies the value of an <input> element.
func (m *MeterTag) Value(value string) *MeterTag {
	m.AddAttribute("value", value)
	return m
}

// Min -
func (m *MeterTag) Min(value string) *MeterTag {
	m.AddAttribute("min", value)
	return m
}

// Max -
func (m *MeterTag) Max(value string) *MeterTag {
	m.AddAttribute("max", value)
	return m
}

// Low -
func (m *MeterTag) Low(value string) *MeterTag {
	m.AddAttribute("low", value)
	return m
}

// High -
func (m *MeterTag) High(value string) *MeterTag {
	m.AddAttribute("high", value)
	return m
}

// Optimum -
func (m *MeterTag) Optimum(value string) *MeterTag {
	m.AddAttribute("optimum", value)
	return m
}

// Form -
func (m *MeterTag) Form(value string) *MeterTag {
	m.AddAttribute("form", value)
	return m
}
