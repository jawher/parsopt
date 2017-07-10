#!./test/libs/bats/bin/bats

load 'libs/bats-support/load'
load 'libs/bats-assert/load'
load 'checks'

@test "int opt" {
    run ./parsopt 'opt xxx -x --xxx :int description'

    assert_success
    assert_var xxx '"0"'
    assert_parsopt
}

@test "int opt default value (invalid)" {
    run ./parsopt 'opt xxx -x --xxx :int =default description'

    assert_parse_error 0 'Invalid default value "default"'
}

@test "int opt default value (valid)" {
    run ./parsopt 'opt xxx -x --xxx :int =42 description'

    assert_success
    assert_var xxx '"42"'
    assert_parsopt
}

@test "int opt env value (invalid)" {
    export XXX_DEFAULT="FROM ENV"
    run ./parsopt 'opt xxx -x --xxx :int =42 $XXX_DEFAULT description'

    assert_success
    assert_var xxx '"42"'
    assert_parsopt
}

@test "int opt env value (valid)" {
    export XXX_DEFAULT="666"
    run ./parsopt 'opt xxx -x --xxx :int =42 $XXX_DEFAULT description'

    assert_success
    assert_var xxx '"666"'
    assert_parsopt
}

@test "int opt user value (short) (invalid) (#1)" {
    export XXX_DEFAULT="666"
    run ./parsopt 'opt xxx -x --xxx :int =42 $XXX_DEFAULT description' -x "FROM USER"

    assert_failure
    refute_parsopt
}

@test "int opt user value (short) (valid) (#1)" {
    export XXX_DEFAULT="666"
    run ./parsopt 'opt xxx -x --xxx :int =42 $XXX_DEFAULT description' -x 7

    assert_success
    assert_var xxx '"7"'
    assert_parsopt
}

@test "int opt user value (short) (valid) (#2)" {
    export XXX_DEFAULT="666"
    run ./parsopt 'opt xxx -x --xxx :int =42 $XXX_DEFAULT description' -x=7

    assert_success
    assert_var xxx '"7"'
    assert_parsopt
}

@test "int opt user value (short) (valid) (#3)" {
    export XXX_DEFAULT="666"
    run ./parsopt 'opt xxx -x --xxx :int =42 $XXX_DEFAULT description' -x7

    assert_success
    assert_var xxx '"7"'
    assert_parsopt
}


@test "int opt user value (long) (invalid) (#1)" {
    export XXX_DEFAULT="666"
    run ./parsopt 'opt xxx -x --xxx :int =42 $XXX_DEFAULT description' --xxx "FROM USER"

    assert_failure
    refute_parsopt
}

@test "int opt user value (long) (valid) (#1)" {
    export XXX_DEFAULT="666"
    run ./parsopt 'opt xxx -x --xxx :int =42 $XXX_DEFAULT description' --xxx 7

    assert_success
    assert_var xxx '"7"'
    assert_parsopt
}

@test "int opt user value (long) (invalid) (#2)" {
    export XXX_DEFAULT="666"
    run ./parsopt 'opt xxx -x --xxx :int =42 $XXX_DEFAULT description' --xxx="FROM USER"

    assert_failure
    refute_parsopt
}

@test "int opt user value (long) (valid) (#2)" {
    export XXX_DEFAULT="666"
    run ./parsopt 'opt xxx -x --xxx :int =42 $XXX_DEFAULT description' --xxx=7

    assert_success
    assert_var xxx '"7"'
    assert_parsopt
}
