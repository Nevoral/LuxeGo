package html

import (
	"LuxeGo/internal/LuxeGo"
	"fmt"
	"slices"
)

type ComponentHtmlTag struct {
	Name       string
	Attributes *LuxeGo.Attributes
	Msg        string
	Children   *[]LuxeGo.Content
}

// AccessKey -
func (c *ComponentHtmlTag) AccessKey(value string) *ComponentHtmlTag {
	c.AddAttribute("accesskey", value)
	return c
}

// Contenteditable -
func (c *ComponentHtmlTag) Contenteditable(value string) *ComponentHtmlTag {
	c.AddAttribute("contenteditable", value)
	return c
}

// Data -
func (c *ComponentHtmlTag) Data(name, value string) *ComponentHtmlTag {
	a := fmt.Sprintf("data-%s", name)
	c.AddAttribute(a, value)
	return c
}

// Hidden -
func (c *ComponentHtmlTag) Hidden(value string) *ComponentHtmlTag {
	c.AddAttribute("hidden", value)
	return c
}

// Id -
func (c *ComponentHtmlTag) Id(value string) *ComponentHtmlTag {
	c.AddAttribute("id", value)
	return c
}

// Nonce -
func (c *ComponentHtmlTag) Nonce(value string) *ComponentHtmlTag {
	c.AddAttribute("nonce", value)
	return c
}

// Autofocus -
func (c *ComponentHtmlTag) Autofocus(value string) *ComponentHtmlTag {
	c.AddAttribute("autofocus", value)
	return c
}

// Style -
func (c *ComponentHtmlTag) Style(value string) *ComponentHtmlTag {
	c.AddAttribute("style", value)
	return c
}

// Tabindex -
func (c *ComponentHtmlTag) Tabindex(value string) *ComponentHtmlTag {
	c.AddAttribute("tabindex", value)
	return c
}

// Class -
func (c *ComponentHtmlTag) Class(value string) *ComponentHtmlTag {
	c.AddAttribute("class", value)
	return c
}

// Role -
func (c *ComponentHtmlTag) Role(value string) *ComponentHtmlTag {
	c.AddAttribute("role", value)
	return c
}

// Aria -
func (c *ComponentHtmlTag) Aria(name, value string) *ComponentHtmlTag {
	a := fmt.Sprintf("aria-%s", name)
	if slices.Contains(LuxeGo.GlobAriaName, name) {
		c.AddAttribute(a, value)
	} else {
		//TODO: Specific aria attributes need to be check before adding to the map of attributes
		c.AddAttribute(a, value)
	}
	return c
}

// ItemRef -
func (c *ComponentHtmlTag) ItemRef(value string) *ComponentHtmlTag {
	c.AddAttribute("itemref", value)
	return c
}

// Part -
func (c *ComponentHtmlTag) Part(value string) *ComponentHtmlTag {
	c.AddAttribute("part", value)
	return c
}

// Autocapitalize -
func (c *ComponentHtmlTag) Autocapitalize(value string) *ComponentHtmlTag {
	c.AddAttribute("autocapitalize", value)
	return c
}

// Dir -
func (c *ComponentHtmlTag) Dir(value string) *ComponentHtmlTag {
	c.AddAttribute("dir", value)
	return c
}

// Draggable -
func (c *ComponentHtmlTag) Draggable(value string) *ComponentHtmlTag {
	c.AddAttribute("draggable", value)
	return c
}

// Lang -
func (c *ComponentHtmlTag) Lang(value string) *ComponentHtmlTag {
	c.AddAttribute("lang", value)
	return c
}

// Spellcheck -
func (c *ComponentHtmlTag) Spellcheck(value string) *ComponentHtmlTag {
	c.AddAttribute("spellcheck", value)
	return c
}

// VirtualKeyBoardPolicy -
func (c *ComponentHtmlTag) VirtualKeyBoardPolicy(value string) *ComponentHtmlTag {
	c.AddAttribute("virtualkeyboardpolicy", value)
	return c
}

// EnterKeyHint -
func (c *ComponentHtmlTag) EnterKeyHint(value string) *ComponentHtmlTag {
	c.AddAttribute("enterkeyhint", value)
	return c
}

// Inert -
func (c *ComponentHtmlTag) Inert(value string) *ComponentHtmlTag {
	c.AddAttribute("inert", value)
	return c
}

// InputMode -
func (c *ComponentHtmlTag) InputMode(value string) *ComponentHtmlTag {
	c.AddAttribute("inputmode", value)
	return c
}

// ItemId -
func (c *ComponentHtmlTag) ItemId(value string) *ComponentHtmlTag {
	c.AddAttribute("itemid", value)
	return c
}

// ItemProp -
func (c *ComponentHtmlTag) ItemProp(value string) *ComponentHtmlTag {
	c.AddAttribute("itemprop", value)
	return c
}

// Popover -
func (c *ComponentHtmlTag) Popover() *ComponentHtmlTag {
	c.AddAttribute("popover", "")
	return c
}

// ExportParts -
func (c *ComponentHtmlTag) ExportParts(value string) *ComponentHtmlTag {
	c.AddAttribute("exportparts", value)
	return c
}

// Is -
func (c *ComponentHtmlTag) Is(value string) *ComponentHtmlTag {
	c.AddAttribute("is", value)
	return c
}

// Slot -
func (c *ComponentHtmlTag) Slot(value string) *ComponentHtmlTag {
	c.AddAttribute("slot", value)
	return c
}

// Translate -
func (c *ComponentHtmlTag) Translate(value string) *ComponentHtmlTag {
	c.AddAttribute("translate", value)
	return c
}

// ItemScope -
func (c *ComponentHtmlTag) ItemScope(value string) *ComponentHtmlTag {
	c.AddAttribute("itemscope", value)
	return c
}

// ItemType -
func (c *ComponentHtmlTag) ItemType(value string) *ComponentHtmlTag {
	c.AddAttribute("itemtype", value)
	return c
}

// Title -
func (c *ComponentHtmlTag) Title(value string) *ComponentHtmlTag {
	c.AddAttribute("title", value)
	return c
}

// Ondragstart -
func (c *ComponentHtmlTag) Ondragstart(value string) *ComponentHtmlTag {
	c.AddAttribute("ondragstart", value)
	return c
}

// Onmousewheel -
func (c *ComponentHtmlTag) Onmousewheel(value string) *ComponentHtmlTag {
	c.AddAttribute("onmousewheel", value)
	return c
}

// Onwaiting -
func (c *ComponentHtmlTag) Onwaiting(value string) *ComponentHtmlTag {
	c.AddAttribute("onwaiting", value)
	return c
}

// Onabort -
func (c *ComponentHtmlTag) Onabort(value string) *ComponentHtmlTag {
	c.AddAttribute("onabort", value)
	return c
}

// Oncanplay -
func (c *ComponentHtmlTag) Oncanplay(value string) *ComponentHtmlTag {
	c.AddAttribute("oncanplay", value)
	return c
}

// Ondragend -
func (c *ComponentHtmlTag) Ondragend(value string) *ComponentHtmlTag {
	c.AddAttribute("ondragend", value)
	return c
}

// Onload -
func (c *ComponentHtmlTag) Onload(value string) *ComponentHtmlTag {
	c.AddAttribute("onload", value)
	return c
}

// Onselect -
func (c *ComponentHtmlTag) Onselect(value string) *ComponentHtmlTag {
	c.AddAttribute("onselect", value)
	return c
}

// Onblur -
func (c *ComponentHtmlTag) Onblur(value string) *ComponentHtmlTag {
	c.AddAttribute("onblur", value)
	return c
}

// Oncontextmenu -
func (c *ComponentHtmlTag) Oncontextmenu(value string) *ComponentHtmlTag {
	c.AddAttribute("oncontextmenu", value)
	return c
}

// Onkeyup -
func (c *ComponentHtmlTag) Onkeyup(value string) *ComponentHtmlTag {
	c.AddAttribute("onkeyup", value)
	return c
}

// Onmouseover -
func (c *ComponentHtmlTag) Onmouseover(value string) *ComponentHtmlTag {
	c.AddAttribute("onmouseover", value)
	return c
}

// Onsuspend -
func (c *ComponentHtmlTag) Onsuspend(value string) *ComponentHtmlTag {
	c.AddAttribute("onsuspend", value)
	return c
}

// Oncanplaythrough -
func (c *ComponentHtmlTag) Oncanplaythrough(value string) *ComponentHtmlTag {
	c.AddAttribute("oncanplaythrough", value)
	return c
}

// Ondragleave -
func (c *ComponentHtmlTag) Ondragleave(value string) *ComponentHtmlTag {
	c.AddAttribute("ondragleave", value)
	return c
}

// Onfocus -
func (c *ComponentHtmlTag) Onfocus(value string) *ComponentHtmlTag {
	c.AddAttribute("onfocus", value)
	return c
}

// Onautocompleteerror -
func (c *ComponentHtmlTag) Onautocompleteerror(value string) *ComponentHtmlTag {
	c.AddAttribute("onautocompleteerror", value)
	return c
}

// Ondragenter -
func (c *ComponentHtmlTag) Ondragenter(value string) *ComponentHtmlTag {
	c.AddAttribute("ondragenter", value)
	return c
}

// Onloadedmetadata -
func (c *ComponentHtmlTag) Onloadedmetadata(value string) *ComponentHtmlTag {
	c.AddAttribute("onloadedmetadata", value)
	return c
}

// Onemptied -
func (c *ComponentHtmlTag) Onemptied(value string) *ComponentHtmlTag {
	c.AddAttribute("onemptied", value)
	return c
}

// Onplaying -
func (c *ComponentHtmlTag) Onplaying(value string) *ComponentHtmlTag {
	c.AddAttribute("onplaying", value)
	return c
}

// Onscroll -
func (c *ComponentHtmlTag) Onscroll(value string) *ComponentHtmlTag {
	c.AddAttribute("onscroll", value)
	return c
}

// Onseeked -
func (c *ComponentHtmlTag) Onseeked(value string) *ComponentHtmlTag {
	c.AddAttribute("onseeked", value)
	return c
}

// Onseeking -
func (c *ComponentHtmlTag) Onseeking(value string) *ComponentHtmlTag {
	c.AddAttribute("onseeking", value)
	return c
}

// Ondragover -
func (c *ComponentHtmlTag) Ondragover(value string) *ComponentHtmlTag {
	c.AddAttribute("ondragover", value)
	return c
}

// Onkeydown -
func (c *ComponentHtmlTag) Onkeydown(value string) *ComponentHtmlTag {
	c.AddAttribute("onkeydown", value)
	return c
}

// Onloadeddata -
func (c *ComponentHtmlTag) Onloadeddata(value string) *ComponentHtmlTag {
	c.AddAttribute("onloadeddata", value)
	return c
}

// Onratechange -
func (c *ComponentHtmlTag) Onratechange(value string) *ComponentHtmlTag {
	c.AddAttribute("onratechange", value)
	return c
}

// Oncuechange -
func (c *ComponentHtmlTag) Oncuechange(value string) *ComponentHtmlTag {
	c.AddAttribute("oncuechange", value)
	return c
}

// Onkeypress -
func (c *ComponentHtmlTag) Onkeypress(value string) *ComponentHtmlTag {
	c.AddAttribute("onkeypress", value)
	return c
}

// Onplay -
func (c *ComponentHtmlTag) Onplay(value string) *ComponentHtmlTag {
	c.AddAttribute("onplay", value)
	return c
}

// Ontimeupdate -
func (c *ComponentHtmlTag) Ontimeupdate(value string) *ComponentHtmlTag {
	c.AddAttribute("ontimeupdate", value)
	return c
}

// Onautocomplete -
func (c *ComponentHtmlTag) Onautocomplete(value string) *ComponentHtmlTag {
	c.AddAttribute("onautocomplete", value)
	return c
}

// Ondblclick -
func (c *ComponentHtmlTag) Ondblclick(value string) *ComponentHtmlTag {
	c.AddAttribute("ondblclick", value)
	return c
}

// Oninvalid -
func (c *ComponentHtmlTag) Oninvalid(value string) *ComponentHtmlTag {
	c.AddAttribute("oninvalid", value)
	return c
}

// Onsort -
func (c *ComponentHtmlTag) Onsort(value string) *ComponentHtmlTag {
	c.AddAttribute("onsort", value)
	return c
}

// Ondrop -
func (c *ComponentHtmlTag) Ondrop(value string) *ComponentHtmlTag {
	c.AddAttribute("ondrop", value)
	return c
}

// Ondurationchange -
func (c *ComponentHtmlTag) Ondurationchange(value string) *ComponentHtmlTag {
	c.AddAttribute("ondurationchange", value)
	return c
}

// Onmousemove -
func (c *ComponentHtmlTag) Onmousemove(value string) *ComponentHtmlTag {
	c.AddAttribute("onmousemove", value)
	return c
}

// Onmouseup -
func (c *ComponentHtmlTag) Onmouseup(value string) *ComponentHtmlTag {
	c.AddAttribute("onmouseup", value)
	return c
}

// Onsubmit -
func (c *ComponentHtmlTag) Onsubmit(value string) *ComponentHtmlTag {
	c.AddAttribute("onsubmit", value)
	return c
}

// Ondrag -
func (c *ComponentHtmlTag) Ondrag(value string) *ComponentHtmlTag {
	c.AddAttribute("ondrag", value)
	return c
}

// Onended -
func (c *ComponentHtmlTag) Onended(value string) *ComponentHtmlTag {
	c.AddAttribute("onended", value)
	return c
}

// Onmouseout -
func (c *ComponentHtmlTag) Onmouseout(value string) *ComponentHtmlTag {
	c.AddAttribute("onmouseout", value)
	return c
}

// Onresize -
func (c *ComponentHtmlTag) Onresize(value string) *ComponentHtmlTag {
	c.AddAttribute("onresize", value)
	return c
}

// Onstalled -
func (c *ComponentHtmlTag) Onstalled(value string) *ComponentHtmlTag {
	c.AddAttribute("onstalled", value)
	return c
}

// Onerror -
func (c *ComponentHtmlTag) Onerror(value string) *ComponentHtmlTag {
	c.AddAttribute("onerror", value)
	return c
}

// Onloadstart -
func (c *ComponentHtmlTag) Onloadstart(value string) *ComponentHtmlTag {
	c.AddAttribute("onloadstart", value)
	return c
}

// Onmousedown -
func (c *ComponentHtmlTag) Onmousedown(value string) *ComponentHtmlTag {
	c.AddAttribute("onmousedown", value)
	return c
}

// Onmouseleave -
func (c *ComponentHtmlTag) Onmouseleave(value string) *ComponentHtmlTag {
	c.AddAttribute("onmouseleave", value)
	return c
}

// Oninput -
func (c *ComponentHtmlTag) Oninput(value string) *ComponentHtmlTag {
	c.AddAttribute("oninput", value)
	return c
}

// Onprogress -
func (c *ComponentHtmlTag) Onprogress(value string) *ComponentHtmlTag {
	c.AddAttribute("onprogress", value)
	return c
}

// Onreset -
func (c *ComponentHtmlTag) Onreset(value string) *ComponentHtmlTag {
	c.AddAttribute("onreset", value)
	return c
}

// Ontoggle -
func (c *ComponentHtmlTag) Ontoggle(value string) *ComponentHtmlTag {
	c.AddAttribute("ontoggle", value)
	return c
}

// Onchange -
func (c *ComponentHtmlTag) Onchange(value string) *ComponentHtmlTag {
	c.AddAttribute("onchange", value)
	return c
}

// Onvolumechange -
func (c *ComponentHtmlTag) Onvolumechange(value string) *ComponentHtmlTag {
	c.AddAttribute("onvolumechange", value)
	return c
}

// Onshow -
func (c *ComponentHtmlTag) Onshow(value string) *ComponentHtmlTag {
	c.AddAttribute("onshow", value)
	return c
}

// Oncancel -
func (c *ComponentHtmlTag) Oncancel(value string) *ComponentHtmlTag {
	c.AddAttribute("oncancel", value)
	return c
}

// Onclick -
func (c *ComponentHtmlTag) Onclick(value string) *ComponentHtmlTag {
	c.AddAttribute("onclick", value)
	return c
}

// Onclose -
func (c *ComponentHtmlTag) Onclose(value string) *ComponentHtmlTag {
	c.AddAttribute("onclose", value)
	return c
}

// Onmouseenter -
func (c *ComponentHtmlTag) Onmouseenter(value string) *ComponentHtmlTag {
	c.AddAttribute("onmouseenter", value)
	return c
}

// Onpause -
func (c *ComponentHtmlTag) Onpause(value string) *ComponentHtmlTag {
	c.AddAttribute("onpause", value)
	return c
}
