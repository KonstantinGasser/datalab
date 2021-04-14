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
		yes, err := shoudlBind(tc.tag)
		if err != nil {
			is.NoErr(err)
		}
		is.True(tc.exp == yes)
	}
}

type tagfields struct {
	F1 string `bind:"yes"`
	F2 string `json:"field" bind:"yes"`
	F3 string
}

func TestMustCompile(t *testing.T) {
	is := is.New(t)

	tt := []struct {
		s   tagfields
		exp error
	}{
		{
			s: tagfields{
				F1: "test_f1",
				F2: "test_f2",
				F3: "",
			},
			exp: nil,
		},
		{
			s: tagfields{
				F1: "",
				F2: "test_f2",
			},
			exp: ErrMissingValues,
		},
		{
			s: tagfields{
				F1: "",
				F2: "",
				F3: "",
			},
			exp: ErrMissingValues,
		},
		{
			s: tagfields{
				F1: "test_f1",
				F2: "",
				F3: "",
			},
			exp: ErrMissingValues,
		},
	}

	for _, tc := range tt {
		err := MustBind(&tc.s)
		is.True(tc.exp == err)
	}
}
