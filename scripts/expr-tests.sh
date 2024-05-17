#!/bin/bash

update=false

help_msg="Script for managing expression evaluation tests.
Must be run from the project's root.

Usage:
  scripts/expr-tests.sh <command>

Commands:
  help    Display help message.
  run     Run tests.
  update  Run tests discarding errors and update stderr.golden files using stderr files.

If no command is provided the run is assumed.
"

while true ; do
	case "$1" in
		help) printf "$help_msg" ; exit 0 ;;
		run) shift ;;
		update) update=true ; shift ;;
		"") shift ; break ;;
		*) echo "invalid argument '$1'" ; exit 1 ;;
	esac
done

if ! $update; then
	set -e
fi

cd tests/expr/

echo -e "\nRunning expression evaluation tests\n"

for dir in $(find . -maxdepth 1 -mindepth 1 -type d | sort);
do
	testname=`basename $dir`
	# Ignore tests starting with '_' character.
	if [ ${testname::1} = "_" ]; then
		continue
	fi

	echo "  $dir"
	cd "$dir"
	../../../fbdl -r bus.fbd > /dev/null 2>stderr || true
	diff --color stderr.golden stderr
	if $update; then
		cp stderr stderr.golden
	fi
	rm stderr
	cd ..
done

if $update; then
	echo -e "\nstderr.golden files updated\n"
else
	echo -e "\nAll \e[1;32mPASSED\e[0m!"
fi
