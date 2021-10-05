#include <err.h>
#include <errno.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "parser.h"

//#define DEBUG

enum TokenType {
	INDENT,
	DEDENT,
	NEWLINE
};

struct Scanner {
	unsigned current_indent;
	unsigned dedents;
};

void * tree_sitter_fbdl_external_scanner_create() {
	struct Scanner* scanner = malloc(sizeof(struct Scanner));
	scanner->current_indent = 0;
	scanner->dedents = 0;

	if (scanner == NULL)
		errx(EXIT_FAILURE, "Can't allocate memory for external scanner");

	return scanner;
}

void tree_sitter_fbdl_external_scanner_destroy(void *payload) {
	struct Scanner *scanner = payload;

#ifdef DEBUG
	printf("On destroy:\n");
	printf("  Current indent: %d\n", scanner->current_indent);
	printf("  Dedents: %d\n", scanner->dedents);
#endif
	free(payload);
}

unsigned tree_sitter_fbdl_external_scanner_serialize(
	void *payload,
	char *buffer
) {
	memcpy(buffer, payload, sizeof(struct Scanner));

	return sizeof(struct Scanner);
}

void tree_sitter_fbdl_external_scanner_deserialize(
	void *payload,
	const char *buffer,
	unsigned length
) {
	if (length > 0) {
		memcpy(payload, buffer, sizeof(struct Scanner));
	}

	return;
}

bool tree_sitter_fbdl_external_scanner_scan(
	void *payload,
	TSLexer *lexer,
	const bool *valid_symbols
) {
#ifdef DEBUG
	static unsigned debug_cnt = 0;
#endif

	struct Scanner *scanner = payload;

	if (valid_symbols[INDENT] || valid_symbols[DEDENT]) {
#ifdef DEBUG
		debug_cnt++;
		// Debug code
		if (valid_symbols[INDENT] && valid_symbols[DEDENT]) {
			printf("%d ID: lookahead %d current indent %d\n", debug_cnt, lexer->lookahead, scanner->current_indent);
		} else if (valid_symbols[DEDENT]) {
			printf("%d  D: lookahead %d current indent %d\n", debug_cnt, lexer->lookahead, scanner->current_indent);
		} else if (valid_symbols[INDENT]) {;
			printf("%d I : lookahead %d currnt indent %d\n", debug_cnt, lexer->lookahead, scanner->current_indent);
		}
#endif

		if (scanner->dedents > 0) {
			if (valid_symbols[DEDENT]) {
#ifdef DEBUG
				printf("scanner.dedents = %d, going to dedent", scanner->dedents);
#endif
				goto dedent;
			}
		}

		// Consume line feeds before indent.
		// This is needed to handle poorly aligned comments.
		if (valid_symbols[INDENT]) {
			while (lexer->lookahead == '\n') {
				lexer->advance(lexer, false);
			}
		}

		// Handle scenario where previous dedent consumed the indentation.
		if (lexer->get_column(lexer) != 0) {
			return false;
		}
		unsigned indent = 0;
		while (lexer->lookahead == '\t') {
			indent++;
			lexer->advance(lexer, true);
		}

		if (valid_symbols[DEDENT] && (indent < scanner->current_indent)) {
			scanner->dedents = scanner->current_indent - indent;
dedent:
			scanner->current_indent--;
			scanner->dedents--;
			lexer->result_symbol = DEDENT;
#ifdef DEBUG
			printf("   D: current indent %d\n", scanner->current_indent);
#endif
			return true;
		}

		if (valid_symbols[INDENT] && (indent > scanner->current_indent)) {
			if (indent != scanner->current_indent + 1) {
				errx(EXIT_FAILURE, "Multi indent detected, indent %d, current indent %d", indent, scanner->current_indent);
			}

			scanner->current_indent++;
			lexer->mark_end(lexer);
			lexer->result_symbol = INDENT;
#ifdef DEBUG
			printf("  I : current_indent %d\n", scanner->current_indent);
#endif
			return true;
		}
	}

	if (valid_symbols[NEWLINE] && lexer->lookahead == '\n') {
		lexer->advance(lexer, false);

		while (lexer->lookahead == '\n') {
			lexer->advance(lexer, true);
		}

		lexer->mark_end(lexer);
		lexer->result_symbol = NEWLINE;
		return true;
	}

	return false;
}
