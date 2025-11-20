#!/usr/bin/env bash
set -euo pipefail

echo "Running sortparams..."
echo ""

# Navigate to student folder and build
cd student/sortparams
go build

# Run the program with test arguments
./sortparams 1 a 2 A 3 b 4 C

# Clean up
rm -f sortparams
cd ../..