package gqlgentype

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/util/typeconvert"
)

// MarshalInt GoのintをGraphQLのIntに変換する。intが32bit符号付き整数に収まらない場合はpanicになる
func MarshalInt(i int) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		//nolint:errcheck // gqlgenのinterfaceに準拠
		io.WriteString(w, strconv.Itoa(int(typeconvert.ToMustInt32(i))))
	})
}

func UnmarshalInt(v any) (int, error) {
	i, err := graphql.UnmarshalInt(v)
	if err != nil {
		return 0, fmt.Errorf("int型に変換できませんでした: %w", err)
	}
	i32, err := typeconvert.ToInt32(i)
	if err != nil {
		return 0, err
	}
	return int(i32), nil
}

func MarshalInt64(i int64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		//nolint:errcheck // gqlgenのinterfaceに準拠
		io.WriteString(w, strconv.Quote(strconv.FormatInt(i, 10)))
	})
}

func UnmarshalInt64(v any) (int64, error) {
	if tmpStr, ok := v.(string); ok {
		i, err := strconv.ParseInt(tmpStr, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("int64型に変換できませんでした: %w", err)
		}
		return i, nil
	}
	return 0, fmt.Errorf("int64型に変換できませんでした")
}
