package svg

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Image -
func Image() *ImageTag {
	return &ImageTag{ComponentSvgTag: &ComponentSvgTag{Name: "image", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type ImageTag struct {
	*ComponentSvgTag
}

// Href -
func (i *ImageTag) Href(value string) *ImageTag {
	i.AddAttribute("href", value)
	return i
}

// X -
func (i *ImageTag) X(value string) *ImageTag {
	i.AddAttribute("x", value)
	return i
}

// Y -
func (i *ImageTag) Y(value string) *ImageTag {
	i.AddAttribute("y", value)
	return i
}

// Width -
func (i *ImageTag) Width(value string) *ImageTag {
	i.AddAttribute("width", value)
	return i
}

// Height -
func (i *ImageTag) Height(value string) *ImageTag {
	i.AddAttribute("height", value)
	return i
}

// PreserveAspectRatio -
func (i *ImageTag) PreserveAspectRatio(value string) *ImageTag {
	i.AddAttribute("preserveaspectratio", value)
	return i
}
