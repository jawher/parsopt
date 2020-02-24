#!./test/libs/bats/bin/bats

load 'libs/bats-support/load'
load 'libs/bats-assert/load'
load 'checks'

@test "strings opt" {
    run ./parsopt 'opt xxx -x --xxx :strings description'

    assert_success
    assert_var xxx '(  )'
    assert_parsopt
}

@test "strings opt default value" {
    run ./parsopt 'opt xxx -x --xxx :strings =default1,default2 description'

    assert_success
    assert_var xxx '( "default1" "default2" )'
    assert_parsopt
}

@test "strings opt env value" {
    export XXX_DEFAULT="ENV VAL 1, ENV VAL 2"
    run ./parsopt '
        opt xxx -x --xxx :strings =default $XXX_DEFAULT description
        spec -x...'

    assert_success
    assert_var xxx '( "ENV VAL 1" "ENV VAL 2" )'
    assert_parsopt
}

@test "strings opt user value (short) (#1)" {
    export XXX_DEFAULT="ENV VAL 1, ENV VAL 2"
    run ./parsopt '
        opt xxx -x --xxx :strings =default $XXX_DEFAULT description
        spec -x...' -x "FROM USER 1" -x "FROM USER 2"

    assert_success
    assert_var xxx '( "FROM USER 1" "FROM USER 2" )'
    assert_parsopt
}

@test "strings opt user value (short) (#2)" {
    export XXX_DEFAULT="ENV VAL 1, ENV VAL 2"
    run ./parsopt '
        opt xxx -x --xxx :strings =default $XXX_DEFAULT description
        spec -x...' -x="FROM USER 1" -x="FROM USER 2"

    assert_success
    assert_var xxx '( "FROM USER 1" "FROM USER 2" )'
    assert_parsopt
}

@test "strings opt user value (short) (#3)" {
    export XXX_DEFAULT="ENV VAL 1, ENV VAL 2"
    run ./parsopt '
        opt xxx -x --xxx :strings =default $XXX_DEFAULT description
        spec -x...' -x"FROM USER 1" -x"FROM USER 2"

    assert_success
    assert_var xxx '( "FROM USER 1" "FROM USER 2" )'
    assert_parsopt
}

@test "strings opt user value (long) (#1)" {
    export XXX_DEFAULT="ENV VAL 1, ENV VAL 2"
    run ./parsopt '
        opt xxx -x --xxx :strings =default $XXX_DEFAULT description
        spec -x...' --xxx "FROM USER 1" --xxx "FROM USER 2"

    assert_success
    assert_var xxx '( "FROM USER 1" "FROM USER 2" )'
    assert_parsopt
}

@test "strings opt user value (long) (#2)" {
    export XXX_DEFAULT="ENV VAL 1, ENV VAL 2"
    run ./parsopt '
        opt xxx -x --xxx :strings =default $XXX_DEFAULT description
        spec -x...' --xxx="FROM USER 1" --xxx="FROM USER 2"

    assert_success
    assert_var xxx '( "FROM USER 1" "FROM USER 2" )'
    assert_parsopt
}
