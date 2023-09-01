#!/bin/bash

# Script for running instantiating tests.
# Must be run from the project's root.

set -e

cd tests/instantiating/

echo -e "\nRunning instantiating tests\n"

for dir in $(find . -maxdepth 1 -mindepth 1 -type d | sort);
do
	echo "    $dir"
	cd $dir
	 ../../../fbdl bus.fbd > /dev/null 2>stderr || true
	diff --color stderr.golden stderr
	rm stderr
	cd ..
done

echo -e "\nAll \e[1;32mPASSED\e[0m!"
