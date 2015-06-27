package react

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/leeola/go-react/reactjs"
)

// wrapComponent creates a React class for the given Component by
// creating a plain js.Object and mapping functions to it, and back.
//
// For example, the React `render()` method is assigned to the object.
// When React calls this method, the given Component's Render() method is
// called. The result of Component.Render() is recursively run through
// Wrap() as well, with the final result being a series of objects given
// to React that map to their respective components.
//
// This means that when a component calls the embedded *React methods,
// such as React.SetState() end up calling the functions in this closure
// which then call the normal React `this.setState()` methods.
//
// A lot of wiring, but the goal is that complexity in a Component is
// isolated from the complexity given to React. So all of the tomfoolery
// that React does to method Binding's are not actually done to the
// methods of your component.
//
// TL;DR: The overhead caused by Wrap() is to avoid React tomfoolery
// multiplying by GopherJS tomfoolery. Resulting in a (in theory)
// net tomfoolery reduction. Science.
//
// TODO: Create the *React pipeline to communicate from the Component
// to the jsObject wrapper.
//
// TODO: Write errors to console.error
//
// TODO: Drop error return. Don't think it's possible to have an error
// outside of the callbacks.
//
// TODO: Handle the following React warning:
//
// 		Warning: Each child in an array or iterator should have a
// 		unique "key" prop. Check the React.render call using <div>.
// 		See https://fb.me/react-warning-keys for more information.
//
// Though, the responsibility of handling this may be better off on the
// Component.
func wrapComponent(c Component) (jsComp *js.Object, err error) {
	fmt.Printf("wrapComponent(<%s>)\n", c.Tag())
	jsObject := js.Global.Get("Object").New()

	// This is a reference to the final, valid React object instance.
	//
	// React does a lot of binding/this tomfoolery, and it heavily
	// screws with GopherJS. As a result, custom functions won't have
	// valid `this` objects.
	//
	// React API functions can get a valid `this`, such as `render` and
	// `setInitialState`, but custom functions fail.
	//
	// Commented out, as it's likely being removed soon (no need for
	// the reference, Components have SetThis()
	//var this *js.Object

	// TODO: Find a way to support sane error handling. It's difficult
	// because the render() method is a callback from JS, and the return
	// value is consumed by React.js itself. Furthermore, even if we could
	// return an error value, we can't get access to it.
	render := func() *js.Object {
		fmt.Printf("<%s>.render\n", c.Tag())

		rendComp := c.Render()

		var jsClass *js.Object
		var content string
		// If a component is Rendered, wrap it and return.
		if rendComp != nil {
			// If this child has Content (String), return it directly.
			content = rendComp.Content()
			if content != "" {
				return reactjs.CreateElement(c.Tag(), c.Props(), content)
			}

			jsClass, _ = wrapComponent(rendComp)

			return reactjs.CreateElement(jsClass)
		}

		// The comp did not render anything, so try to create a Tag with
		// it's children.
		//
		// If it has no tag, we can't do that. Error out.
		if c.Tag() == "" {
			fmt.Println("Error: Component has no Children and no Tag")
			return nil
		}

		children := c.Children()
		childrenCount := len(children)

		// If this comp has no children, create an empty jsComp (eg: <foo />)
		if childrenCount == 0 {
			return reactjs.CreateElement(c.Tag(), c.Props())
		}

		// The == 1 check is done to avoid creating the slice, looping,
		// and (most importantly) avoid returning an array. By doing it
		// here, we can save a bit of overhead.. in theory.
		//
		// Premature optimization without benchmarks? heh
		var childComp Component
		if childrenCount == 1 {
			childComp = children[0]

			// If this child has Content (String), return it directly.
			content = childComp.Content()
			if content != "" {
				return reactjs.CreateElement(c.Tag(), c.Props(), content)
			}

			jsClass, _ = wrapComponent(childComp)

			return reactjs.CreateElement(c.Tag(), c.Props(),
				reactjs.CreateElement(jsClass))
		}

		// A slice of this Element's Childen. Note that in the event of
		// a Content child, string is added to this slice instead of a
		// jsObject.
		jsChildren := make([]interface{}, childrenCount)

		var i int
		for i, childComp = range children {
			// If this child has Content (String), push it onto Children
			// directly.
			content = childComp.Content()
			if content != "" {
				jsChildren[i] = content
				continue
			}

			// Wrap the component, returning it's class
			jsClass, _ = wrapComponent(childComp)
			// And create an element from the class, adding it to the slice.
			jsChildren[i] = reactjs.CreateElement(jsClass)
		}

		return reactjs.CreateElement(c.Tag(), c.Props(), jsChildren)
	}

	jsObject.Set("render", render)
	jsObject.Set("componentDidMount", c.ComponentDidMount)
	jsObject.Set("componentWillUnmount", c.ComponentWillUnmount)

	// Temporary implementation for early testing.
	jsObject.Set("getInitialState", js.MakeFunc(
		func(reactThis *js.Object, _ []*js.Object) interface{} {
			if err := c.SetThis(reactThis); err != nil {
				fmt.Println("Error: ", err.Error())
			}
			return nil
		},
	))

	return reactjs.CreateClass(jsObject), nil
}
