package utils

type ContextKey string

func (c ContextKey) String() string {
	return "mykey" + string(c)
}
