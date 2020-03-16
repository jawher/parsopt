package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseError is returned when the provided spec is invalid
type ParseError struct {
	Line    string
	Row     int
	Message string
}

const (
	dirAPP  = "app"
	dirOPT  = "opt"
	dirARG  = "arg"
	dirSPEC = "spec"
)

type app struct {
	Name string
	Desc string
	Spec string
	Opts []*decl
	Args []*decl
}

type decl struct {
	Var     string
	Type    string
	Names   []string
	Env     string
	Default interface{}
	Desc    string

	storage interface{}
}

var (
	types = []string{
		"string",
		"strings",
		"int",
		"ints",
		"bool",
	}
)

func (e *ParseError) Error() string {
	return fmt.Sprintf("Parse error at row %d: %s\nOffending line:\n%s", e.Row, e.Message, e.Line)
}

func parse(directives string) (*app, error) {
	lines := strings.Split(directives, "\n")
	app := &app{}

	for row, line := range lines {
		line := strings.TrimSpace(line)
		if line == "" {
			continue
		}

		//Syntax:
		// app app_name [desc]
		// opt var_name [:type] [=value] [$env] -o --offer [desc]
		// arg var_name  [:type] [=value] [$env] [desc]
		// spec spec_string
		parts := words(line)
		if len(parts) < 2 {
			return nil, mkParseError(row, line, "Invalid syntax: expected DIRECTIVE ARG...")
		}

		directive := parts[0]

		switch directive {
		case dirAPP:
			err := parseApp(parts, app)
			if err != nil {
				return nil, err
			}
		case dirSPEC:
			err := parseSpec(parts, app)
			if err != nil {
				return nil, err
			}
		case dirOPT:
			decl, err := parseOption(row, line, parts)
			if err != nil {
				return nil, err
			}
			app.Opts = append(app.Opts, decl)
		case dirARG:
			decl, err := parseArg(row, line, parts)
			if err != nil {
				return nil, err
			}
			app.Args = append(app.Args, decl)
		default:
			return nil, mkParseError(row, line, "Unknown directive %q", directive)
		}
	}

	return app, nil
}

func parseApp(parts []string, app *app) error {
	app.Name = parts[1]
	app.Desc = strings.Join(parts[2:], " ")
	return nil
}

func parseSpec(parts []string, app *app) error {
	app.Spec = strings.Join(parts[1:], " ")
	return nil
}

func parseOption(row int, line string, parts []string) (*decl, error) {
	// opt var_name [:type] [=value] [$env] -o --offer [desc]
	res := &decl{
		Var: parts[1],
	}
	stop := false
	defaultValue := ""
	for idx, word := range parts[2:] {
		if stop {
			break
		}
		switch {
		case strings.HasPrefix(word, ":"):
			res.Type = word[1:]
		case strings.HasPrefix(word, "$"):
			res.Env = word[1:]
		case strings.HasPrefix(word, "="):
			defaultValue = word[1:]
		case strings.HasPrefix(word, "-"):
			res.Names = append(res.Names, word)
		default:
			res.Desc = strings.Join(parts[2+idx:], " ")
			stop = true
		}
	}

	switch {
	case len(res.Names) == 0:
		return nil, mkParseError(row, line, "At least one option name must be provided")
	case res.Type == "":
		res.Type = "string"
	case !inList(res.Type, types):
		return nil, mkParseError(row, line, "Provided type %q is not one of: %s", res.Type, strings.Join(types, ", "))
	}

	dv, err := parseDefaultValue(res.Type, defaultValue)
	if err != nil {
		return nil, mkParseError(row, line, "Invalid default value %q", defaultValue)
	}
	res.Default = dv
	return res, nil
}

func parseArg(row int, line string, parts []string) (*decl, error) {
	// arg var_name [:type] [=value] [$env] [desc]

	res := &decl{
		Var:   parts[1],
		Names: []string{strings.ToUpper(parts[1])},
	}
	stop := false
	defaultValue := ""
	for idx, word := range parts[2:] {
		if stop {
			break
		}
		switch {
		case strings.HasPrefix(word, ":"):
			res.Type = word[1:]
		case strings.HasPrefix(word, "$"):
			res.Env = word[1:]
		case strings.HasPrefix(word, "="):
			defaultValue = word[1:]
		default:
			res.Desc = strings.Join(parts[2+idx:], " ")
			stop = true
		}
	}

	switch {
	case res.Type == "":
		res.Type = "string"
	case !inList(res.Type, types):
		return nil, mkParseError(row, line, "Provided type %q is not one of: %s", res.Type, strings.Join(types, ", "))
	}

	dv, err := parseDefaultValue(res.Type, defaultValue)
	if err != nil {
		fmt.Printf("Err=%v\n", err)
		return nil, mkParseError(row, line, "Invalid default value %q", defaultValue)
	}
	res.Default = dv

	return res, nil
}

func mkParseError(row int, line string, msg string, args ...interface{}) error {
	return &ParseError{
		Row:     row,
		Line:    line,
		Message: fmt.Sprintf(msg, args...),
	}
}

func words(line string) []string {
	line = strings.Replace(line, "\t", " ", -1)

	parts := strings.Fields(line)

	return append(make([]string, 0, len(parts)), parts...)
}

func inList(x string, xs []string) bool {
	for _, y := range xs {
		if x == y {
			return true
		}
	}
	return false
}

func parseDefaultValue(typ, def string) (res interface{}, err error) {
	switch typ {
	case "string":
		res, err = def, nil
		return
	case "int":
		if def == "" {
			res, err = 0, nil
			return
		}
		res, err = strconv.Atoi(def)
		return
	case "bool":
		if def == "" {
			res, err = false, nil
			return
		}
		res, err = strconv.ParseBool(def)
		return
	case "strings":
		if def == "" {
			res, err = []string(nil), nil
			return
		}
		res, err = strings.Split(def, ","), nil
		return
	case "ints":
		if def == "" {
			res, err = []int(nil), nil
			return
		}
		xs := strings.Split(def, ",")
		ixs := make([]int, 0, len(xs))
		for _, x := range xs {
			var ix int
			ix, err = strconv.Atoi(x)
			if err != nil {
				return
			}
			ixs = append(ixs, ix)
		}
		res, err = ixs, nil
		return
	default:
		panic(fmt.Sprintf("Unhandled type %q", typ))
	}
}
