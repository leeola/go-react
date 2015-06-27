package react

// Content is a special Class that accepts a string and returns a
// component that will return the given string via `Content()`.
//
// For further explanation of this quirky Class, see the Content()
// docstring of the Component interface.
func Content(s string) Component {
	return &htmlComponent{This: &This{}, content: s}
}

// HtmlFactory is an easy func for creating plain Tag classes. This
// is used to create the entire base suite of Html Components in the
// react package.
func HtmlFactory(tag string) Class {
	return func(p Props, c ...Component) Component {
		return &htmlComponent{
			This:     &This{},
			children: c,
			props:    p,
			tag:      tag,
		}
	}
}

type htmlComponent struct {
	*This

	tag      string
	content  string
	props    Props
	children Components
}

func (c *htmlComponent) Children() Components { return c.children }
func (c *htmlComponent) Content() string      { return c.content }
func (c *htmlComponent) Props() Props         { return c.props }
func (c *htmlComponent) Tag() string          { return c.tag }

func FuncClass(fn func(Component) Component) Class {
	return func(p Props, c ...Component) Component {
		return &funcComponent{This: &This{}, fn: fn, props: p, children: c}
	}
}

type funcComponent struct {
	*This

	fn       func(Component) Component
	props    Props
	children Components
}

func (c *funcComponent) Children() Components { return c.children }
func (c *funcComponent) Props() Props         { return c.props }
func (c *funcComponent) Render() Component    { return c.fn(c) }
