#!/usr/bin/env bash
set -euo pipefail

echo "Running printprogramname..."
echo ""

# Navigate to student folder and build
cd student/printprogramname
go build

# Run the program
./printprogramname

# Clean up
rm -f printprogramname
cd ../..