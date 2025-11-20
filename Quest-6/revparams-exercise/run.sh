#!/usr/bin/env bash
set -euo pipefail

echo "Running revparams..."
echo ""

# Navigate to student folder and build
cd student/revparams
go build

# Run the program with test arguments
./revparams choumi is the best cat

# Clean up
rm -f revparams
cd ../..