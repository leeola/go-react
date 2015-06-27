package react

type Props map[string]interface{}

func (p Props) ClassName() (s string, ok bool) {
	return p.GetString("className")
}

func (p Props) GetBool(key string) (v bool, ok bool) {
	v, ok = p[key].(bool)
	return v, ok
}

func (p Props) GetString(key string) (v string, ok bool) {
	v, ok = p[key].(string)
	return v, ok
}
