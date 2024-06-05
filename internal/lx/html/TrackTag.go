package html

import (
	"LuxeGo/internal/lx"
)

//Track - 
func Track() *TrackTag {
	return &TrackTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "track", Attributes: &lx.Attributes{}, Children: nil}}
}

type TrackTag struct {
	*ComponentHtmlTag
}

//Default - 
func (t *TrackTag) Default() *TrackTag {
	t.AddAttribute("default", "")
	return t
}

//Kind - 
func (t *TrackTag) Kind(value string) *TrackTag {
	t.AddAttribute("kind", value)
	return t
}

//Label - 
func (t *TrackTag) Label(value string) *TrackTag {
	t.AddAttribute("label", value)
	return t
}

//Src - Specifies the URL of an image.
func (t *TrackTag) Src(value string) *TrackTag {
	t.AddAttribute("src", value)
	return t
}

//Srclang - 
func (t *TrackTag) Srclang(value string) *TrackTag {
	t.AddAttribute("srclang", value)
	return t
}
