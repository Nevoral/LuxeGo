package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Img -
func Img() *ImgTag {
	return &ImgTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "img", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type ImgTag struct {
	*ComponentHtmlTag
}

// Alt -  Provides alternative text for an image.
func (i *ImgTag) Alt(value string) *ImgTag {
	i.AddAttribute("alt", value)
	return i
}

// Crossorigin - How to handle cross-origin requests for the element.
func (i *ImgTag) Crossorigin(value string) *ImgTag {
	i.AddAttribute("crossorigin", value)
	return i
}

// Decoding -
func (i *ImgTag) Decoding(value string) *ImgTag {
	i.AddAttribute("decoding", value)
	return i
}

// Elementtiming -
func (i *ImgTag) Elementtiming(value string) *ImgTag {
	i.AddAttribute("elementtiming", value)
	return i
}

// Fetchpriority -
func (i *ImgTag) Fetchpriority(value string) *ImgTag {
	i.AddAttribute("fetchpriority", value)
	return i
}

// Height -
func (i *ImgTag) Height(value string) *ImgTag {
	i.AddAttribute("height", value)
	return i
}

// Ismap -
func (i *ImgTag) Ismap() *ImgTag {
	i.AddAttribute("ismap", "")
	return i
}

// Loading -
func (i *ImgTag) Loading(value string) *ImgTag {
	i.AddAttribute("loading", value)
	return i
}

// Referrerpolicy -
func (i *ImgTag) Referrerpolicy(value string) *ImgTag {
	i.AddAttribute("referrerpolicy", value)
	return i
}

// Sizes -
func (i *ImgTag) Sizes(value string) *ImgTag {
	i.AddAttribute("sizes", value)
	return i
}

// Src - Specifies the URL of an image.
func (i *ImgTag) Src(value string) *ImgTag {
	i.AddAttribute("src", value)
	return i
}

// Srcset -
func (i *ImgTag) Srcset(value string) *ImgTag {
	i.AddAttribute("srcset", value)
	return i
}

// Width -
func (i *ImgTag) Width(value string) *ImgTag {
	i.AddAttribute("width", value)
	return i
}

// Usemap -
func (i *ImgTag) Usemap(value string) *ImgTag {
	i.AddAttribute("usemap", value)
	return i
}
