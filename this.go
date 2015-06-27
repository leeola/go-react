package react

import (
	"errors"

	"github.com/gopherjs/gopherjs/js"
)

type This struct {
	this *js.Object
}

func (t *This) SetThis(this *js.Object) error {
	if t.this != nil {
		return errors.New(
			"This.SetThis Error: SetThis() can only be assigned once")
	}
	t.this = this
	return nil
}

// TODO: (Maybe) Add a callback to the js.Call, and block for it's return.
// Why maybe? Well, the blocking call may not be worth it. Especially if
// ForceUpdate ends up being how we render in go-react
func (t *This) ForceUpdate() error {
	if t.this == nil {
		return errors.New(
			"ForceUpdate Error: This.this pointer has not been assigned")
	}
	t.this.Call("forceUpdate")
	return nil
}

func (t *This) ComponentDidMount()    {}
func (t *This) ComponentWillUnmount() {}
func (t *This) Children() Components  { return nil }
func (t *This) Content() string       { return "" }
func (t *This) Render() Component     { return nil }
func (t *This) Props() Props          { return nil }
func (t *This) Tag() string           { return "" }
