#!./test/libs/bats/bin/bats

load 'libs/bats-support/load'
load 'libs/bats-assert/load'
load 'checks'

export TEST_SPEC_1='
   app test1

   spec ([-b|-s] | --ii...) ARG1 ARG2...

   opt bool1 -b :bool
   opt str1  -s :string =default1
   opt ints1  --ii :ints

   arg arg1 :int
   arg arg2 :strings =default1,default2
'

@test "IT :: Test 1" {
    eval "$(./parsopt "$TEST_SPEC_1" -b 3)"

    assert_equal "${bool1}" "true"
    assert_equal "${str1}" "default1"
    assert_equal "${ints1}" "(  )"
    assert_equal "${arg1}" "3"
    assert_equal "${arg2[*]}" 'default1 default2'

    assert_equal "${PARSOPT_OK}" "1"

}

@test "IT :: Test 2" {
    eval "$(./parsopt "$TEST_SPEC_1" -s aloha --ii=1 --ii=2 3)"

    assert_equal "${bool1}" ""
    assert_equal "${str1}" "aloha"
    assert_equal "${ints1}" '( "1" "2" )'
    assert_equal "${arg1}" "3"
    assert_equal "${arg2[*]}" 'default1 default2'

    assert_equal "${PARSOPT_OK}" "1"
}

@test "IT :: Test 3" {
    eval "$(./parsopt "$TEST_SPEC_1" -s aloha 42 a b c)"

    assert_equal "${bool1}" ""
    assert_equal "${str1}" "aloha"
    assert_equal "${arg1}" "42"
    assert_equal "${arg2[*]}" 'a b c'


    assert_equal "${PARSOPT_OK}" "1"
}


@test "IT :: Incorrect usage" {
    eval "$(./parsopt "$TEST_SPEC_1" -b -s aloha 42)"

    fail "should have returned"
}


@test "IT :: Incorrect usage (2)" {
    eval "$(./parsopt "$TEST_SPEC_1" --ii invalid arg)"

    fail "should have returned"
}
