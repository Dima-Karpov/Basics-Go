package generativeTesting

import (
	"log"
	"strings"
)

func Pad(s string, max uint) string {
	log.Printf("Testing Len: %d, Str: %s\n", max, s)
	ln := uint(len(s))
	if ln > max {
		return s[:max-1]
	}
	s += strings.Repeat(" ", int(max-ln))
	return s
}
