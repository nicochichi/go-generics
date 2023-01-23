package convert

import (
	"strconv"
)

type (
	IDs interface {
		~uint64
	}
)

const (
	base10 = 10
)

func ToString[T IDs](x T) string {
	return strconv.FormatUint(uint64(x), base10)
}
