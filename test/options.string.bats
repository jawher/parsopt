#!./test/libs/bats/bin/bats

load 'libs/bats-support/load'
load 'libs/bats-assert/load'
load 'checks'

@test "string opt" {
    run ./parsopt 'opt xxx -x --xxx description'

    assert_success
    assert_var xxx '""'
    assert_parsopt
}

@test "string opt default value" {
    run ./parsopt 'opt xxx -x --xxx =default description'

    assert_success
    assert_var xxx '"default"'
    assert_parsopt
}

@test "string opt env value" {
    export XXX_DEFAULT="FROM ENV"
    run ./parsopt 'opt xxx -x --xxx =default $XXX_DEFAULT description'

    assert_success
    assert_var xxx '"FROM ENV"'
    assert_parsopt
}

@test "string opt user value (short) (#1)" {
    export XXX_DEFAULT="FROM ENV"
    run ./parsopt 'opt xxx -x --xxx =default $XXX_DEFAULT description' -x "FROM USER"

    assert_success
    assert_var xxx '"FROM USER"'
    assert_parsopt
}

@test "string opt user value (short) (#2)" {
    export XXX_DEFAULT="FROM ENV"
    run ./parsopt 'opt xxx -x --xxx =default $XXX_DEFAULT description' -x="FROM USER"

    assert_success
    assert_var xxx '"FROM USER"'
    assert_parsopt
}

@test "string opt user value (short) (#3)" {
    export XXX_DEFAULT="FROM ENV"
    run ./parsopt 'opt xxx -x --xxx =default $XXX_DEFAULT description' -x"FROM USER"

    assert_success
    assert_var xxx '"FROM USER"'
    assert_parsopt
}

@test "string opt user value (long) (#1)" {
    export XXX_DEFAULT="FROM ENV"
    run ./parsopt 'opt xxx -x --xxx =default $XXX_DEFAULT description' --xxx "FROM USER"

    assert_success
    assert_var xxx '"FROM USER"'
    assert_parsopt
}

@test "string opt user value (long) (#2)" {
    export XXX_DEFAULT="FROM ENV"
    run ./parsopt 'opt xxx -x --xxx =default $XXX_DEFAULT description' --xxx="FROM USER"

    assert_success
    assert_var xxx '"FROM USER"'
    assert_parsopt
}
