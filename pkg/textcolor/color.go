package textcolor

import "errors"

const (
	RED = iota
	GREEN
	YELLOW
	BLUE
	PURPLE
	CYAN
	WHITE
	RESET
)

var colors = map[color]string{
	RED:    "\033[31m",
	GREEN:  "\033[32m",
	YELLOW: "\033[33m",
	BLUE:   "\033[34m",
	PURPLE: "\033[35m",
	CYAN:   "\033[36m",
	WHITE:  "\033[37m",
	RESET:  "\033[0m",
}

type color int

func AddColor(s string, c color) (string, error) {
	if c > RESET {
		return s, errors.New("invalid color")
	}
	return colors[c] + s + colors[RESET], nil
}
