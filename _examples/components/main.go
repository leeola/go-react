//
// This Go package is comparable to the following React JSX:
//
//		var CommentList = React.createClass({
//		  render: function() {
//		    return (
//		      <div className="commentList">
//		        Hello, world! I am a CommentList.
//		      </div>
//		    );
//		  }
//		});
//		var CommentForm = React.createClass({
//		  render: function() {
//		    return (
//		      <div className="commentForm">
//		        Hello, world! I am a CommentForm.
//		      </div>
//		    );
//		  }
//		});
//		var CommentBox = React.createClass({displayName: 'CommentBox',
//		  render: function() {
//		    return (
//		      <div className="commentBox">
//		        <h1>Comments</h1>
//		        <CommentList />
//		        <CommentForm />
//		      </div>
//		      )
//		    );
//		  }
//		});
//		React.render(
//		  React.createElement(CommentBox, null),
//		  document.getElementById('content')
//		);
//
package main

import . "github.com/leeola/go-react"

// Our main looks roughly the same as Hello World, except we're using
// a custom Class this time. CommentBox.
func main() {
	RenderToDomId(CommentBox(nil), "example")
}

// The CommentBox Class is being created with the help of a convinience
// function, FuncClass. FuncClass takes the signature of the Render()
// method and binds it to a struct, and then returns a working Class.
//
// This is optimal if you don't need the additional functionality that
// a full blown Struct gives you. See the FuncClass docstring for
// further explanation.
var CommentBox Class = FuncClass(func(_ Component) Component {
	return Div(Props{"className": "commentBox"},
		H1(nil, Content("Comments")),
		CommentList(),
		CommentForm(nil),
	)
})

// With this Component and Class, we're creating a valid Component, but
// a simple "Class". While the CommentList func does not match the type
// Class, we're still able to use it as a shorthand to create the
// Component.
//
// Because the Class doesn't really matter at all, it's just a shorthand,
// this is perfectly fine to do.
func CommentList() Component {
	return &commentList{This: &This{}}
}

type commentList struct{ *This }

func (c *commentList) Render() Component {
	return Div(Props{"className": "commentList"},
		Content("Hello, world! I am a CommentList."),
	)
}

// And finally, a full Class and Component. Nothing complex, but it
// contains a signature that users of your component will expect.
func CommentForm(p Props, c ...Component) Component {
	return &commentForm{This: &This{}, props: p, children: c}
}

type commentForm struct {
	*This
	props    Props
	children Components
}

func (c *commentForm) Render() Component {
	children := append(c.children,
		Content("Hello, world! I am a CommentForm."))

	return Div(Props{"className": "commentForm"}, children...)
}
