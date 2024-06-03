package svg

import (
	"LuxeGo/internal/lx"
)

type Component struct {
	*lx.WebComponent
}

func (s *Component) Id(value string) *Component {
	s.AddAttribute("id", value)
	return s
}

func A(tags ...lx.Content) *ATag {
	return &ATag{&Component{&lx.WebComponent{Name: "a", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type ATag struct {
	*Component
}

func (a *ATag) Download(value string) *ATag {
	a.AddAttribute("download", value)
	return a
}

func Animate(tags ...lx.Content) *AnimateSvgTag {
	return &AnimateSvgTag{&Component{&lx.WebComponent{Name: "animate", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type AnimateSvgTag struct {
	*Component
}
