package LuxeGo

import (
	"context"
	"io"
)

type Content interface {
	Render(ctx context.Context, writer io.Writer) error
	//AddAttribute(name, value string)
	//RenderString(level int) string
}
