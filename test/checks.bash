
assert_parsopt() {
	assert_line "local PARSOPT_OK=1"
	assert_parsopt_guard
}

refute_parsopt() {
	refute_line "local PARSOPT_OK=1"
	assert_parsopt_guard
}

assert_parsopt_guard() {
	assert_output --regexp '.*if \[\[ -z "\$PARSOPT_OK" ]] ; then[[:space:]]*return[[:space:]]*fi[[:space:]]*$'
}

assert_var() {
    name=$1
    val=$2
	assert_line "local ${name}=${val}"
}

assert_parse_error() {
    row=$1
    msg=$2
    assert_failure
	assert_line "Parse error at row ${row}: ${msg}"
}
