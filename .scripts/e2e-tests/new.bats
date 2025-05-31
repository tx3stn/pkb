#!/usr/bin/env bats

# e2e tests for the `pkb new` command

setup() {
	echo "### test setup ###"
	bats_load_library bats-support
	bats_load_library bats-assert
}

@test "pkb new: help" {
	run pkb --help
	assert_success
}
