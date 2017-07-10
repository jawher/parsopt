#!./test/libs/bats/bin/bats

load 'libs/bats-support/load'
load 'libs/bats-assert/load'
load 'checks'

@test "bool opt" {
    run ./parsopt 'opt xxx -x --xxx :bool description'

    assert_success
    assert_var xxx
    assert_parsopt
}

@test "bool opt default value (invalid)" {
    run ./parsopt 'opt xxx -x --xxx :bool =default description'

    assert_parse_error 0 'Invalid default value "default"'
}

@test "bool opt default value (valid) (true)" {
    run ./parsopt 'opt xxx -x --xxx :bool =true description'

    assert_success
    assert_var xxx '"true"'
    assert_parsopt
}

@test "bool opt default value (valid) (false)" {
    run ./parsopt 'opt xxx -x --xxx :bool =false description'

    assert_success
    assert_var xxx
    assert_parsopt
}

@test "bool opt env value (invalid)" {
    export XXX_DEFAULT="FROM ENV"
    run ./parsopt 'opt xxx -x --xxx :bool $XXX_DEFAULT description'

    assert_success
    assert_var xxx
    assert_parsopt
}

@test "bool opt env value (valid)" {
    export XXX_DEFAULT="true"
    run ./parsopt 'opt xxx -x --xxx :bool $XXX_DEFAULT description'

    assert_success
    assert_var xxx '"true"'
    assert_parsopt
}

@test "bool opt user value (short)" {
    export XXX_DEFAULT="false"
    run ./parsopt 'opt xxx -x --xxx :bool $XXX_DEFAULT description' -x

    assert_success
    assert_var xxx '"true"'
    assert_parsopt
}

@test "bool opt user value (long)" {
    export XXX_DEFAULT="false"
    run ./parsopt 'opt xxx -x --xxx :bool $XXX_DEFAULT description' --xxx

    assert_success
    assert_var xxx '"true"'
    assert_parsopt
}
