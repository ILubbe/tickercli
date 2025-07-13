package colors

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func DetermineColor(value float64) string {
	switch {
	case value > 0:
		return Green
	case value < 0:
		return Red
	default:
		return Yellow
	}
}
