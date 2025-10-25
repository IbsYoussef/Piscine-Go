#!/usr/bin/env bash

echo "Running your solution in test-data directory:"
echo ""
result=$(cd test-data && bash ../student/lookagain.sh)

if [ -z "$result" ]; then
    echo "(No .sh files found)"
else
    echo "$result"
fi