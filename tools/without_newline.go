package test

func WithoutNewLine(s string) string {
	return s[:len(s)-1]
}
