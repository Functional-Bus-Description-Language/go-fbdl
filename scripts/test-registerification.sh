#!/bin/bash

# Script for running registerification tests.
# Must be run from the project's root.

set -e

cd tests/registerification/

echo -e "Running registerification tests\n"

for dir in */ ;
do
	echo "  $dir"
	cd $dir
	../../../fbdl -z -r bus.fbd
	diff golden.json reg.json
	rm reg.json
	cd ../..
done

echo -e "\nAll \e[1;32mPASSED\e[0m!"
