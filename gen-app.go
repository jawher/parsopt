package main

import (
	"strings"

	"fmt"

	"flag"

	"os"

	cli "github.com/jawher/mow.cli"
)

func genApp(app *app) *cli.Cli {
	res := cli.App(app.Name, app.Desc)
	res.ErrorHandling = flag.ContinueOnError

	if app.Spec != "" {
		res.Spec = app.Spec
	}

	for _, opt := range app.Opts {
		switch opt.Type {
		case "bool":
			opt.storage = res.Bool(cli.BoolOpt{
				Name:   optNames(opt),
				Desc:   opt.Desc,
				EnvVar: opt.Env,
				Value:  opt.Default.(bool),
			})
		case "string":
			opt.storage = res.String(cli.StringOpt{
				Name:   optNames(opt),
				Desc:   opt.Desc,
				EnvVar: opt.Env,
				Value:  opt.Default.(string),
			})
		case "int":
			opt.storage = res.Int(cli.IntOpt{
				Name:   optNames(opt),
				Desc:   opt.Desc,
				EnvVar: opt.Env,
				Value:  opt.Default.(int),
			})
		case "strings":
			opt.storage = res.Strings(cli.StringsOpt{
				Name:   optNames(opt),
				Desc:   opt.Desc,
				EnvVar: opt.Env,
				Value:  opt.Default.([]string),
			})
		case "ints":
			opt.storage = res.Ints(cli.IntsOpt{
				Name:   optNames(opt),
				Desc:   opt.Desc,
				EnvVar: opt.Env,
				Value:  opt.Default.([]int),
			})
		default:
			panic(fmt.Sprintf("Unhandled type %q", opt.Type))
		}
	}

	for _, arg := range app.Args {
		switch arg.Type {
		case "bool":
			arg.storage = res.Bool(cli.BoolArg{
				Name:   arg.Names[0],
				Desc:   arg.Desc,
				EnvVar: arg.Env,
				Value:  arg.Default.(bool),
			})
		case "string":
			arg.storage = res.String(cli.StringArg{
				Name:   arg.Names[0],
				Desc:   arg.Desc,
				EnvVar: arg.Env,
				Value:  arg.Default.(string),
			})
		case "int":
			arg.storage = res.Int(cli.IntArg{
				Name:   arg.Names[0],
				Desc:   arg.Desc,
				EnvVar: arg.Env,
				Value:  arg.Default.(int),
			})
		case "strings":
			arg.storage = res.Strings(cli.StringsArg{
				Name:   arg.Names[0],
				Desc:   arg.Desc,
				EnvVar: arg.Env,
				Value:  arg.Default.([]string),
			})
		case "ints":
			arg.storage = res.Ints(cli.IntsArg{
				Name:   arg.Names[0],
				Desc:   arg.Desc,
				EnvVar: arg.Env,
				Value:  arg.Default.([]int),
			})
		default:
			panic(fmt.Sprintf("Unhandled type %q", arg.Type))

		}
	}

	res.Action = func() {
		fmt.Println("local PARSOPT_OK=1")
		for _, opt := range app.Opts {
			fmt.Printf("local %s=%s\n", opt.Var, bashValue(opt.storage))
		}
		for _, arg := range app.Args {
			fmt.Printf("local %s=%s\n", arg.Var, bashValue(arg.storage))
		}
	}

	return res
}

func runApp(app *cli.Cli, args []string) {
	err := app.Run(args)

	printGuard()

	if err != nil {
		os.Exit(1)
	}
}

func printGuard() {
	fmt.Println(`
if [[ -z "$PARSOPT_OK" ]] ; then
	return
fi
		`)
}

func optNames(d *decl) string {
	res := ""
	for _, n := range d.Names {
		n = strings.TrimPrefix(n, "--")
		n = strings.TrimPrefix(n, "-")
		res = res + " " + n
	}
	return strings.TrimSpace(res)
}

func bashValue(v interface{}) string {
	switch v := v.(type) {
	case *bool:
		if *v {
			return fmt.Sprintf(`"%v"`, *v)
		}
		return ""
	case *int:
		return fmt.Sprintf(`"%d"`, *v)
	case *string:
		return fmt.Sprintf(`"%s"`, *v)
	case *[]int:
		{
			xs := make([]string, len(*v))
			for idx, value := range *v {
				xs[idx] = fmt.Sprintf(`"%d"`, value)
			}
			return fmt.Sprintf(`( %s )`, strings.Join(xs, " "))
		}
	case *[]string:
		{
			xs := make([]string, len(*v))
			for idx, value := range *v {
				xs[idx] = fmt.Sprintf(`"%s"`, value)
			}
			return fmt.Sprintf(`( %s )`, strings.Join(xs, " "))
		}
	default:
		panic("Unsupported case")
	}
}
