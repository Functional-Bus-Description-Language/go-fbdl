#!/bin/bash

# Script for running parsing tests.
# Must be run from the project's root.

set -e

cd tests/parsing/

echo -e "Running parsing tests\n"

#echo -e "  Testing valid constructs\n"
#for dir in $(find valid/ -maxdepth 1 -mindepth 1 -type d);
#do
#	echo "    $dir"
#	cd $dir
#	python3 ../../../../main.py -p packages bus.fbd > /dev/null 2>&1
#	sed '/\/home\//d' packages > packages.sed
#	sed '/\/home\//d' packages.golden > packages.golden.sed
#	diff packages.golden.sed packages.sed
#	rm packages
#	rm *.sed
#	cd ../..
#done

echo -e "\n  Testing invalid constructs\n"
for dir in $(find invalid/ -maxdepth 1 -mindepth 1 -type d);
do
	echo "    $dir"
	cd $dir
	../../../../fbdl bus.fbd > /dev/null 2>stderr || true
	diff --color stderr.golden stderr
	rm stderr
	cd ../..
done

echo -e "\nAll \e[1;32mPASSED\e[0m!"
