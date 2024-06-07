package svg

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Line -
func Line() *LineTag {
	return &LineTag{ComponentSvgTag: &ComponentSvgTag{Name: "line", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type LineTag struct {
	*ComponentSvgTag
}

// X1 -
func (l *LineTag) X1(value string) *LineTag {
	l.AddAttribute("x1", value)
	return l
}

// Y1 -
func (l *LineTag) Y1(value string) *LineTag {
	l.AddAttribute("y1", value)
	return l
}

// X2 -
func (l *LineTag) X2(value string) *LineTag {
	l.AddAttribute("x2", value)
	return l
}

// Y2 -
func (l *LineTag) Y2(value string) *LineTag {
	l.AddAttribute("y2", value)
	return l
}

// PathLength -
func (l *LineTag) PathLength(value string) *LineTag {
	l.AddAttribute("pathlength", value)
	return l
}
