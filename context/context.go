package context

import "context"

type Key string

// String converts Key into built-in string
func (k Key) String() string {
	return string(k)
}

// GetString gets a value for provided key and asserts it into string
func GetString(ctx context.Context, k Key) string {
	val, _ := ctx.Value(k).(string)
	return val
}
