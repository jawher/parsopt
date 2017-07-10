#!./test/libs/bats/bin/bats

load 'libs/bats-support/load'
load 'libs/bats-assert/load'
load 'checks'

@test "Should show help with no args" {
    run ./parsopt

    assert_failure 2
    assert_line "Error: incorrect usage"
}

@test "Should accept empty spec" {
    run ./parsopt ''

    assert_success
    assert_parsopt
}

@test "Should handle help" {
    run ./parsopt '' -h

    assert_success
    refute_parsopt
}


@test "Should handle app directive" {
    run ./parsopt 'app Hello desc two words' -h

    assert_success
    assert_output --partial "Usage: Hello"
    assert_output --partial "desc two words"
    refute_parsopt
}
