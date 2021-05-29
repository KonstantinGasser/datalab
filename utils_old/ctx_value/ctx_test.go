package ctx_value

import (
	"context"
	"testing"

	"github.com/matryer/is"
)

func TestAddValue(t *testing.T) {
	is := is.New(t)
	tt := []struct {
		have interface{}
		want interface{}
	}{
		{have: "string", want: "string"},
		{have: 24, want: 24},
		{have: map[string]interface{}{"hello": "friend"}, want: map[string]interface{}{"hello": "friend"}},
	}

	for _, tc := range tt {
		freshCtx := context.Background()
		newCtx := AddValue(freshCtx, "test", tc.have)
		ctxVal := newCtx.Value(ctxKey("test"))
		is.Equal(ctxVal, tc.want)
	}
}

func TestGetString(t *testing.T) {
	is := is.New(t)
	tt := []struct {
		have interface{}
		want interface{}
	}{
		{have: nil, want: ""},
		{have: "stri ng2", want: "stri ng2"},
		{have: "string3", want: "string3"},
	}

	for _, tc := range tt {
		freshCtx := context.WithValue(context.Background(), ctxKey("test"), tc.have)
		value := GetString(freshCtx, "test")
		is.Equal(value, tc.want)
	}

}
