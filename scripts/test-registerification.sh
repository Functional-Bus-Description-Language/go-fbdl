#!/bin/bash

# Script for running registerification tests.
# Must be run from the project's root.

set -e

cd tests/registerification/

echo -e "Running registerification tests\n"

for dir in $(find . -maxdepth 3 -mindepth 3 -type d);
do
	echo "  $dir"
	cd $dir
	../../../../../fbdl -z -r bus.fbd
	diff --color golden.json reg.json
	rm reg.json
	cd ../../..
done

echo -e "\nAll \e[1;32mPASSED\e[0m!"
