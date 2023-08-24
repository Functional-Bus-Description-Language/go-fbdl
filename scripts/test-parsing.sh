#!/bin/bash

# Script for running parsing tests.
# Must be run from the project's root.

set -e

cd tests/parsing/

echo -e "\nRunning parsing tests\n"

for dir in $(find . -maxdepth 1 -mindepth 1 -type d);
do
	echo "    $dir"
	cd $dir
	../../../fbdl bus.fbd > /dev/null 2>stderr || true
	diff --color stderr.golden stderr
	rm stderr
	cd ..
done

echo -e "\nAll \e[1;32mPASSED\e[0m!"
