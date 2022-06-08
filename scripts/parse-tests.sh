#! /bin/sh

# Pase the unit test output and create a simple markdown table from the results
# to display the code coverage summary in the GITHUB_STEP_SUMMARY output.

# reads from /dev/stdin so you can pipe tests to the script.
summary=$(cat -)

echo '## 📈 Coverage Summary'
echo ''
echo '| package | coverage |'
echo '| --- | --- |'
echo "$summary" | while IFS= read -r line
do
  pkg=$(echo "$line" | cut -f2)
  coverage=$(echo "$line" | cut -f4 | cut -d' ' -f 2)
  echo "| $pkg | $coverage |"
done
