#!/bin/bash

set -e

if [ -z "$1" ]; then
  echo "Usage: $0 <iterations>"
  exit 1
fi

for ((i = 1; i <= $1; i++)); do
  echo "Running iteration $i"
  result=$(claude --allowedTools "Bash,Read,Write" --permission-mode acceptEdits -p "@PRD.md @progress.txt \
    1. Read the PRD and progress file. \
    2. Find the next incomplete task and implement it. \
    DO YOUR BEST TO KEEP EACH COMMIT TO LESS THAN 200 LINES OF CODE. \
    3. Commit your changes. \
    4. Update progress.txt with what you did. \
    WHEN WRITING TESTS IN GO, ALWAYS USE THE testify LIBRARIES. \
    ONLY DO ONE TASK AT A TIME. \
    If the PRD is complete, output <promise>COMPLETE</promise>.")

  echo "$result"

  if [[ "$result" == *"<promise>COMPLETE</promise>"* ]]; then
    echo "PRD complete after $i iterations."
    exit 0
  fi
done
