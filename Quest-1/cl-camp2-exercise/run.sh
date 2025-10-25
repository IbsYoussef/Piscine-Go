#!/usr/bin/env bash

echo "Your file content:"
grep -v '^\s*#' student/r | grep -v '^\s*$' | cat -e