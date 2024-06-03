package lx

type Content interface {
	AddAttribute(name, value string)
	Render() string
}
