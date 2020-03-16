# parsopt : getopt[s] on steroids

![CI](https://github.com/jawher/parsopt/workflows/CI/badge.svg)

The parsopt utility is used to quickly and easily accept and validate options and arguments in shell procedures.
It's much more powerful and intuitive compared to getopt[s] as it handles short and long options, arguments, automatic help generation. etc.

## Quickstart

### Install parsopt

Download your platform binary from this page: https://github.com/jawher/parsopt/releases/tag/v0.1


### Use it

To showcase what parsopt does, we'll build a command line utility to convert between currencies.
[fixer.io](http://fixer.io) will be used for the conversion part.
Also, this will use [HTTPie](https://github.com/jakubroztocil/httpie) (to perform the API calls) and [jq](https://stedolan.github.io/jq/) to parse the json responses.


parsopt should be called from a shell function.
Place this code snippet in your `.bashrc`, `.zshrc` or whatever file is loaded by your shell:


```bash
fixer() {
    eval "$(parsopt '
            app fixer Fixer API from the CLI

            spec [OPTIONS] FROM [TO]

            opt https  -s --use-https  :bool                      Use HTTPS
            opt date   -d --date       =latest                    Use this date''s historical rate

            arg from                                              The currency to convert from
            arg to                     =EUR      $FIXER_TO        Return rates based on this currency
        ' "$@")"



    protocol="http"
    [[ -n "${https}" ]] && protocol="https"

    http get "${protocol}://api.fixer.io/${date:-latest}" "base==${to}" "symbols==${from}" | jq ".rates | .${from}"
}

```

Make sure to open a new terminal window or tab (or to reload).

If you run:

```
fixer
```

You'll get the following output:

```
$ fixer
Error: incorrect usage

Usage: fixer [OPTIONS] FROM [TO]

Fixer API from the CLI

Arguments:
  FROM=""      The currency to convert from
  TO="EUR"     Return rates based on this currency ($FIXER_TO)

Options:
  -s, --use-https=false   Use HTTPS
  -d, --date="latest"     Use this dates historical rate
```

(Same with `fixer -h` or `fixer --help`).

You just witnessed parsopt options/arguments validation and help message generation in action.


If you call it again with a valid syntax:

```
fixer USD EUR
```

Will proceed to call the fixer.io API and convert 1 USD to EUR.

## Reference

### Usage
parsopt must be called from within a shell function, as follows:


```bash
function something() {
    eval "$(parsopt 'SPEC_STRING' "$@")"

    # Use the variables defined in SPEC_STRING
}
```

### Spec

The `SPEC_STRING` is a multiline string, where every line is a directive.

A directive has the following syntax:


```
DIRECTIVE_NAME VALUE VALUES*
```

parsopt supports the following directives:

#### app

Optional. Let's you set the name of the utility and the description to be shown in the usage message.

**Syntax**

```
app APP_NAME DESC?
```

**Examples***

```
app fixer
```

```
app fixer a CLI to convert between currencies
```


#### opt

Optional. Let's you declare an option/flag.

**Syntax**

```
opt VAR_NAME ('-'SHORT_NAME | '--'LONG_NAME)+ (':'TYPE) ('='DEFAULT_VALUE) ('$'ENV_VAR) DESC?
```

Where:
* `VAR_NAME`: the shell variable that will be populated from the option
* `SHORT_NAME` and `LONG_NAME`: the option names (1 letter names for short names, more for long names), e.g. `-d`, `--debug`, etc.
* `':'TYPE`: one of `:string`, `:int`, `:bool`, `:ints`, `:strings` Default is `:string`
* `'='DEFAULT_VALUE`: the default value if the option is not set, e.g. `=/tmp/out`, `=true`
* `'$'ENV_VAR`: set the value from the specified env var if the option is not set, e.g. `$HOME`

**Examples***

```
opt debug -d --debug :bool $DEBUG_MODE =false Enable debug log
```

#### arg

Optional. Let's you declare an argument.

```
arg VAR_NAME (':'TYPE) ('='DEFAULT_VALUE) ('$'ENV_VAR) DESC?
```

Where:
* `VAR_NAME`: the shell variable that will be populated from the argument
* `':'TYPE`: one of `:string`, `:int`, `:bool`, `:ints`, `:strings` Default is `:string`
* `'='DEFAULT_VALUE`: the default value if the argument is not set, e.g. `=/tmp/out`, `=true`
* `'$'ENV_VAR`: set the value from the specified env var if the argument is not set, e.g. `$HOME`

**Examples***

```
arg file_name $DEBUG_MODE Enable debug log
```

#### spec

Optional. Let's you override the default spec string for the app generated by [mow.cli](https://github.com/jawher/mow.cli#spec).

You can refer to the:

* options (declared using the opt directive) using their short or long names (e.g. -r, --verbose)
* arguments (declared with the arg directive) by uppercasing their shell variable name


```
spec SPEC
```

**Examples***

```
spec [ -f | -d ] [-RHL] SRC DST...
```
