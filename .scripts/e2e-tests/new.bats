#!/usr/bin/env bats

# e2e tests for the `pkb new` command

out_dir=/tmp/pkb-output

setup_file() {
	echo "### suite setup ###"
	mkdir "$out_dir"
	cp -r "$BATS_TEST_DIRNAME/templates" "$out_dir/templates"
}

teardown_file() {
	echo "### suite teardown ###"
	rm -rf "$out_dir"
}

setup() {
	echo "### test setup ###"
	bats_load_library bats-support
	bats_load_library bats-assert
}

teardown() {
	echo "### test teardown ###"
}

@test "pkb new: file is created with --no-edit" {
	eval "run $BATS_TEST_DIRNAME/new-no-edit.exp $BATS_TEST_DIRNAME/pkb.json"
	assert_success

	expected_path="$out_dir/simple/foo.md"
	assert_line --index 2 --partial "file created: $expected_path"

	run cat "$expected_path"
	assert_success
	assert_line --index 0 "# foo document"
	assert_line --index 1 "simple generation for test"
}
