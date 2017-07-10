package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInput(t *testing.T) {
	spec := `
	app TestApp this is a test

	opt string_opt    -s --string   =value    :string     $STRING_OPT    The string opt description
	opt int_opt       -i --int      =2000     :int        $INT_OPT       The int opt description
	opt bool_opt      -b --bool     =false    :bool       $BOOL_OPT      The bool opt description
	opt strings_opt   -x --strings  =a,b,c    :strings    $STRINGS_OPT   The strings opt description
	opt ints_opt      -j --ints     =1,2,3    :ints       $INTS_OPT      The ints opt description

	arg string_arg                  =value    :string     $STRING_ARG    The string arg description
	arg int_arg                     =2000     :int        $INT_ARG       The int arg description
	arg bool_arg                    =false    :bool       $BOOL_ARG      The bool arg description
	arg strings_arg                 =a,b,c    :strings    $STRINGS_ARG   The strings arg description
	arg ints_arg                    =1,2,3    :ints       $INTS_ARG      The ints arg description

	spec -o [FILE...]
	`

	app, err := parse(spec)
	require.NoError(t, err)

	require.Equal(t, "TestApp", app.Name)
	require.Equal(t, "this is a test", app.Desc)
	require.Equal(t, "-o [FILE...]", app.Spec)

	// Options
	require.Len(t, app.Opts, 5)
	require.Equal(t, decl{
		Var:     "string_opt",
		Names:   []string{"-s", "--string"},
		Type:    "string",
		Default: "value",
		Env:     "STRING_OPT",
		Desc:    "The string opt description",
	}, *app.Opts[0])
	require.Equal(t, decl{
		Var:     "int_opt",
		Names:   []string{"-i", "--int"},
		Type:    "int",
		Default: 2000,
		Env:     "INT_OPT",
		Desc:    "The int opt description",
	}, *app.Opts[1])
	require.Equal(t, decl{
		Var:     "bool_opt",
		Names:   []string{"-b", "--bool"},
		Type:    "bool",
		Default: false,
		Env:     "BOOL_OPT",
		Desc:    "The bool opt description",
	}, *app.Opts[2])
	require.Equal(t, decl{
		Var:     "strings_opt",
		Names:   []string{"-x", "--strings"},
		Type:    "strings",
		Default: []string{"a", "b", "c"},
		Env:     "STRINGS_OPT",
		Desc:    "The strings opt description",
	}, *app.Opts[3])
	require.Equal(t, decl{
		Var:     "ints_opt",
		Names:   []string{"-j", "--ints"},
		Type:    "ints",
		Default: []int{1, 2, 3},
		Env:     "INTS_OPT",
		Desc:    "The ints opt description",
	}, *app.Opts[4])

	// Args
	require.Len(t, app.Args, 5)
	require.Equal(t, decl{
		Var:     "string_arg",
		Names:   []string{"STRING_ARG"},
		Type:    "string",
		Default: "value",
		Env:     "STRING_ARG",
		Desc:    "The string arg description",
	}, *app.Args[0])
	require.Equal(t, decl{
		Var:     "int_arg",
		Names:   []string{"INT_ARG"},
		Type:    "int",
		Default: 2000,
		Env:     "INT_ARG",
		Desc:    "The int arg description",
	}, *app.Args[1])
	require.Equal(t, decl{
		Var:     "bool_arg",
		Names:   []string{"BOOL_ARG"},
		Type:    "bool",
		Default: false,
		Env:     "BOOL_ARG",
		Desc:    "The bool arg description",
	}, *app.Args[2])
	require.Equal(t, decl{
		Var:     "strings_arg",
		Names:   []string{"STRINGS_ARG"},
		Type:    "strings",
		Default: []string{"a", "b", "c"},
		Env:     "STRINGS_ARG",
		Desc:    "The strings arg description",
	}, *app.Args[3])
	require.Equal(t, decl{
		Var:     "ints_arg",
		Names:   []string{"INTS_ARG"},
		Type:    "ints",
		Default: []int{1, 2, 3},
		Env:     "INTS_ARG",
		Desc:    "The ints arg description",
	}, *app.Args[4])

}
