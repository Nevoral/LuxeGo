package lx

type Content interface {
	AddAttribute(name, value string)
	Render(level int) string
}
