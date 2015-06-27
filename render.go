package react

import "github.com/leeola/go-react/reactjs"

func RenderToDomId(c Component, domId string) error {
	wrap, err := wrapComponent(c)
	if err != nil {
		return err
	}

	reactjs.RenderToDomId(reactjs.CreateElement(wrap), domId)

	return nil
}
