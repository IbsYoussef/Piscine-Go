#!/usr/bin/env bash
set -euo pipefail

echo "Running point..."
echo ""

# Navigate to student folder and build
cd student/point
go build

./point

# Clean up
rm -f point
cd ../..