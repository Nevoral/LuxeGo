package svg

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// FeTurbulence -
func FeTurbulence() *FeTurbulenceTag {
	return &FeTurbulenceTag{ComponentSvgTag: &ComponentSvgTag{Name: "feturbulence", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type FeTurbulenceTag struct {
	*ComponentSvgTag
}

// BaseFrequency -
func (f *FeTurbulenceTag) BaseFrequency(value string) *FeTurbulenceTag {
	f.AddAttribute("basefrequency", value)
	return f
}

// NumOctaves -
func (f *FeTurbulenceTag) NumOctaves(value string) *FeTurbulenceTag {
	f.AddAttribute("numoctaves", value)
	return f
}

// Seed -
func (f *FeTurbulenceTag) Seed(value string) *FeTurbulenceTag {
	f.AddAttribute("seed", value)
	return f
}

// StitchTiles -
func (f *FeTurbulenceTag) StitchTiles(value string) *FeTurbulenceTag {
	f.AddAttribute("stitchtiles", value)
	return f
}

// Type -
func (f *FeTurbulenceTag) Type(value string) *FeTurbulenceTag {
	f.AddAttribute("type", value)
	return f
}
