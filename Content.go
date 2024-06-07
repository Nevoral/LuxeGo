package LuxeGo

import (
	"context"
	"io"
)

type Content interface {
	AddAttribute(name, value string)
	Render(ctx context.Context, writer io.Writer) error
	RenderString(level int) string
}
