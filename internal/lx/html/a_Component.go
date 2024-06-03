package html

import "LuxeGo/internal/lx"

type ComponentTag struct {
	*lx.WebComponent
}


//Autofocus - 
func (c *ComponentTag) Autofocus(value string) *ComponentTag {
	c.AddAttribute("autofocus", value)
	return c
}

//Data - 
func (c *ComponentTag) Data(value string) *ComponentTag {
	c.AddAttribute("data", value)
	return c
}

//InputMode - 
func (c *ComponentTag) InputMode(value string) *ComponentTag {
	c.AddAttribute("inputmode", value)
	return c
}

//Popover - 
func (c *ComponentTag) Popover() *ComponentTag {
	c.AddAttribute("popover", "")
	return c
}

//VirtualKeyBoardPolicy - 
func (c *ComponentTag) VirtualKeyBoardPolicy(value string) *ComponentTag {
	c.AddAttribute("virtualkeyboardpolicy", value)
	return c
}

//Aria - 
func (c *ComponentTag) Aria(value string) *ComponentTag {
	c.AddAttribute("aria", value)
	return c
}

//Autocapitalize - 
func (c *ComponentTag) Autocapitalize(value string) *ComponentTag {
	c.AddAttribute("autocapitalize", value)
	return c
}

//ItemType - 
func (c *ComponentTag) ItemType(value string) *ComponentTag {
	c.AddAttribute("itemtype", value)
	return c
}

//Nonce - 
func (c *ComponentTag) Nonce(value string) *ComponentTag {
	c.AddAttribute("nonce", value)
	return c
}

//AccessKey - 
func (c *ComponentTag) AccessKey(value string) *ComponentTag {
	c.AddAttribute("accesskey", value)
	return c
}

//Contenteditable - 
func (c *ComponentTag) Contenteditable(value string) *ComponentTag {
	c.AddAttribute("contenteditable", value)
	return c
}

//Is - 
func (c *ComponentTag) Is(value string) *ComponentTag {
	c.AddAttribute("is", value)
	return c
}

//ItemRef - 
func (c *ComponentTag) ItemRef(value string) *ComponentTag {
	c.AddAttribute("itemref", value)
	return c
}

//ItemScope - 
func (c *ComponentTag) ItemScope(value string) *ComponentTag {
	c.AddAttribute("itemscope", value)
	return c
}

//Title - 
func (c *ComponentTag) Title(value string) *ComponentTag {
	c.AddAttribute("title", value)
	return c
}

//Draggable - 
func (c *ComponentTag) Draggable(value string) *ComponentTag {
	c.AddAttribute("draggable", value)
	return c
}

//Inert - 
func (c *ComponentTag) Inert(value string) *ComponentTag {
	c.AddAttribute("inert", value)
	return c
}

//Role - 
func (c *ComponentTag) Role(value string) *ComponentTag {
	c.AddAttribute("role", value)
	return c
}

//Style - 
func (c *ComponentTag) Style(value string) *ComponentTag {
	c.AddAttribute("style", value)
	return c
}

//Id - 
func (c *ComponentTag) Id(value string) *ComponentTag {
	c.AddAttribute("id", value)
	return c
}

//Lang - 
func (c *ComponentTag) Lang(value string) *ComponentTag {
	c.AddAttribute("lang", value)
	return c
}

//ItemProp - 
func (c *ComponentTag) ItemProp(value string) *ComponentTag {
	c.AddAttribute("itemprop", value)
	return c
}

//Spellcheck - 
func (c *ComponentTag) Spellcheck(value string) *ComponentTag {
	c.AddAttribute("spellcheck", value)
	return c
}

//Dir - 
func (c *ComponentTag) Dir(value string) *ComponentTag {
	c.AddAttribute("dir", value)
	return c
}

//Hidden - 
func (c *ComponentTag) Hidden(value string) *ComponentTag {
	c.AddAttribute("hidden", value)
	return c
}

//Part - 
func (c *ComponentTag) Part(value string) *ComponentTag {
	c.AddAttribute("part", value)
	return c
}

//EnterKeyHint - 
func (c *ComponentTag) EnterKeyHint(value string) *ComponentTag {
	c.AddAttribute("enterkeyhint", value)
	return c
}

//ExportParts - 
func (c *ComponentTag) ExportParts(value string) *ComponentTag {
	c.AddAttribute("exportparts", value)
	return c
}

//Tabindex - 
func (c *ComponentTag) Tabindex(value string) *ComponentTag {
	c.AddAttribute("tabindex", value)
	return c
}

//Slot - 
func (c *ComponentTag) Slot(value string) *ComponentTag {
	c.AddAttribute("slot", value)
	return c
}

//Translate - 
func (c *ComponentTag) Translate(value string) *ComponentTag {
	c.AddAttribute("translate", value)
	return c
}

//Class - 
func (c *ComponentTag) Class(value string) *ComponentTag {
	c.AddAttribute("class", value)
	return c
}

//ItemId - 
func (c *ComponentTag) ItemId(value string) *ComponentTag {
	c.AddAttribute("itemid", value)
	return c
}

//Onmouseenter - 
func (c *ComponentTag) Onmouseenter(value string) *ComponentTag {
	c.AddAttribute("onmouseenter", value)
	return c
}

//Onreset - 
func (c *ComponentTag) Onreset(value string) *ComponentTag {
	c.AddAttribute("onreset", value)
	return c
}

//Onseeking - 
func (c *ComponentTag) Onseeking(value string) *ComponentTag {
	c.AddAttribute("onseeking", value)
	return c
}

//Onselect - 
func (c *ComponentTag) Onselect(value string) *ComponentTag {
	c.AddAttribute("onselect", value)
	return c
}

//Onkeydown - 
func (c *ComponentTag) Onkeydown(value string) *ComponentTag {
	c.AddAttribute("onkeydown", value)
	return c
}

//Onloadeddata - 
func (c *ComponentTag) Onloadeddata(value string) *ComponentTag {
	c.AddAttribute("onloadeddata", value)
	return c
}

//Onloadedmetadata - 
func (c *ComponentTag) Onloadedmetadata(value string) *ComponentTag {
	c.AddAttribute("onloadedmetadata", value)
	return c
}

//Onmousewheel - 
func (c *ComponentTag) Onmousewheel(value string) *ComponentTag {
	c.AddAttribute("onmousewheel", value)
	return c
}

//Oncontextmenu - 
func (c *ComponentTag) Oncontextmenu(value string) *ComponentTag {
	c.AddAttribute("oncontextmenu", value)
	return c
}

//Onerror - 
func (c *ComponentTag) Onerror(value string) *ComponentTag {
	c.AddAttribute("onerror", value)
	return c
}

//Onmousedown - 
func (c *ComponentTag) Onmousedown(value string) *ComponentTag {
	c.AddAttribute("onmousedown", value)
	return c
}

//Onplay - 
func (c *ComponentTag) Onplay(value string) *ComponentTag {
	c.AddAttribute("onplay", value)
	return c
}

//Ontimeupdate - 
func (c *ComponentTag) Ontimeupdate(value string) *ComponentTag {
	c.AddAttribute("ontimeupdate", value)
	return c
}

//Onvolumechange - 
func (c *ComponentTag) Onvolumechange(value string) *ComponentTag {
	c.AddAttribute("onvolumechange", value)
	return c
}

//Ondragover - 
func (c *ComponentTag) Ondragover(value string) *ComponentTag {
	c.AddAttribute("ondragover", value)
	return c
}

//Oncancel - 
func (c *ComponentTag) Oncancel(value string) *ComponentTag {
	c.AddAttribute("oncancel", value)
	return c
}

//Oninvalid - 
func (c *ComponentTag) Oninvalid(value string) *ComponentTag {
	c.AddAttribute("oninvalid", value)
	return c
}

//Onkeyup - 
func (c *ComponentTag) Onkeyup(value string) *ComponentTag {
	c.AddAttribute("onkeyup", value)
	return c
}

//Onblur - 
func (c *ComponentTag) Onblur(value string) *ComponentTag {
	c.AddAttribute("onblur", value)
	return c
}

//Ondurationchange - 
func (c *ComponentTag) Ondurationchange(value string) *ComponentTag {
	c.AddAttribute("ondurationchange", value)
	return c
}

//Onsuspend - 
func (c *ComponentTag) Onsuspend(value string) *ComponentTag {
	c.AddAttribute("onsuspend", value)
	return c
}

//Onclick - 
func (c *ComponentTag) Onclick(value string) *ComponentTag {
	c.AddAttribute("onclick", value)
	return c
}

//Onended - 
func (c *ComponentTag) Onended(value string) *ComponentTag {
	c.AddAttribute("onended", value)
	return c
}

//Oninput - 
func (c *ComponentTag) Oninput(value string) *ComponentTag {
	c.AddAttribute("oninput", value)
	return c
}

//Onmouseleave - 
func (c *ComponentTag) Onmouseleave(value string) *ComponentTag {
	c.AddAttribute("onmouseleave", value)
	return c
}

//Onautocomplete - 
func (c *ComponentTag) Onautocomplete(value string) *ComponentTag {
	c.AddAttribute("onautocomplete", value)
	return c
}

//Oncuechange - 
func (c *ComponentTag) Oncuechange(value string) *ComponentTag {
	c.AddAttribute("oncuechange", value)
	return c
}

//Onstalled - 
func (c *ComponentTag) Onstalled(value string) *ComponentTag {
	c.AddAttribute("onstalled", value)
	return c
}

//Onabort - 
func (c *ComponentTag) Onabort(value string) *ComponentTag {
	c.AddAttribute("onabort", value)
	return c
}

//Onratechange - 
func (c *ComponentTag) Onratechange(value string) *ComponentTag {
	c.AddAttribute("onratechange", value)
	return c
}

//Onemptied - 
func (c *ComponentTag) Onemptied(value string) *ComponentTag {
	c.AddAttribute("onemptied", value)
	return c
}

//Ondrop - 
func (c *ComponentTag) Ondrop(value string) *ComponentTag {
	c.AddAttribute("ondrop", value)
	return c
}

//Onfocus - 
func (c *ComponentTag) Onfocus(value string) *ComponentTag {
	c.AddAttribute("onfocus", value)
	return c
}

//Onkeypress - 
func (c *ComponentTag) Onkeypress(value string) *ComponentTag {
	c.AddAttribute("onkeypress", value)
	return c
}

//Onloadstart - 
func (c *ComponentTag) Onloadstart(value string) *ComponentTag {
	c.AddAttribute("onloadstart", value)
	return c
}

//Onprogress - 
func (c *ComponentTag) Onprogress(value string) *ComponentTag {
	c.AddAttribute("onprogress", value)
	return c
}

//Onsubmit - 
func (c *ComponentTag) Onsubmit(value string) *ComponentTag {
	c.AddAttribute("onsubmit", value)
	return c
}

//Ontoggle - 
func (c *ComponentTag) Ontoggle(value string) *ComponentTag {
	c.AddAttribute("ontoggle", value)
	return c
}

//Ondragenter - 
func (c *ComponentTag) Ondragenter(value string) *ComponentTag {
	c.AddAttribute("ondragenter", value)
	return c
}

//Onwaiting - 
func (c *ComponentTag) Onwaiting(value string) *ComponentTag {
	c.AddAttribute("onwaiting", value)
	return c
}

//Ondragleave - 
func (c *ComponentTag) Ondragleave(value string) *ComponentTag {
	c.AddAttribute("ondragleave", value)
	return c
}

//Onpause - 
func (c *ComponentTag) Onpause(value string) *ComponentTag {
	c.AddAttribute("onpause", value)
	return c
}

//Oncanplaythrough - 
func (c *ComponentTag) Oncanplaythrough(value string) *ComponentTag {
	c.AddAttribute("oncanplaythrough", value)
	return c
}

//Onload - 
func (c *ComponentTag) Onload(value string) *ComponentTag {
	c.AddAttribute("onload", value)
	return c
}

//Onresize - 
func (c *ComponentTag) Onresize(value string) *ComponentTag {
	c.AddAttribute("onresize", value)
	return c
}

//Ondrag - 
func (c *ComponentTag) Ondrag(value string) *ComponentTag {
	c.AddAttribute("ondrag", value)
	return c
}

//Onplaying - 
func (c *ComponentTag) Onplaying(value string) *ComponentTag {
	c.AddAttribute("onplaying", value)
	return c
}

//Onautocompleteerror - 
func (c *ComponentTag) Onautocompleteerror(value string) *ComponentTag {
	c.AddAttribute("onautocompleteerror", value)
	return c
}

//Ondblclick - 
func (c *ComponentTag) Ondblclick(value string) *ComponentTag {
	c.AddAttribute("ondblclick", value)
	return c
}

//Ondragstart - 
func (c *ComponentTag) Ondragstart(value string) *ComponentTag {
	c.AddAttribute("ondragstart", value)
	return c
}

//Onmouseout - 
func (c *ComponentTag) Onmouseout(value string) *ComponentTag {
	c.AddAttribute("onmouseout", value)
	return c
}

//Onseeked - 
func (c *ComponentTag) Onseeked(value string) *ComponentTag {
	c.AddAttribute("onseeked", value)
	return c
}

//Onchange - 
func (c *ComponentTag) Onchange(value string) *ComponentTag {
	c.AddAttribute("onchange", value)
	return c
}

//Onmousemove - 
func (c *ComponentTag) Onmousemove(value string) *ComponentTag {
	c.AddAttribute("onmousemove", value)
	return c
}

//Onmouseover - 
func (c *ComponentTag) Onmouseover(value string) *ComponentTag {
	c.AddAttribute("onmouseover", value)
	return c
}

//Onmouseup - 
func (c *ComponentTag) Onmouseup(value string) *ComponentTag {
	c.AddAttribute("onmouseup", value)
	return c
}

//Onscroll - 
func (c *ComponentTag) Onscroll(value string) *ComponentTag {
	c.AddAttribute("onscroll", value)
	return c
}

//Oncanplay - 
func (c *ComponentTag) Oncanplay(value string) *ComponentTag {
	c.AddAttribute("oncanplay", value)
	return c
}

//Ondragend - 
func (c *ComponentTag) Ondragend(value string) *ComponentTag {
	c.AddAttribute("ondragend", value)
	return c
}

//Onshow - 
func (c *ComponentTag) Onshow(value string) *ComponentTag {
	c.AddAttribute("onshow", value)
	return c
}

//Onsort - 
func (c *ComponentTag) Onsort(value string) *ComponentTag {
	c.AddAttribute("onsort", value)
	return c
}

//Onclose - 
func (c *ComponentTag) Onclose(value string) *ComponentTag {
	c.AddAttribute("onclose", value)
	return c
}
