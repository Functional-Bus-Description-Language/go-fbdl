#!/bin/bash

update=false

help_msg="Script for managing registerification tests.
Must be run from the project's root.

Usage:
  scripts/reg-tests.sh <command>

Commands:
  help    Display help message.
  run     Run tests.
  update  Run tests discarding errors and update golden.json files using reg.json files.

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

cd tests/registerification/

echo -e "Running registerification tests\n"

for dir in $(find . -maxdepth 3 -mindepth 3 -type d);
do
	echo "  $dir"
	cd "$dir"
	../../../../../fbdl -r bus.fbd
	diff --color golden.json reg.json
	if $update; then
		cp reg.json golden.json
	fi
	rm reg.json
	cd ../../..
done

if $update; then
	echo -e "\ngolden.json files updated\n"
else
	echo -e "\nAll \e[1;32mPASSED\e[0m!"
fi
