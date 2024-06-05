package html

import (
	"LuxeGo/internal/lx"
)

//Audio - 
func Audio(tags ...lx.Content) *AudioTag {
	return &AudioTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "audio", Attributes: &lx.Attributes{}, Children: &tags}}
}

type AudioTag struct {
	*ComponentHtmlTag
}

//Autoplay - Specifies that the audio or video should start playing as soon as it is ready.
func (a *AudioTag) Autoplay() *AudioTag {
	a.AddAttribute("autoplay", "")
	return a
}

//Controls - Specifies that audio/video controls should be displayed.
func (a *AudioTag) Controls() *AudioTag {
	a.AddAttribute("controls", "")
	return a
}

//Controlslist - 
func (a *AudioTag) Controlslist(value string) *AudioTag {
	a.AddAttribute("controlslist", value)
	return a
}

//Crossorigin - How to handle cross-origin requests for the element.
func (a *AudioTag) Crossorigin(value string) *AudioTag {
	a.AddAttribute("crossorigin", value)
	return a
}

//Disableremoteplayback - 
func (a *AudioTag) Disableremoteplayback(value string) *AudioTag {
	a.AddAttribute("disableremoteplayback", value)
	return a
}

//Loop - 
func (a *AudioTag) Loop() *AudioTag {
	a.AddAttribute("loop", "")
	return a
}

//Muted - 
func (a *AudioTag) Muted() *AudioTag {
	a.AddAttribute("muted", "")
	return a
}

//Preload - 
func (a *AudioTag) Preload(value string) *AudioTag {
	a.AddAttribute("preload", value)
	return a
}

//Src - Specifies the URL of an image.
func (a *AudioTag) Src(value string) *AudioTag {
	a.AddAttribute("src", value)
	return a
}
