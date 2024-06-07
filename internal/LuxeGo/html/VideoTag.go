package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Video -
func Video(tags ...LuxeGo.Content) *VideoTag {
	return &VideoTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "video", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type VideoTag struct {
	*ComponentHtmlTag
}

// Autoplay - Specifies that the audio or video should start playing as soon as it is ready.
func (v *VideoTag) Autoplay() *VideoTag {
	v.AddAttribute("autoplay", "")
	return v
}

// Controls - Specifies that audio/video controls should be displayed.
func (v *VideoTag) Controls() *VideoTag {
	v.AddAttribute("controls", "")
	return v
}

// Controlslist -
func (v *VideoTag) Controlslist(value string) *VideoTag {
	v.AddAttribute("controlslist", value)
	return v
}

// Crossorigin - How to handle cross-origin requests for the element.
func (v *VideoTag) Crossorigin(value string) *VideoTag {
	v.AddAttribute("crossorigin", value)
	return v
}

// Disablepictureinpicture -
func (v *VideoTag) Disablepictureinpicture(value string) *VideoTag {
	v.AddAttribute("disablepictureinpicture", value)
	return v
}

// Disableremoteplayback -
func (v *VideoTag) Disableremoteplayback(value string) *VideoTag {
	v.AddAttribute("disableremoteplayback", value)
	return v
}

// Height -
func (v *VideoTag) Height(value string) *VideoTag {
	v.AddAttribute("height", value)
	return v
}

// Loop -
func (v *VideoTag) Loop() *VideoTag {
	v.AddAttribute("loop", "")
	return v
}

// Muted -
func (v *VideoTag) Muted() *VideoTag {
	v.AddAttribute("muted", "")
	return v
}

// Playsinline -
func (v *VideoTag) Playsinline(value string) *VideoTag {
	v.AddAttribute("playsinline", value)
	return v
}

// Poster -
func (v *VideoTag) Poster(value string) *VideoTag {
	v.AddAttribute("poster", value)
	return v
}

// Preload -
func (v *VideoTag) Preload(value string) *VideoTag {
	v.AddAttribute("preload", value)
	return v
}

// Src - Specifies the URL of an image.
func (v *VideoTag) Src(value string) *VideoTag {
	v.AddAttribute("src", value)
	return v
}

// Width -
func (v *VideoTag) Width(value string) *VideoTag {
	v.AddAttribute("width", value)
	return v
}
