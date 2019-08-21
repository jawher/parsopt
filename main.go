package main

import (
	"os"

	"flag"

	"fmt"

	"strings"

	"github.com/fatih/color"
	cli "github.com/jawher/mow.cli"
)

var (
	version string
)

func main() {
	if version == "" {
		version = "tip"
	}

	outerApp := cli.App("parsopt", shortDesc)
	outerApp.Version("v version", version)
	outerApp.Spec = "DIRECTIVES -- [ARGS...]"
	outerApp.LongDesc = longHelp
	outerApp.ErrorHandling = flag.ExitOnError

	var (
		directives = outerApp.StringArg("DIRECTIVES", "", "The accepted options and arguments. Run parsopt --help for the syntax reference")
		args       = outerApp.StringsArg("ARGS", nil, "The args to be passed to the generated CLI app")
	)

	outerApp.Action = func() {
		appDecl, err := parse(*directives)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			printGuard()
			cli.Exit(1)
		}

		app := genApp(appDecl)

		appArgs := make([]string, len(*args)+1)
		appArgs[0] = appDecl.Name
		for i, a := range *args {
			appArgs[i+1] = a
		}
		runApp(app, appArgs)
	}

	if err := outerApp.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

var (
	title  = color.New(color.Bold).SprintfFunc()
	title2 = color.New(color.FgBlue, color.Bold).SprintfFunc()
	title3 = color.New(color.Underline).SprintfFunc()
	code   = color.New(color.FgGreen).SprintfFunc()
	output = color.New(color.FgHiBlack).SprintfFunc()

	shortDesc = `The parsopt utility is used to quickly and easily accept and validate options and arguments in shell procedures.
It's much more powerful compared to getopt[s] as it handles short and long options, arguments, default values, automatic help generation. etc.`
	longHelp = shortDesc + "\n" +

		title("How to use parsopt") + `

parsopt should be called from a shell function.
Place this code snippet in your .bashrc, .zshrc or whatever file is loaded by your shell:
` + code(`
awesome() {
  eval "$(parsopt '
    app AwesomeAppName awesome description

    spec [-r] SRC... DST

    opt recursive  -r --recursive  :bool        Copy the entire src subtree

    arg src                        :strings     The source file to copy
    arg dst                                     The destination file or directory
  ' "$@")"

  # Make use of the ${recursive}, ${src} and ${dst} variables
}`) + `

Make sure to open a new terminal window or tab (or to reload).

If you run:

` + code("$ awesome") + `

You'll get the following output:
` + output(`
Error: incorrect usage

Usage: AwesomeAppName [-r] SRC... DST

awesome description

Arguments:
  SRC=[""]     The source file to copy
  DST=""       The destination file or directory

Options:
  -r, --recursive=false   Copy the entire src subtree

`) + `

(Same with "awesome -h" or "fixer --help").

You just witnessed parsopt options/arguments validation and help message generation in action.


If you call it again with a valid syntax:

` + code("$ awesome file.png /tmp/dir") + `

The variables ${recursive}, ${src} and ${dst} will be populated with "", "file.png" and "/tmp/dir" respectively.
You can then use these variables to perform the desired operation.

` + title("Spec reference") + `

The first argument to parsopt is a multiline string, with one directive per-lien (blank lines are skipped).

A directive has the following syntax:

` + code("DIRECTIVE_NAME VALUE VALUES*") + `

parsopt supports the following directives:

` +
		dirHelp(
			"app",
			"Optional. Let's you set the name of the utility and the description to be shown in the usage message.",
			"app APP_NAME DESC?",
			nil,
			[]string{
				"app awesome",
				"app awesome and here is the description",
			},
		) +
		dirHelp(
			"opt",
			"Optional. Let's you declare an option/flag.",
			"opt VAR_NAME ('-'SHORT_NAME | '--'LONG_NAME)+ (':'TYPE) ('='DEFAULT_VALUE) ('$'ENV_VAR) DESC?",
			[]string{
				"VAR_NAME: the shell variable that will be populated from the option",
				":TYPE: one of \":string\", \":int\", \":bool\", \":ints\", \":strings\" Default is \":string\"",
				"=DEFAULT_VALUE: the default value if the argument is not set, e.g. \"=/tmp/out\", \"=true\"",
				"'$'ENV_VAR: set the value from the specified env var if the option is not set, e.g. \"$HOME\"",
			},
			[]string{"opt debug -d --debug :bool $DEBUG_MODE =false Enable debug log"},
		) +
		dirHelp(
			"arg",
			"Optional. Let's you declare an argument.",
			"arg VAR_NAME (':'TYPE) ('='DEFAULT_VALUE) ('$'ENV_VAR) DESC?",
			[]string{
				"VAR_NAME: the shell variable that will be populated from the argument",
				":TYPE: one of \":string\", \":int\", \":bool\", \":ints\", \":strings\" Default is \":string\"",
				"=DEFAULT_VALUE: the default value if the argument is not set, e.g. \"=/tmp/out\", \"=true\"",
				"'$'ENV_VAR: set the value from the specified env var if the argument is not set, e.g. \"$HOME\"",
			},
			[]string{
				"arg file_name     $DEBUG_MODE      Enable debug log",
				"arg files_to_copy :strings         The files to copy",
			},
		) +
		dirHelp(
			"spec",
			"Optional. Let's you override the default spec string for the app generated by mow.cli (https://github.com/jawher/mow.cli#spec).\n\n"+
				"You can refer to the:\n"+
				"* options (declared using the opt directive) using their short or long names (e.g. -r, --verbose)\n"+
				"* arguments (declared with the arg directive) by uppercasing their shell variable name",
			"spec SPEC",
			nil,
			[]string{
				"spec [ -f | -d ] [-RHL] SRC DST...",
			},
		)
)

func dirHelp(name, desc, syntax string, where []string, examples []string) string {
	res := title2(name) + "\n\n" + desc + "\n\n" +
		title3("Syntax") + "\n\n" + code(syntax) + "\n\n"

	if len(where) > 0 {
		res += "Where:\n"
		for _, w := range where {
			res += "* " + w + "\n"
		}
		res += "\n\n"
	}
	res += title3("Examples") + "\n\n" + code(strings.Join(examples, "\n\n")) + "\n\n"
	return res
}
