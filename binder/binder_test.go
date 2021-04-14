package binder

import (
	"testing"

	"github.com/matryer/is"
)

func TestShouldBind(t *testing.T) {
	is := is.New(t)

	tt := []struct {
		tag string
		exp bool
	}{
		{
			tag: `bind:"yes"`,
			exp: true,
		},
		{
			tag: ``,
			exp: false,
		},
		{
			tag: `bind:" yes"`,
			exp: false,
		},
		{
			tag: `bind:"y e s"`,
			exp: false,
		},
		{
			tag: `bind:"yes "`,
			exp: false,
		},
	}

	for _, tc := range tt {
		ok := shoudlBind(tc.tag)
		is.True(tc.exp == ok)
	}
}

type tagfields struct {
	F1 string      `bind:"yes, min=6, max=12"`
	F2 string      `json:"field" bind:"yes"`
	F3 interface{} `bind:"yes"`
}

func TestMustBind(t *testing.T) {
	is := is.New(t)

	tt := []struct {
		s   tagfields
		exp error
	}{
		{
			s: tagfields{
				F1: "test_f1",
				F2: "test_f2",
				F3: 42,
			},
			exp: nil,
		},
		{
			s: tagfields{
				F1: "",
				F2: "test_f2",
				F3: 42,
			},
			exp: ErrNoFullBind{invalid: []string{"F1"}},
		},
	}

	for _, tc := range tt {
		err := MustBind(&tc.s)
		is.True(tc.exp == err)
	}
}

// {
// 	s: tagfields{
// 		F1: "test_f1",
// 		F2: "test_f2",
// 		F3: int8(0),
// 	},
// 	exp: ErrMissingValues,
// },
// {
// 	s: tagfields{
// 		F1: "test_f1",
// 		F2: "test_f2",
// 		F3: int16(0),
// 	},
// 	exp: ErrMissingValues,
// },
// {
// 	s: tagfields{
// 		F1: "test_f1",
// 		F2: "test_f2",
// 		F3: int32(0),
// 	},
// 	exp: ErrMissingValues,
// },
// {
// 	s: tagfields{
// 		F1: "test_f1",
// 		F2: "test_f2",
// 		F3: int64(0),
// 	},
// 	exp: ErrMissingValues,
// },
// {
// 	s: tagfields{
// 		F1: "test_f1",
// 		F2: "test_f2",
// 		F3: uint(0),
// 	},
// 	exp: ErrMissingValues,
// },
// {
// 	s: tagfields{
// 		F1: "test_f1",
// 		F2: "test_f2",
// 		F3: uint8(0),
// 	},
// 	exp: ErrMissingValues,
// },
// {
// 	s: tagfields{
// 		F1: "test_f1",
// 		F2: "test_f2",
// 		F3: uint16(0),
// 	},
// 	exp: ErrMissingValues,
// },
// {
// 	s: tagfields{
// 		F1: "test_f1",
// 		F2: "test_f2",
// 		F3: uint32(0),
// 	},
// 	exp: ErrMissingValues,
// },
// {
// 	s: tagfields{
// 		F1: "test_f1",
// 		F2: "test_f2",
// 		F3: uint64(0),
// 	},
// 	exp: ErrMissingValues,
// },
