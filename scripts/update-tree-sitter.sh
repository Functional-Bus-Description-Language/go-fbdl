#!/bin/bash

curl https://raw.githubusercontent.com/Functional-Bus-Description-Language/tree-sitter-fbdl/master/src/parser.c > ./internal/ts/parser.c
curl https://raw.githubusercontent.com/Functional-Bus-Description-Language/tree-sitter-fbdl/master/src/scanner.c > ./internal/ts/scanner.c
curl https://raw.githubusercontent.com/Functional-Bus-Description-Language/tree-sitter-fbdl/master/src/tree_sitter/parser.h > ./internal/ts/tree_sitter/parser.h
