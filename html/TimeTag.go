package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Time -
func Time(tags ...interface{}) *TimeTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		case []LuxeGo.Content:
			children = append(children, v...)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	return &TimeTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "time", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type TimeTag struct {
	*ComponentHtmlTag
}

// Datetime - Specifies the date and time.
func (t *TimeTag) Datetime(value string) *TimeTag {
	t.AddAttribute("datetime", value)
	return t
}
