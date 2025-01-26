package typeconvert

import (
	"fmt"
	"math"
)

type convertibleToInt32 interface {
	~int | ~int64
}

func ToPtr[T any](v T) *T {
	return &v
}


func ToInt32[T1 convertibleToInt32](t1 T1) (int32, error) {
	if t1 < math.MinInt32 || t1 > math.MaxInt32 {
		return 0, fmt.Errorf("32bit符号付き整数の範囲を超えています")
	}
	return int32(t1), nil
}

func ToMustInt32[T1 convertibleToInt32](t1 T1) int32 {
	i32, err := ToInt32(t1)
	if err != nil {
		panic(err)
	}
	return i32
}
