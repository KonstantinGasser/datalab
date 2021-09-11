package ctxkey

type (
	KeyString string
)

func Str(str string) KeyString {
	return KeyString(str)
}
