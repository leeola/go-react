package react

import "github.com/gopherjs/gopherjs/js"

type Component interface {
	SetThis(*js.Object) error

	// See: https://facebook.github.io/react/docs/component-api.html#forceupdate
	ForceUpdate() error

	// See: https://facebook.github.io/react/docs/component-specs.html#mounting-componentdidmount
	ComponentDidMount()

	// See: https://facebook.github.io/react/docs/component-specs.html#unmounting-componentwillunmount
	ComponentWillUnmount()

	Render() Component

	Children() Components
	Content() string
	Props() Props
	Tag() string
}

// Class is a function that takes Props and Components, returning a new
// Component. It does not provide the actual implmentation of a Component,
// but rather it is a loose specification to easily create instances of
// components.
//
// In other words, you should create a function for each of your
// components, with your desired component name for users of your
// component to cleanly create instances of it.
//
// To go into rather absurd detail, lets look at some examples.
// In essence, a Class is similar to JSX. In JSX, we could use:
//
// 		<div></div>
//
// To create a div instance. In go-react, we would use:
//
// 		Div()
//
// Still not clear?, lets say we have a hand made html Component. It may look
// something like this:
//
//		type Div struct {
//			props    Props
//			children []Component
//		}
//		func (c *Div) Render() []Component {
//			return c.children
//		}
//		func (c *Div) Tag() string { return "div" }
//		func (c *Div) Props() string { return nil }
//		func (c *Div) Content() string { return "" }
//
// If we had a whole series of html elements like this, and wanted to
// initialize them for actual use, it might look like this:
//
// 		Div{props: nil, children: []Component{
//			Header{props: nil, children: []Component{
//				H1{props: Props{"className": "title"}, children: []Component{
//					Span{props: nil, []Component{react.Content("My Website!")}},
//					},
//				},
//				Ul{props: nil, children: []Component{}},
//			}
// 		}
//
// Now lets compare this to what the JSX would look like:
//
//		<div>
//			<header>
//				<h1 className="title">
//					<span>My Website!</span>
//				</h1>
//				<ul></ul>
//		</div>
//
// The JSX presents a far more readable representation of the given
// components. To emulate that nice JSX, we can create functions that
// initialize our Go Components. Visually, it will look far closer to
// the JSX. Example:
//
// 		Div(nil,
//			Header(nil,
//				H1(react.Props{"className": "title"},
//					Span(nil, react.Content("My Website!")),
//				),
//				Ul(),
//			),
// 		)
//
// Both of the Go examples create the same Components, they're just
// presented in a much cleaner way. By providing a function of type Class
// for your own Components, you can drastically improve the cleanliness
// of your component creations.
type Class func(Props, ...Component) Component

type Components []Component
