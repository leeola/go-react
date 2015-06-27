package reactjs

import "github.com/gopherjs/gopherjs/js"

var React *js.Object
var Document *js.Object

func init() {
	React = js.Global.Get("React")
	Document = js.Global.Get("document")
}

func CreateClass(obj *js.Object) *js.Object {
	return React.Call("createClass", obj)
}

func CreateElement(args ...interface{}) *js.Object {
	return React.Call("createElement", args...)
}

func Render(elm *js.Object, dom *js.Object) {
	React.Call("render", elm, dom)
}

func RenderToDomId(elm *js.Object, domId string) {
	Render(elm, Document.Call("getElementById", domId))
}
