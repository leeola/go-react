//
// This Go package is comparable to the following React JSX:
//
//		React.render(
//			<h1>Hello, world!</h1>,
//			document.getElementById('example')
//		);
//
package main

import "github.com/leeola/go-react"

func main() {
	react.RenderToDomId(react.Div(nil,
		react.H1(nil, react.Content("Hello, World!"))),
		"example",
	)
}
