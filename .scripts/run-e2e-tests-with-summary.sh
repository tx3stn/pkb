#!/bin/bash

# Script to run end to edn tests in CI and gracefully handle the status to display
# all of the info you need to debug and output the results to the
# GITHUB_STEP_SUMMARY environment variable so the test summary is always
# displayed if things pass or fail.

TEST_OUTPUT=$(bats --verbose-run --formatter tap .scripts/e2e-tests/)
EXIT_CODE="$?"

if [ "$EXIT_CODE" = 0 ]; then
	echo ""
	echo "# 🎉 end to end tests passed 🎉"
	echo ""
	SUMMARY=$(echo "$TEST_OUTPUT" | ./.scripts/parse-e2e-tests.sh)
	echo "$SUMMARY"
	exit 0
fi

echo "# 🚩 some end to end tests failed 🚩"
echo "\`\`\`"
echo "$TEST_OUTPUT"
echo "\`\`\`"
exit 1
