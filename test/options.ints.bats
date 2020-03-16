#!./test/libs/bats/bin/bats

load 'libs/bats-support/load'
load 'libs/bats-assert/load'
load 'checks'

@test "ints opt" {
    run ./parsopt 'opt xxx -x --xxx :ints description'

    assert_success
    assert_var xxx '(  )'
    assert_parsopt
}

@test "ints opt default value (invalid)" {
    run ./parsopt 'opt xxx -x --xxx :ints =default description'

    assert_parse_error 0 'Invalid default value "default"'
}

@test "ints opt default value (valid)" {
    run ./parsopt 'opt xxx -x --xxx :ints =7,42 description'

    assert_success
    assert_var xxx '( "7" "42" )'
    assert_parsopt
}

@test "ints opt env value (valid)" {
    export XXX_DEFAULT="9,666"
    run ./parsopt 'opt xxx -x --xxx :ints =7,42 $XXX_DEFAULT description'

    assert_success
    assert_var xxx '( "9" "666" )'
    assert_parsopt
}

@test "ints opt user value (short) (invalid) (#1)" {
    export XXX_DEFAULT="9,666"
    run ./parsopt 'opt xxx -x --xxx :ints =7,42 $XXX_DEFAULT description' -x "FROM USER"

    assert_failure
    refute_parsopt
}

@test "ints opt user value (short) (valid) (#1)" {
    export XXX_DEFAULT="9,666"
    run ./parsopt '
        opt xxx -x --xxx :ints =7,42 $XXX_DEFAULT description
        spec -x...' -x 3 -x 7

    assert_success
    assert_var xxx '( "3" "7" )'
    assert_parsopt
}

@test "ints opt user value (short) (valid) (#2)" {
    export XXX_DEFAULT="9,666"
    run ./parsopt '
        opt xxx -x --xxx :ints =7,42 $XXX_DEFAULT description
        spec -x...' -x=7 -x=11

    assert_success
    assert_var xxx '( "7" "11" )'
    assert_parsopt
}

@test "ints opt user value (short) (valid) (#3)" {
    export XXX_DEFAULT="9,666"
    run ./parsopt '
        opt xxx -x --xxx :ints =7,42 $XXX_DEFAULT description
        spec -x...' -x5 -x9

    assert_success
    assert_var xxx '( "5" "9" )'
    assert_parsopt
}


@test "ints opt user value (long) (invalid) (#1)" {
    export XXX_DEFAULT="9,666"
    run ./parsopt '
        opt xxx -x --xxx :ints =7,42 $XXX_DEFAULT description
        spec -x...' --xxx "FROM USER"

    assert_failure
    refute_parsopt
}

@test "ints opt user value (long) (valid) (#1)" {
    export XXX_DEFAULT="9,666"
    run ./parsopt '
        opt xxx -x --xxx :ints =7,42 $XXX_DEFAULT description
        spec -x...' --xxx 5 --xxx 9

    assert_success
    assert_var xxx '( "5" "9" )'
    assert_parsopt
}

@test "ints opt user value (long) (invalid) (#2)" {
    export XXX_DEFAULT="9,666"
    run ./parsopt '
        opt xxx -x --xxx :ints =7,42 $XXX_DEFAULT description
        spec -x...' --xxx="FROM USER"

    assert_failure
    refute_parsopt
}

@test "ints opt user value (long) (valid) (#2)" {
    export XXX_DEFAULT="9,666"
    run ./parsopt '
        opt xxx -x --xxx :ints =7,42 $XXX_DEFAULT description
        spec -x...' --xxx=5 --xxx=9

    assert_success
    assert_var xxx '( "5" "9" )'
    assert_parsopt
}
