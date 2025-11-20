#!/usr/bin/env bash
set -euo pipefail

echo "Running printparams..."
echo ""

# Navigate to student folder and build
cd student/printparams
go build

# Run the program with test arguments
./printparams choumi is the best cat

# Clean up
rm -f printparams
cd ../..