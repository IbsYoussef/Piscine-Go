#!/usr/bin/env bash

echo "Running your solution in test-data directory:"
echo ""
result=$(cd test-data && grep -v '^\s*#' ../student/look | grep -v '^\s*$' | bash)

if [ -z "$result" ]; then
    echo "(No matching files found)"
else
    echo "$result"
fi