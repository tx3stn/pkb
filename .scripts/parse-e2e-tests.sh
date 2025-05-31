#!/bin/sh

# Parse the e2e test output and create a simple markdown table from the results
# to display as a summary in the GITHUB_STEP_SUMMARY output.

# reads from /dev/stdin so you can pipe tests to the script.
summary=$(cat -)

echo '## 🧪 end to end test results'
echo ''
echo '| status | command | test name |'
echo '| --- | --- | --- |'
echo "$summary" | while IFS= read -r line; do
	if [ "$line" != "${line#1..}" ]; then
		continue
	fi

	status=$(echo "$line" | cut -d' ' -f1)

	if [ "$status" = "ok" ]; then
		result_icon='✓'
	else
		result_icon='✕'
	fi

	cmd_name=$(echo "$line" | cut -d' ' -f4)
	test_name=$(echo "$line" | cut -d':' -f2 | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//')

	echo "| $result_icon | pkb $cmd_name | $test_name |"
done
