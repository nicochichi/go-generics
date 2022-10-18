package convert

import (
	"strconv"

	"nicochichi/generics/domain"
)

type (
	IDs interface {
		uint64 | int64 | domain.CustomerID | domain.ShipmentID
	}
)

const (
	base10 = 10
)

func ToString[T IDs](x T) string {
	return strconv.FormatUint(uint64(x), base10)
}
