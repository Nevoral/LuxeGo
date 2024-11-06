package base

import (
	"context"
	"io"

	"github.com/Nevoral/LuxeGo/pkg/SpecGenzo/model"
)

type Content interface {
	GetName() string
	GetNamespace() model.Namespace
	Render(ctx context.Context, writer io.Writer, renderFormat RenderFormat) error

	getTag() *node
	getAttributes() *attributes
	getNodeType() model.NodeType
	getChildren() []*node
	getNestingLevel() int
	updateNestingLevel(offset int)
	// getSupportedAttributes() map[string][]string
	// getParent() *node
	// getRenderFormat() RenderFormat
	// setNamespace(namespace spec.Namespace)
	// setTextContent(content string)
	// registerAttribute(name string, boolean bool, value string)
	// registerAttributes(attributes ...*Attribute)
	// registerChildren(children ...*node)
	// registerParent(parent *node)
	// updateParent(parent *node)
	// updateRenderFormat(fn RenderFormat)
	// Children(tags ...any) Content
}

type node struct {
	name       string
	attributes *attributes
	children   []*node
	*rendererData
}

func NewDoctypeNode(namespace model.Namespace) Content {
	return &node{
		name:       "!DOCTYPE",
		attributes: nil,
		children:   nil,
		rendererData: &rendererData{
			nodeType:    model.DoctypeType,
			namespace:   namespace,
			nestedLevel: 0,
		},
	}
}

func NewSelfClosingNode(name string, namespace model.Namespace, attributes Attributes) Content {
	return &node{
		name:       name,
		attributes: attributes.getAttributes(),
		children:   nil,
		rendererData: &rendererData{
			nodeType:    model.SelfClosingType,
			namespace:   namespace,
			nestedLevel: 0,
		},
	}
}

func NewFullTagNode(name string, namespace model.Namespace, attributes Attributes, children *[]Content) Content {
	return &node{
		name:       name,
		attributes: attributes.getAttributes(),
		children:   extractNodes(children),
		rendererData: &rendererData{
			nodeType:    model.FullTagType,
			namespace:   namespace,
			nestedLevel: 0,
		},
	}
}

func NewTextNode(nodeType model.NodeType, namespace model.Namespace, textContent string) Content {
	return &node{
		name:       textContent,
		attributes: nil,
		children:   nil,
		rendererData: &rendererData{
			nodeType:    nodeType,
			namespace:   namespace,
			nestedLevel: 0,
		},
	}
}

func (n *node) GetName() string {
	return n.name
}

func (n *node) GetNamespace() model.Namespace {
	return n.namespace
}

func (n *node) Render(ctx context.Context, writer io.Writer, renderFormat RenderFormat) error {
	n.updateNestingLevel(n.nestedLevel)
	return renderFormat.render(ctx, writer, n)
}

func (n *node) getTag() *node {
	return n
}

func (n *node) getAttributes() *attributes {
	return n.attributes
}

func (n *node) getNodeType() model.NodeType {
	return n.nodeType
}

func (n *node) getChildren() []*node {
	return n.children
}

func (n *node) getNestingLevel() int {
	return n.nestedLevel
}

func (n *node) updateNestingLevel(offset int) {
	n.nestedLevel = offset
	for _, child := range n.children {
		child.nestedLevel = offset + 1
		child.updateNestingLevel(offset + 1)
	}
}

func extractNodes(tags *[]Content) []*node {
	var nodes []*node
	for _, tag := range *tags {
		nodes = append(nodes, tag.getTag())
	}
	return nodes
}

// func (t *node) getParent() *node {
// return t.parent
// }
// func (t *node) getRenderFormat() RenderFormat {
// return t.renderFormat
// }
//
// func (t *node) setNamespace(namespace spec.Namespace) {
// t.namespace = namespace
// }
//
// func (t *node) setTextContent(content string) {
// t.textContent = content
// }
// func (t *node) registerAttribute(name string, boolean bool, value string) {
// if boolean {
// t.attributes = append(t.attributes, &Attribute{
// Name:    name,
// Value:   "",
// Boolean: true,
// })
// return
// }
//
// t.attributes = append(t.attributes, &Attribute{
// Name:    name,
// Value:   value,
// Boolean: false,
// })
// }

// func (t *node) registerAttributes(attributes ...*Attribute) {
// if t.attributes == nil {
// t.attributes = make([]*Attribute, 0)
// }
// for _, attribute := range attributes {
// t.attributes = append(t.attributes, &Attribute{
// Name:          attribute.Name,
// Value:         attribute.Value,
// Boolean:       attribute.Boolean,
// attributeType: custom,
// })
// }
// }
// func (t *node) registerChildren(children ...*node) {
// 	if t.children == nil {
// 		t.children = make([]*node, 0)
// 	t.children = append(t.children, children...)
// 	for _, child := range children {
// 		child.registerParent(t)
// 	}
// }

// func (t *node) registerParent(parent *node) {
// t.parent = parent
// }
//
// func (t *node) updateParent(parent *node) {
// t.parent = parent
// for _, child := range t.children {
// child.updateParent(t)
// }
// }

//
// func (t *node) updateRenderFormat(fn RenderFormat) {
// t.renderFormat = fn
// for _, child := range t.children {
// child.updateRenderFormat(t.renderFormat)
// }
// }

// func (t *node) isSupported(name string) bool {
// if t.supportedAttributes == nil || len(t.supportedAttributes) == 0 {
// return false
// }
//
// for _, attributes := range t.supportedAttributes {
// if slices.Contains(attributes, name) {
// return true
// }
// }
// return false
// }
//
// func (t *node) getCategoryName(name string) string {
// for category, attributes := range t.supportedAttributes {
// if slices.Contains(attributes, name) {
// return category
// }
// }
// panic("Error: this should not happened, because this method is called after was checked that this attribute is Supported by this tag.")
// }
