package svg

import (
	"LuxeGo/internal/lx"
)

// FeTurbulence -
func FeTurbulence() *FeTurbulenceTag {
	return &FeTurbulenceTag{ComponentSvgTag: &ComponentSvgTag{Name: "feturbulence", Attributes: &lx.Attributes{}, Children: nil}}
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
