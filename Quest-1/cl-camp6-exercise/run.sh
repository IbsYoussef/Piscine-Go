#!/usr/bin/env bash

echo "Running your solution in test-data directory:"
echo ""
result=$(cd test-data && bash ../student/countfiles.sh)
echo "Count: $result"