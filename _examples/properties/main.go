package main

import (
	"fmt"

	. "github.com/leeola/go-react"
)

func main() {
	RenderToDomId(
		Author(Props{
			"className": "moderator",
			"author":    "John Doe",
			"signature": "Tractors rock!",
		}), "example")
}

var Author Class = FuncClass(func(c Component) Component {
	props := c.Props()

	// For common React properties, Props provides methods to retrieve
	// them.
	className, _ := props.ClassName()

	// There are also convenience methods for retrieving types
	author, _ := props.GetString("author")

	// And of course, Props is just a map, so you can just get values
	// directly.
	signature, _ := props["signature"].(string)

	return Div(Props{"className": className},
		H1(Props{"className": "author"},
			Content(fmt.Sprintf("Author: %s", author))),
		Span(nil, Content(signature)),
	)
})
