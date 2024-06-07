package svg

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
	"slices"
)

type ComponentSvgTag struct {
	Name       string
	Attributes *LuxeGo.Attributes
	Msg        string
	Children   *[]LuxeGo.Content
}

// Class -
func (c *ComponentSvgTag) Class(value string) *ComponentSvgTag {
	c.AddAttribute("class", value)
	return c
}

// Style -
func (c *ComponentSvgTag) Style(value string) *ComponentSvgTag {
	c.AddAttribute("style", value)
	return c
}

// Tabindex -
func (c *ComponentSvgTag) Tabindex(value string) *ComponentSvgTag {
	c.AddAttribute("tabindex", value)
	return c
}

// Role -
func (c *ComponentSvgTag) Role(value string) *ComponentSvgTag {
	c.AddAttribute("role", value)
	return c
}

// Data -
func (c *ComponentSvgTag) Data(name, value string) *ComponentSvgTag {
	var a string
	if name == "" {
		a = fmt.Sprintf("data-%s", name)
	} else {
		a = fmt.Sprintf("data-%s", name)
	}
	c.AddAttribute(a, value)
	return c
}

// Id -
func (c *ComponentSvgTag) Id(value string) *ComponentSvgTag {
	c.AddAttribute("id", value)
	return c
}

// XmlBase -
func (c *ComponentSvgTag) XmlBase(value string) *ComponentSvgTag {
	c.AddAttribute("xml:base", value)
	return c
}

// XmlLang -
func (c *ComponentSvgTag) XmlLang(value string) *ComponentSvgTag {
	c.AddAttribute("xml:lang", value)
	return c
}

// XmlSpace -
func (c *ComponentSvgTag) XmlSpace(value string) *ComponentSvgTag {
	c.AddAttribute("xml:space", value)
	return c
}

// Aria -
func (c *ComponentSvgTag) Aria(name, value string) *ComponentSvgTag {
	var a string
	if name == "" {
		a = fmt.Sprintf("aria-%s", name)
	} else {
		a = fmt.Sprintf("aria-%s", name)
	}
	if slices.Contains(LuxeGo.GlobAriaName, name) {
		c.AddAttribute(a, value)
	} else {
		//TODO: Specific aria attributes need to be check before adding to the map of attributes
		c.AddAttribute(a, value)
	}
	return c
}

// Onload -
func (c *ComponentSvgTag) Onload(value string) *ComponentSvgTag {
	c.AddAttribute("onload", value)
	return c
}

// Ondragenter -
func (c *ComponentSvgTag) Ondragenter(value string) *ComponentSvgTag {
	c.AddAttribute("ondragenter", value)
	return c
}

// Onmouseup -
func (c *ComponentSvgTag) Onmouseup(value string) *ComponentSvgTag {
	c.AddAttribute("onmouseup", value)
	return c
}

// Onscroll -
func (c *ComponentSvgTag) Onscroll(value string) *ComponentSvgTag {
	c.AddAttribute("onscroll", value)
	return c
}

// Ondrop -
func (c *ComponentSvgTag) Ondrop(value string) *ComponentSvgTag {
	c.AddAttribute("ondrop", value)
	return c
}

// Onmouseenter -
func (c *ComponentSvgTag) Onmouseenter(value string) *ComponentSvgTag {
	c.AddAttribute("onmouseenter", value)
	return c
}

// Onkeydown -
func (c *ComponentSvgTag) Onkeydown(value string) *ComponentSvgTag {
	c.AddAttribute("onkeydown", value)
	return c
}

// Onmousedown -
func (c *ComponentSvgTag) Onmousedown(value string) *ComponentSvgTag {
	c.AddAttribute("onmousedown", value)
	return c
}

// Onratechange -
func (c *ComponentSvgTag) Onratechange(value string) *ComponentSvgTag {
	c.AddAttribute("onratechange", value)
	return c
}

// Onchange -
func (c *ComponentSvgTag) Onchange(value string) *ComponentSvgTag {
	c.AddAttribute("onchange", value)
	return c
}

// Onmouseleave -
func (c *ComponentSvgTag) Onmouseleave(value string) *ComponentSvgTag {
	c.AddAttribute("onmouseleave", value)
	return c
}

// Onseeking -
func (c *ComponentSvgTag) Onseeking(value string) *ComponentSvgTag {
	c.AddAttribute("onseeking", value)
	return c
}

// Ondurationchange -
func (c *ComponentSvgTag) Ondurationchange(value string) *ComponentSvgTag {
	c.AddAttribute("ondurationchange", value)
	return c
}

// Ondrag -
func (c *ComponentSvgTag) Ondrag(value string) *ComponentSvgTag {
	c.AddAttribute("ondrag", value)
	return c
}

// Onpause -
func (c *ComponentSvgTag) Onpause(value string) *ComponentSvgTag {
	c.AddAttribute("onpause", value)
	return c
}

// Onwaiting -
func (c *ComponentSvgTag) Onwaiting(value string) *ComponentSvgTag {
	c.AddAttribute("onwaiting", value)
	return c
}

// Onautocompleteerror -
func (c *ComponentSvgTag) Onautocompleteerror(value string) *ComponentSvgTag {
	c.AddAttribute("onautocompleteerror", value)
	return c
}

// Ondragleave -
func (c *ComponentSvgTag) Ondragleave(value string) *ComponentSvgTag {
	c.AddAttribute("ondragleave", value)
	return c
}

// Onerror -
func (c *ComponentSvgTag) Onerror(value string) *ComponentSvgTag {
	c.AddAttribute("onerror", value)
	return c
}

// Onkeyup -
func (c *ComponentSvgTag) Onkeyup(value string) *ComponentSvgTag {
	c.AddAttribute("onkeyup", value)
	return c
}

// Onmousemove -
func (c *ComponentSvgTag) Onmousemove(value string) *ComponentSvgTag {
	c.AddAttribute("onmousemove", value)
	return c
}

// Onplay -
func (c *ComponentSvgTag) Onplay(value string) *ComponentSvgTag {
	c.AddAttribute("onplay", value)
	return c
}

// Ontimeupdate -
func (c *ComponentSvgTag) Ontimeupdate(value string) *ComponentSvgTag {
	c.AddAttribute("ontimeupdate", value)
	return c
}

// Ontoggle -
func (c *ComponentSvgTag) Ontoggle(value string) *ComponentSvgTag {
	c.AddAttribute("ontoggle", value)
	return c
}

// Onblur -
func (c *ComponentSvgTag) Onblur(value string) *ComponentSvgTag {
	c.AddAttribute("onblur", value)
	return c
}

// Oncontextmenu -
func (c *ComponentSvgTag) Oncontextmenu(value string) *ComponentSvgTag {
	c.AddAttribute("oncontextmenu", value)
	return c
}

// Onloadstart -
func (c *ComponentSvgTag) Onloadstart(value string) *ComponentSvgTag {
	c.AddAttribute("onloadstart", value)
	return c
}

// Onmouseout -
func (c *ComponentSvgTag) Onmouseout(value string) *ComponentSvgTag {
	c.AddAttribute("onmouseout", value)
	return c
}

// Onprogress -
func (c *ComponentSvgTag) Onprogress(value string) *ComponentSvgTag {
	c.AddAttribute("onprogress", value)
	return c
}

// Onsubmit -
func (c *ComponentSvgTag) Onsubmit(value string) *ComponentSvgTag {
	c.AddAttribute("onsubmit", value)
	return c
}

// Oncanplaythrough -
func (c *ComponentSvgTag) Oncanplaythrough(value string) *ComponentSvgTag {
	c.AddAttribute("oncanplaythrough", value)
	return c
}

// Onclose -
func (c *ComponentSvgTag) Onclose(value string) *ComponentSvgTag {
	c.AddAttribute("onclose", value)
	return c
}

// Onfocus -
func (c *ComponentSvgTag) Onfocus(value string) *ComponentSvgTag {
	c.AddAttribute("onfocus", value)
	return c
}

// Oncancel -
func (c *ComponentSvgTag) Oncancel(value string) *ComponentSvgTag {
	c.AddAttribute("oncancel", value)
	return c
}

// Onmousewheel -
func (c *ComponentSvgTag) Onmousewheel(value string) *ComponentSvgTag {
	c.AddAttribute("onmousewheel", value)
	return c
}

// Onclick -
func (c *ComponentSvgTag) Onclick(value string) *ComponentSvgTag {
	c.AddAttribute("onclick", value)
	return c
}

// Onloadedmetadata -
func (c *ComponentSvgTag) Onloadedmetadata(value string) *ComponentSvgTag {
	c.AddAttribute("onloadedmetadata", value)
	return c
}

// Onvolumechange -
func (c *ComponentSvgTag) Onvolumechange(value string) *ComponentSvgTag {
	c.AddAttribute("onvolumechange", value)
	return c
}

// Oninput -
func (c *ComponentSvgTag) Oninput(value string) *ComponentSvgTag {
	c.AddAttribute("oninput", value)
	return c
}

// Ondblclick -
func (c *ComponentSvgTag) Ondblclick(value string) *ComponentSvgTag {
	c.AddAttribute("ondblclick", value)
	return c
}

// Oninvalid -
func (c *ComponentSvgTag) Oninvalid(value string) *ComponentSvgTag {
	c.AddAttribute("oninvalid", value)
	return c
}

// Onselect -
func (c *ComponentSvgTag) Onselect(value string) *ComponentSvgTag {
	c.AddAttribute("onselect", value)
	return c
}

// Onstalled -
func (c *ComponentSvgTag) Onstalled(value string) *ComponentSvgTag {
	c.AddAttribute("onstalled", value)
	return c
}

// Oncanplay -
func (c *ComponentSvgTag) Oncanplay(value string) *ComponentSvgTag {
	c.AddAttribute("oncanplay", value)
	return c
}

// Onemptied -
func (c *ComponentSvgTag) Onemptied(value string) *ComponentSvgTag {
	c.AddAttribute("onemptied", value)
	return c
}

// Onplaying -
func (c *ComponentSvgTag) Onplaying(value string) *ComponentSvgTag {
	c.AddAttribute("onplaying", value)
	return c
}

// Onresize -
func (c *ComponentSvgTag) Onresize(value string) *ComponentSvgTag {
	c.AddAttribute("onresize", value)
	return c
}

// Onsuspend -
func (c *ComponentSvgTag) Onsuspend(value string) *ComponentSvgTag {
	c.AddAttribute("onsuspend", value)
	return c
}

// Onabort -
func (c *ComponentSvgTag) Onabort(value string) *ComponentSvgTag {
	c.AddAttribute("onabort", value)
	return c
}

// Ondragover -
func (c *ComponentSvgTag) Ondragover(value string) *ComponentSvgTag {
	c.AddAttribute("ondragover", value)
	return c
}

// Onkeypress -
func (c *ComponentSvgTag) Onkeypress(value string) *ComponentSvgTag {
	c.AddAttribute("onkeypress", value)
	return c
}

// Onmouseover -
func (c *ComponentSvgTag) Onmouseover(value string) *ComponentSvgTag {
	c.AddAttribute("onmouseover", value)
	return c
}

// Onshow -
func (c *ComponentSvgTag) Onshow(value string) *ComponentSvgTag {
	c.AddAttribute("onshow", value)
	return c
}

// Ondragend -
func (c *ComponentSvgTag) Ondragend(value string) *ComponentSvgTag {
	c.AddAttribute("ondragend", value)
	return c
}

// Onended -
func (c *ComponentSvgTag) Onended(value string) *ComponentSvgTag {
	c.AddAttribute("onended", value)
	return c
}

// Onloadeddata -
func (c *ComponentSvgTag) Onloadeddata(value string) *ComponentSvgTag {
	c.AddAttribute("onloadeddata", value)
	return c
}

// Onreset -
func (c *ComponentSvgTag) Onreset(value string) *ComponentSvgTag {
	c.AddAttribute("onreset", value)
	return c
}

// Onseeked -
func (c *ComponentSvgTag) Onseeked(value string) *ComponentSvgTag {
	c.AddAttribute("onseeked", value)
	return c
}

// Onsort -
func (c *ComponentSvgTag) Onsort(value string) *ComponentSvgTag {
	c.AddAttribute("onsort", value)
	return c
}

// Ondragstart -
func (c *ComponentSvgTag) Ondragstart(value string) *ComponentSvgTag {
	c.AddAttribute("ondragstart", value)
	return c
}

// Oncuechange -
func (c *ComponentSvgTag) Oncuechange(value string) *ComponentSvgTag {
	c.AddAttribute("oncuechange", value)
	return c
}

// Onautocomplete -
func (c *ComponentSvgTag) Onautocomplete(value string) *ComponentSvgTag {
	c.AddAttribute("onautocomplete", value)
	return c
}
