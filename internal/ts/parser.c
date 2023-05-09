#include <tree_sitter/parser.h>

#if defined(__GNUC__) || defined(__clang__)
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wmissing-field-initializers"
#endif

#define LANGUAGE_VERSION 13
#define STATE_COUNT 267
#define LARGE_STATE_COUNT 2
#define SYMBOL_COUNT 112
#define ALIAS_COUNT 0
#define TOKEN_COUNT 61
#define EXTERNAL_TOKEN_COUNT 3
#define FIELD_COUNT 0
#define MAX_ALIAS_SEQUENCE_LENGTH 7
#define PRODUCTION_ID_COUNT 1

enum {
  sym_comment = 1,
  aux_sym_identifier_token1 = 2,
  anon_sym_access = 3,
  anon_sym_add_DASHenable = 4,
  anon_sym_atomic = 5,
  anon_sym_byte_DASHwrite_DASHenable = 6,
  anon_sym_clear = 7,
  anon_sym_delay = 8,
  anon_sym_enable_DASHinit_DASHvalue = 9,
  anon_sym_enable_DASHreset_DASHvalue = 10,
  anon_sym_groups = 11,
  anon_sym_init_DASHvalue = 12,
  anon_sym_in_DASHtrigger = 13,
  anon_sym_masters = 14,
  anon_sym_out_DASHtrigger = 15,
  anon_sym_range = 16,
  anon_sym_read_DASHlatency = 17,
  anon_sym_read_DASHvalue = 18,
  anon_sym_reset = 19,
  anon_sym_reset_DASHvalue = 20,
  anon_sym_size = 21,
  anon_sym_width = 22,
  anon_sym_DOT = 23,
  anon_sym_import = 24,
  anon_sym_const = 25,
  anon_sym_EQ = 26,
  anon_sym_COMMA = 27,
  anon_sym_LPAREN = 28,
  anon_sym_RPAREN = 29,
  anon_sym_type = 30,
  anon_sym_SEMI = 31,
  anon_sym_LBRACK = 32,
  anon_sym_RBRACK = 33,
  anon_sym_DQUOTE = 34,
  aux_sym_string_literal_token1 = 35,
  sym_zero_literal = 36,
  sym_decimal_literal = 37,
  sym_binary_literal = 38,
  sym_octal_literal = 39,
  sym_hex_literal = 40,
  sym_bit_literal = 41,
  anon_sym_ns = 42,
  anon_sym_us = 43,
  anon_sym_ms = 44,
  anon_sym_s = 45,
  anon_sym_and = 46,
  anon_sym_or = 47,
  anon_sym_PLUS = 48,
  anon_sym_DASH = 49,
  anon_sym_STAR = 50,
  anon_sym_SLASH = 51,
  anon_sym_PERCENT = 52,
  anon_sym_STAR_STAR = 53,
  anon_sym_LT_LT = 54,
  anon_sym_GT_GT = 55,
  anon_sym_true = 56,
  anon_sym_false = 57,
  sym__indent = 58,
  sym__dedent = 59,
  sym__newline = 60,
  sym_description = 61,
  sym_identifier = 62,
  sym_declared_identifier = 63,
  sym_property_identifier = 64,
  sym_qualified_identifier = 65,
  sym__meta_statement = 66,
  sym__import_statement = 67,
  sym_single_import_statement = 68,
  sym_multi_import_statement = 69,
  sym__constant_definition = 70,
  sym_single_constant_definition = 71,
  sym_multi_constant_definition = 72,
  sym__parameter = 73,
  sym__parameters = 74,
  sym_parameter_list = 75,
  sym__type_definition = 76,
  sym_single_line_type_definition = 77,
  sym_multi_line_type_definition = 78,
  sym_element_body = 79,
  sym__property_assignment = 80,
  sym_property_assignments = 81,
  sym__instantiation = 82,
  sym__argument = 83,
  sym__arguments = 84,
  sym_argument_list = 85,
  sym_array_marker = 86,
  sym_single_line_instantiation = 87,
  sym_multi_line_instantiation = 88,
  sym_string_literal = 89,
  sym__integer_literal = 90,
  sym__zero_or_decimal_literal = 91,
  sym_float_literal = 92,
  sym_time_unit = 93,
  sym_time_literal = 94,
  sym_unary_operation = 95,
  sym_binary_operation = 96,
  sym_call = 97,
  sym_parenthesized_expression = 98,
  sym_subscript = 99,
  sym_primary_expression = 100,
  sym__expressions = 101,
  sym_expression_list = 102,
  sym__expression = 103,
  aux_sym_description_repeat1 = 104,
  aux_sym_multi_import_statement_repeat1 = 105,
  aux_sym_multi_constant_definition_repeat1 = 106,
  aux_sym__parameters_repeat1 = 107,
  aux_sym_element_body_repeat1 = 108,
  aux_sym_property_assignments_repeat1 = 109,
  aux_sym__arguments_repeat1 = 110,
  aux_sym__expressions_repeat1 = 111,
};

static const char * const ts_symbol_names[] = {
  [ts_builtin_sym_end] = "end",
  [sym_comment] = "comment",
  [aux_sym_identifier_token1] = "identifier_token1",
  [anon_sym_access] = "access",
  [anon_sym_add_DASHenable] = "add-enable",
  [anon_sym_atomic] = "atomic",
  [anon_sym_byte_DASHwrite_DASHenable] = "byte-write-enable",
  [anon_sym_clear] = "clear",
  [anon_sym_delay] = "delay",
  [anon_sym_enable_DASHinit_DASHvalue] = "enable-init-value",
  [anon_sym_enable_DASHreset_DASHvalue] = "enable-reset-value",
  [anon_sym_groups] = "groups",
  [anon_sym_init_DASHvalue] = "init-value",
  [anon_sym_in_DASHtrigger] = "in-trigger",
  [anon_sym_masters] = "masters",
  [anon_sym_out_DASHtrigger] = "out-trigger",
  [anon_sym_range] = "range",
  [anon_sym_read_DASHlatency] = "read-latency",
  [anon_sym_read_DASHvalue] = "read-value",
  [anon_sym_reset] = "reset",
  [anon_sym_reset_DASHvalue] = "reset-value",
  [anon_sym_size] = "size",
  [anon_sym_width] = "width",
  [anon_sym_DOT] = ".",
  [anon_sym_import] = "import",
  [anon_sym_const] = "const",
  [anon_sym_EQ] = "=",
  [anon_sym_COMMA] = ",",
  [anon_sym_LPAREN] = "(",
  [anon_sym_RPAREN] = ")",
  [anon_sym_type] = "type",
  [anon_sym_SEMI] = ";",
  [anon_sym_LBRACK] = "[",
  [anon_sym_RBRACK] = "]",
  [anon_sym_DQUOTE] = "\"",
  [aux_sym_string_literal_token1] = "string_literal_token1",
  [sym_zero_literal] = "zero_literal",
  [sym_decimal_literal] = "decimal_literal",
  [sym_binary_literal] = "binary_literal",
  [sym_octal_literal] = "octal_literal",
  [sym_hex_literal] = "hex_literal",
  [sym_bit_literal] = "bit_literal",
  [anon_sym_ns] = "ns",
  [anon_sym_us] = "us",
  [anon_sym_ms] = "ms",
  [anon_sym_s] = "s",
  [anon_sym_and] = "and",
  [anon_sym_or] = "or",
  [anon_sym_PLUS] = "+",
  [anon_sym_DASH] = "-",
  [anon_sym_STAR] = "*",
  [anon_sym_SLASH] = "/",
  [anon_sym_PERCENT] = "%",
  [anon_sym_STAR_STAR] = "**",
  [anon_sym_LT_LT] = "<<",
  [anon_sym_GT_GT] = ">>",
  [anon_sym_true] = "true",
  [anon_sym_false] = "false",
  [sym__indent] = "_indent",
  [sym__dedent] = "_dedent",
  [sym__newline] = "_newline",
  [sym_description] = "description",
  [sym_identifier] = "identifier",
  [sym_declared_identifier] = "declared_identifier",
  [sym_property_identifier] = "property_identifier",
  [sym_qualified_identifier] = "qualified_identifier",
  [sym__meta_statement] = "_meta_statement",
  [sym__import_statement] = "_import_statement",
  [sym_single_import_statement] = "single_import_statement",
  [sym_multi_import_statement] = "multi_import_statement",
  [sym__constant_definition] = "_constant_definition",
  [sym_single_constant_definition] = "single_constant_definition",
  [sym_multi_constant_definition] = "multi_constant_definition",
  [sym__parameter] = "_parameter",
  [sym__parameters] = "_parameters",
  [sym_parameter_list] = "parameter_list",
  [sym__type_definition] = "_type_definition",
  [sym_single_line_type_definition] = "single_line_type_definition",
  [sym_multi_line_type_definition] = "multi_line_type_definition",
  [sym_element_body] = "element_body",
  [sym__property_assignment] = "_property_assignment",
  [sym_property_assignments] = "property_assignments",
  [sym__instantiation] = "_instantiation",
  [sym__argument] = "_argument",
  [sym__arguments] = "_arguments",
  [sym_argument_list] = "argument_list",
  [sym_array_marker] = "array_marker",
  [sym_single_line_instantiation] = "single_line_instantiation",
  [sym_multi_line_instantiation] = "multi_line_instantiation",
  [sym_string_literal] = "string_literal",
  [sym__integer_literal] = "_integer_literal",
  [sym__zero_or_decimal_literal] = "_zero_or_decimal_literal",
  [sym_float_literal] = "float_literal",
  [sym_time_unit] = "time_unit",
  [sym_time_literal] = "time_literal",
  [sym_unary_operation] = "unary_operation",
  [sym_binary_operation] = "binary_operation",
  [sym_call] = "call",
  [sym_parenthesized_expression] = "parenthesized_expression",
  [sym_subscript] = "subscript",
  [sym_primary_expression] = "primary_expression",
  [sym__expressions] = "_expressions",
  [sym_expression_list] = "expression_list",
  [sym__expression] = "_expression",
  [aux_sym_description_repeat1] = "description_repeat1",
  [aux_sym_multi_import_statement_repeat1] = "multi_import_statement_repeat1",
  [aux_sym_multi_constant_definition_repeat1] = "multi_constant_definition_repeat1",
  [aux_sym__parameters_repeat1] = "_parameters_repeat1",
  [aux_sym_element_body_repeat1] = "element_body_repeat1",
  [aux_sym_property_assignments_repeat1] = "property_assignments_repeat1",
  [aux_sym__arguments_repeat1] = "_arguments_repeat1",
  [aux_sym__expressions_repeat1] = "_expressions_repeat1",
};

static const TSSymbol ts_symbol_map[] = {
  [ts_builtin_sym_end] = ts_builtin_sym_end,
  [sym_comment] = sym_comment,
  [aux_sym_identifier_token1] = aux_sym_identifier_token1,
  [anon_sym_access] = anon_sym_access,
  [anon_sym_add_DASHenable] = anon_sym_add_DASHenable,
  [anon_sym_atomic] = anon_sym_atomic,
  [anon_sym_byte_DASHwrite_DASHenable] = anon_sym_byte_DASHwrite_DASHenable,
  [anon_sym_clear] = anon_sym_clear,
  [anon_sym_delay] = anon_sym_delay,
  [anon_sym_enable_DASHinit_DASHvalue] = anon_sym_enable_DASHinit_DASHvalue,
  [anon_sym_enable_DASHreset_DASHvalue] = anon_sym_enable_DASHreset_DASHvalue,
  [anon_sym_groups] = anon_sym_groups,
  [anon_sym_init_DASHvalue] = anon_sym_init_DASHvalue,
  [anon_sym_in_DASHtrigger] = anon_sym_in_DASHtrigger,
  [anon_sym_masters] = anon_sym_masters,
  [anon_sym_out_DASHtrigger] = anon_sym_out_DASHtrigger,
  [anon_sym_range] = anon_sym_range,
  [anon_sym_read_DASHlatency] = anon_sym_read_DASHlatency,
  [anon_sym_read_DASHvalue] = anon_sym_read_DASHvalue,
  [anon_sym_reset] = anon_sym_reset,
  [anon_sym_reset_DASHvalue] = anon_sym_reset_DASHvalue,
  [anon_sym_size] = anon_sym_size,
  [anon_sym_width] = anon_sym_width,
  [anon_sym_DOT] = anon_sym_DOT,
  [anon_sym_import] = anon_sym_import,
  [anon_sym_const] = anon_sym_const,
  [anon_sym_EQ] = anon_sym_EQ,
  [anon_sym_COMMA] = anon_sym_COMMA,
  [anon_sym_LPAREN] = anon_sym_LPAREN,
  [anon_sym_RPAREN] = anon_sym_RPAREN,
  [anon_sym_type] = anon_sym_type,
  [anon_sym_SEMI] = anon_sym_SEMI,
  [anon_sym_LBRACK] = anon_sym_LBRACK,
  [anon_sym_RBRACK] = anon_sym_RBRACK,
  [anon_sym_DQUOTE] = anon_sym_DQUOTE,
  [aux_sym_string_literal_token1] = aux_sym_string_literal_token1,
  [sym_zero_literal] = sym_zero_literal,
  [sym_decimal_literal] = sym_decimal_literal,
  [sym_binary_literal] = sym_binary_literal,
  [sym_octal_literal] = sym_octal_literal,
  [sym_hex_literal] = sym_hex_literal,
  [sym_bit_literal] = sym_bit_literal,
  [anon_sym_ns] = anon_sym_ns,
  [anon_sym_us] = anon_sym_us,
  [anon_sym_ms] = anon_sym_ms,
  [anon_sym_s] = anon_sym_s,
  [anon_sym_and] = anon_sym_and,
  [anon_sym_or] = anon_sym_or,
  [anon_sym_PLUS] = anon_sym_PLUS,
  [anon_sym_DASH] = anon_sym_DASH,
  [anon_sym_STAR] = anon_sym_STAR,
  [anon_sym_SLASH] = anon_sym_SLASH,
  [anon_sym_PERCENT] = anon_sym_PERCENT,
  [anon_sym_STAR_STAR] = anon_sym_STAR_STAR,
  [anon_sym_LT_LT] = anon_sym_LT_LT,
  [anon_sym_GT_GT] = anon_sym_GT_GT,
  [anon_sym_true] = anon_sym_true,
  [anon_sym_false] = anon_sym_false,
  [sym__indent] = sym__indent,
  [sym__dedent] = sym__dedent,
  [sym__newline] = sym__newline,
  [sym_description] = sym_description,
  [sym_identifier] = sym_identifier,
  [sym_declared_identifier] = sym_declared_identifier,
  [sym_property_identifier] = sym_property_identifier,
  [sym_qualified_identifier] = sym_qualified_identifier,
  [sym__meta_statement] = sym__meta_statement,
  [sym__import_statement] = sym__import_statement,
  [sym_single_import_statement] = sym_single_import_statement,
  [sym_multi_import_statement] = sym_multi_import_statement,
  [sym__constant_definition] = sym__constant_definition,
  [sym_single_constant_definition] = sym_single_constant_definition,
  [sym_multi_constant_definition] = sym_multi_constant_definition,
  [sym__parameter] = sym__parameter,
  [sym__parameters] = sym__parameters,
  [sym_parameter_list] = sym_parameter_list,
  [sym__type_definition] = sym__type_definition,
  [sym_single_line_type_definition] = sym_single_line_type_definition,
  [sym_multi_line_type_definition] = sym_multi_line_type_definition,
  [sym_element_body] = sym_element_body,
  [sym__property_assignment] = sym__property_assignment,
  [sym_property_assignments] = sym_property_assignments,
  [sym__instantiation] = sym__instantiation,
  [sym__argument] = sym__argument,
  [sym__arguments] = sym__arguments,
  [sym_argument_list] = sym_argument_list,
  [sym_array_marker] = sym_array_marker,
  [sym_single_line_instantiation] = sym_single_line_instantiation,
  [sym_multi_line_instantiation] = sym_multi_line_instantiation,
  [sym_string_literal] = sym_string_literal,
  [sym__integer_literal] = sym__integer_literal,
  [sym__zero_or_decimal_literal] = sym__zero_or_decimal_literal,
  [sym_float_literal] = sym_float_literal,
  [sym_time_unit] = sym_time_unit,
  [sym_time_literal] = sym_time_literal,
  [sym_unary_operation] = sym_unary_operation,
  [sym_binary_operation] = sym_binary_operation,
  [sym_call] = sym_call,
  [sym_parenthesized_expression] = sym_parenthesized_expression,
  [sym_subscript] = sym_subscript,
  [sym_primary_expression] = sym_primary_expression,
  [sym__expressions] = sym__expressions,
  [sym_expression_list] = sym_expression_list,
  [sym__expression] = sym__expression,
  [aux_sym_description_repeat1] = aux_sym_description_repeat1,
  [aux_sym_multi_import_statement_repeat1] = aux_sym_multi_import_statement_repeat1,
  [aux_sym_multi_constant_definition_repeat1] = aux_sym_multi_constant_definition_repeat1,
  [aux_sym__parameters_repeat1] = aux_sym__parameters_repeat1,
  [aux_sym_element_body_repeat1] = aux_sym_element_body_repeat1,
  [aux_sym_property_assignments_repeat1] = aux_sym_property_assignments_repeat1,
  [aux_sym__arguments_repeat1] = aux_sym__arguments_repeat1,
  [aux_sym__expressions_repeat1] = aux_sym__expressions_repeat1,
};

static const TSSymbolMetadata ts_symbol_metadata[] = {
  [ts_builtin_sym_end] = {
    .visible = false,
    .named = true,
  },
  [sym_comment] = {
    .visible = true,
    .named = true,
  },
  [aux_sym_identifier_token1] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_access] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_add_DASHenable] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_atomic] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_byte_DASHwrite_DASHenable] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_clear] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_delay] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_enable_DASHinit_DASHvalue] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_enable_DASHreset_DASHvalue] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_groups] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_init_DASHvalue] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_in_DASHtrigger] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_masters] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_out_DASHtrigger] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_range] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_read_DASHlatency] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_read_DASHvalue] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_reset] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_reset_DASHvalue] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_size] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_width] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DOT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_import] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_const] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COMMA] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LPAREN] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_RPAREN] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_type] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SEMI] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LBRACK] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_RBRACK] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DQUOTE] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_string_literal_token1] = {
    .visible = false,
    .named = false,
  },
  [sym_zero_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_decimal_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_binary_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_octal_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_hex_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_bit_literal] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_ns] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_us] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_ms] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_s] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_and] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_or] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PLUS] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DASH] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_STAR] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SLASH] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PERCENT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_STAR_STAR] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LT_LT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_GT_GT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_true] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_false] = {
    .visible = true,
    .named = false,
  },
  [sym__indent] = {
    .visible = false,
    .named = true,
  },
  [sym__dedent] = {
    .visible = false,
    .named = true,
  },
  [sym__newline] = {
    .visible = false,
    .named = true,
  },
  [sym_description] = {
    .visible = true,
    .named = true,
  },
  [sym_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym_declared_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym_property_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym_qualified_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym__meta_statement] = {
    .visible = false,
    .named = true,
  },
  [sym__import_statement] = {
    .visible = false,
    .named = true,
  },
  [sym_single_import_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_multi_import_statement] = {
    .visible = true,
    .named = true,
  },
  [sym__constant_definition] = {
    .visible = false,
    .named = true,
  },
  [sym_single_constant_definition] = {
    .visible = true,
    .named = true,
  },
  [sym_multi_constant_definition] = {
    .visible = true,
    .named = true,
  },
  [sym__parameter] = {
    .visible = false,
    .named = true,
  },
  [sym__parameters] = {
    .visible = false,
    .named = true,
  },
  [sym_parameter_list] = {
    .visible = true,
    .named = true,
  },
  [sym__type_definition] = {
    .visible = false,
    .named = true,
  },
  [sym_single_line_type_definition] = {
    .visible = true,
    .named = true,
  },
  [sym_multi_line_type_definition] = {
    .visible = true,
    .named = true,
  },
  [sym_element_body] = {
    .visible = true,
    .named = true,
  },
  [sym__property_assignment] = {
    .visible = false,
    .named = true,
  },
  [sym_property_assignments] = {
    .visible = true,
    .named = true,
  },
  [sym__instantiation] = {
    .visible = false,
    .named = true,
  },
  [sym__argument] = {
    .visible = false,
    .named = true,
  },
  [sym__arguments] = {
    .visible = false,
    .named = true,
  },
  [sym_argument_list] = {
    .visible = true,
    .named = true,
  },
  [sym_array_marker] = {
    .visible = true,
    .named = true,
  },
  [sym_single_line_instantiation] = {
    .visible = true,
    .named = true,
  },
  [sym_multi_line_instantiation] = {
    .visible = true,
    .named = true,
  },
  [sym_string_literal] = {
    .visible = true,
    .named = true,
  },
  [sym__integer_literal] = {
    .visible = false,
    .named = true,
  },
  [sym__zero_or_decimal_literal] = {
    .visible = false,
    .named = true,
  },
  [sym_float_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_time_unit] = {
    .visible = true,
    .named = true,
  },
  [sym_time_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_unary_operation] = {
    .visible = true,
    .named = true,
  },
  [sym_binary_operation] = {
    .visible = true,
    .named = true,
  },
  [sym_call] = {
    .visible = true,
    .named = true,
  },
  [sym_parenthesized_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_subscript] = {
    .visible = true,
    .named = true,
  },
  [sym_primary_expression] = {
    .visible = true,
    .named = true,
  },
  [sym__expressions] = {
    .visible = false,
    .named = true,
  },
  [sym_expression_list] = {
    .visible = true,
    .named = true,
  },
  [sym__expression] = {
    .visible = false,
    .named = true,
  },
  [aux_sym_description_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_multi_import_statement_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_multi_constant_definition_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__parameters_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_element_body_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_property_assignments_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__arguments_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__expressions_repeat1] = {
    .visible = false,
    .named = false,
  },
};

static const TSSymbol ts_alias_sequences[PRODUCTION_ID_COUNT][MAX_ALIAS_SEQUENCE_LENGTH] = {
  [0] = {0},
};

static const uint16_t ts_non_terminal_alias_map[] = {
  0,
};

static inline bool sym_bit_literal_character_set_1(int32_t c) {
  return (c < 'Z'
    ? (c < 'H'
      ? (c < '0'
        ? c == '-'
        : (c <= '9' || (c >= 'A' && c <= 'F')))
      : (c <= 'H' || (c < 'U'
        ? c == 'L'
        : (c <= 'U' || (c >= 'W' && c <= 'X')))))
    : (c <= 'Z' || (c < 'u'
      ? (c < 'h'
        ? (c >= 'a' && c <= 'f')
        : (c <= 'h' || c == 'l'))
      : (c <= 'u' || (c < 'z'
        ? (c >= 'w' && c <= 'x')
        : c <= 'z')))));
}

static bool ts_lex(TSLexer *lexer, TSStateId state) {
  START_LEXER();
  eof = lexer->eof(lexer);
  switch (state) {
    case 0:
      if (eof) ADVANCE(88);
      if (lookahead == '"') ADVANCE(215);
      if (lookahead == '#') ADVANCE(90);
      if (lookahead == '%') ADVANCE(238);
      if (lookahead == '(') ADVANCE(209);
      if (lookahead == ')') ADVANCE(210);
      if (lookahead == '*') ADVANCE(236);
      if (lookahead == '+') ADVANCE(234);
      if (lookahead == ',') ADVANCE(208);
      if (lookahead == '-') ADVANCE(235);
      if (lookahead == '.') ADVANCE(204);
      if (lookahead == '/') ADVANCE(237);
      if (lookahead == '0') ADVANCE(218);
      if (lookahead == ';') ADVANCE(212);
      if (lookahead == '<') ADVANCE(9);
      if (lookahead == '=') ADVANCE(207);
      if (lookahead == '>') ADVANCE(10);
      if (lookahead == 'B' ||
          lookahead == 'O' ||
          lookahead == 'X' ||
          lookahead == 'x') ADVANCE(93);
      if (lookahead == '[') ADVANCE(213);
      if (lookahead == ']') ADVANCE(214);
      if (lookahead == 'a') ADVANCE(101);
      if (lookahead == 'b') ADVANCE(92);
      if (lookahead == 'c') ADVANCE(104);
      if (lookahead == 'd') ADVANCE(116);
      if (lookahead == 'e') ADVANCE(140);
      if (lookahead == 'f') ADVANCE(181);
      if (lookahead == 'g') ADVANCE(151);
      if (lookahead == 'i') ADVANCE(105);
      if (lookahead == 'm') ADVANCE(107);
      if (lookahead == 'n') ADVANCE(154);
      if (lookahead == 'o') ADVANCE(91);
      if (lookahead == 'r') ADVANCE(103);
      if (lookahead == 's') ADVANCE(231);
      if (lookahead == 't') ADVANCE(106);
      if (lookahead == 'u') ADVANCE(155);
      if (lookahead == 'w') ADVANCE(132);
      if (lookahead == '\t' ||
          lookahead == '\n' ||
          lookahead == '\r' ||
          lookahead == ' ') SKIP(0)
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('h' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(219);
      END_STATE();
    case 1:
      if (lookahead == '"') ADVANCE(215);
      if (lookahead == '#') ADVANCE(90);
      if (lookahead == '(') ADVANCE(209);
      if (lookahead == ')') ADVANCE(210);
      if (lookahead == '+') ADVANCE(234);
      if (lookahead == '-') ADVANCE(235);
      if (lookahead == '.') ADVANCE(204);
      if (lookahead == '0') ADVANCE(218);
      if (lookahead == 'B' ||
          lookahead == 'O' ||
          lookahead == 'X' ||
          lookahead == 'b' ||
          lookahead == 'o' ||
          lookahead == 'x') ADVANCE(93);
      if (lookahead == '[') ADVANCE(213);
      if (lookahead == ']') ADVANCE(214);
      if (lookahead == 'f') ADVANCE(181);
      if (lookahead == 't') ADVANCE(149);
      if (lookahead == '\t' ||
          lookahead == '\n' ||
          lookahead == '\r' ||
          lookahead == ' ') SKIP(1)
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(219);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 2:
      if (lookahead == '"') ADVANCE(215);
      if (lookahead == '#') ADVANCE(90);
      if (lookahead == '(') ADVANCE(209);
      if (lookahead == ')') ADVANCE(210);
      if (lookahead == ',') ADVANCE(208);
      if (lookahead == '=') ADVANCE(207);
      if (lookahead == '[') ADVANCE(213);
      if (lookahead == '\t' ||
          lookahead == '\n' ||
          lookahead == '\r' ||
          lookahead == ' ') SKIP(2)
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 3:
      if (lookahead == '"') ADVANCE(223);
      if (sym_bit_literal_character_set_1(lookahead)) ADVANCE(3);
      END_STATE();
    case 4:
      if (lookahead == '#') ADVANCE(90);
      if (lookahead == '%') ADVANCE(238);
      if (lookahead == ')') ADVANCE(210);
      if (lookahead == '*') ADVANCE(236);
      if (lookahead == '+') ADVANCE(234);
      if (lookahead == ',') ADVANCE(208);
      if (lookahead == '-') ADVANCE(235);
      if (lookahead == '.') ADVANCE(204);
      if (lookahead == '/') ADVANCE(237);
      if (lookahead == ';') ADVANCE(212);
      if (lookahead == '<') ADVANCE(9);
      if (lookahead == '>') ADVANCE(10);
      if (lookahead == ']') ADVANCE(214);
      if (lookahead == 'm') ADVANCE(63);
      if (lookahead == 'n') ADVANCE(64);
      if (lookahead == 's') ADVANCE(230);
      if (lookahead == 'u') ADVANCE(65);
      if (lookahead == '\t' ||
          lookahead == '\n' ||
          lookahead == '\r' ||
          lookahead == ' ') SKIP(4)
      END_STATE();
    case 5:
      if (lookahead == '#') ADVANCE(90);
      if (lookahead == 'a') ADVANCE(102);
      if (lookahead == 'b') ADVANCE(174);
      if (lookahead == 'c') ADVANCE(104);
      if (lookahead == 'd') ADVANCE(116);
      if (lookahead == 'e') ADVANCE(140);
      if (lookahead == 'g') ADVANCE(151);
      if (lookahead == 'i') ADVANCE(138);
      if (lookahead == 'm') ADVANCE(182);
      if (lookahead == 'o') ADVANCE(171);
      if (lookahead == 'r') ADVANCE(103);
      if (lookahead == 's') ADVANCE(130);
      if (lookahead == 't') ADVANCE(176);
      if (lookahead == 'w') ADVANCE(132);
      if (lookahead == '\t' ||
          lookahead == '\n' ||
          lookahead == '\r' ||
          lookahead == ' ') SKIP(5)
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('f' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 6:
      if (lookahead == '-') ADVANCE(80);
      END_STATE();
    case 7:
      if (lookahead == '-') ADVANCE(81);
      END_STATE();
    case 8:
      if (lookahead == '-') ADVANCE(35);
      END_STATE();
    case 9:
      if (lookahead == '<') ADVANCE(240);
      END_STATE();
    case 10:
      if (lookahead == '>') ADVANCE(241);
      END_STATE();
    case 11:
      if (lookahead == 'a') ADVANCE(19);
      END_STATE();
    case 12:
      if (lookahead == 'a') ADVANCE(46);
      END_STATE();
    case 13:
      if (lookahead == 'a') ADVANCE(70);
      END_STATE();
    case 14:
      if (lookahead == 'a') ADVANCE(20);
      END_STATE();
    case 15:
      if (lookahead == 'a') ADVANCE(50);
      END_STATE();
    case 16:
      if (lookahead == 'a') ADVANCE(51);
      END_STATE();
    case 17:
      if (lookahead == 'a') ADVANCE(52);
      END_STATE();
    case 18:
      if (lookahead == 'a') ADVANCE(53);
      END_STATE();
    case 19:
      if (lookahead == 'b') ADVANCE(48);
      END_STATE();
    case 20:
      if (lookahead == 'b') ADVANCE(49);
      END_STATE();
    case 21:
      if (lookahead == 'c') ADVANCE(83);
      END_STATE();
    case 22:
      if (lookahead == 'e') ADVANCE(54);
      END_STATE();
    case 23:
      if (lookahead == 'e') ADVANCE(66);
      END_STATE();
    case 24:
      if (lookahead == 'e') ADVANCE(185);
      END_STATE();
    case 25:
      if (lookahead == 'e') ADVANCE(8);
      END_STATE();
    case 26:
      if (lookahead == 'e') ADVANCE(193);
      END_STATE();
    case 27:
      if (lookahead == 'e') ADVANCE(199);
      END_STATE();
    case 28:
      if (lookahead == 'e') ADVANCE(201);
      END_STATE();
    case 29:
      if (lookahead == 'e') ADVANCE(187);
      END_STATE();
    case 30:
      if (lookahead == 'e') ADVANCE(190);
      END_STATE();
    case 31:
      if (lookahead == 'e') ADVANCE(191);
      END_STATE();
    case 32:
      if (lookahead == 'e') ADVANCE(59);
      END_STATE();
    case 33:
      if (lookahead == 'e') ADVANCE(60);
      END_STATE();
    case 34:
      if (lookahead == 'e') ADVANCE(55);
      END_STATE();
    case 35:
      if (lookahead == 'e') ADVANCE(57);
      END_STATE();
    case 36:
      if (lookahead == 'e') ADVANCE(72);
      END_STATE();
    case 37:
      if (lookahead == 'g') ADVANCE(38);
      END_STATE();
    case 38:
      if (lookahead == 'g') ADVANCE(32);
      END_STATE();
    case 39:
      if (lookahead == 'g') ADVANCE(33);
      END_STATE();
    case 40:
      if (lookahead == 'g') ADVANCE(39);
      END_STATE();
    case 41:
      if (lookahead == 'i') ADVANCE(37);
      END_STATE();
    case 42:
      if (lookahead == 'i') ADVANCE(56);
      if (lookahead == 'r') ADVANCE(23);
      END_STATE();
    case 43:
      if (lookahead == 'i') ADVANCE(69);
      END_STATE();
    case 44:
      if (lookahead == 'i') ADVANCE(68);
      END_STATE();
    case 45:
      if (lookahead == 'i') ADVANCE(40);
      END_STATE();
    case 46:
      if (lookahead == 'l') ADVANCE(73);
      END_STATE();
    case 47:
      if (lookahead == 'l') ADVANCE(13);
      if (lookahead == 'v') ADVANCE(15);
      END_STATE();
    case 48:
      if (lookahead == 'l') ADVANCE(24);
      END_STATE();
    case 49:
      if (lookahead == 'l') ADVANCE(29);
      END_STATE();
    case 50:
      if (lookahead == 'l') ADVANCE(74);
      END_STATE();
    case 51:
      if (lookahead == 'l') ADVANCE(75);
      END_STATE();
    case 52:
      if (lookahead == 'l') ADVANCE(76);
      END_STATE();
    case 53:
      if (lookahead == 'l') ADVANCE(77);
      END_STATE();
    case 54:
      if (lookahead == 'n') ADVANCE(11);
      END_STATE();
    case 55:
      if (lookahead == 'n') ADVANCE(21);
      END_STATE();
    case 56:
      if (lookahead == 'n') ADVANCE(44);
      END_STATE();
    case 57:
      if (lookahead == 'n') ADVANCE(14);
      END_STATE();
    case 58:
      if (lookahead == 'r') ADVANCE(41);
      END_STATE();
    case 59:
      if (lookahead == 'r') ADVANCE(194);
      END_STATE();
    case 60:
      if (lookahead == 'r') ADVANCE(196);
      END_STATE();
    case 61:
      if (lookahead == 'r') ADVANCE(43);
      END_STATE();
    case 62:
      if (lookahead == 'r') ADVANCE(45);
      END_STATE();
    case 63:
      if (lookahead == 's') ADVANCE(228);
      END_STATE();
    case 64:
      if (lookahead == 's') ADVANCE(224);
      END_STATE();
    case 65:
      if (lookahead == 's') ADVANCE(226);
      END_STATE();
    case 66:
      if (lookahead == 's') ADVANCE(36);
      END_STATE();
    case 67:
      if (lookahead == 't') ADVANCE(58);
      END_STATE();
    case 68:
      if (lookahead == 't') ADVANCE(6);
      END_STATE();
    case 69:
      if (lookahead == 't') ADVANCE(25);
      END_STATE();
    case 70:
      if (lookahead == 't') ADVANCE(34);
      END_STATE();
    case 71:
      if (lookahead == 't') ADVANCE(62);
      END_STATE();
    case 72:
      if (lookahead == 't') ADVANCE(7);
      END_STATE();
    case 73:
      if (lookahead == 'u') ADVANCE(26);
      END_STATE();
    case 74:
      if (lookahead == 'u') ADVANCE(27);
      END_STATE();
    case 75:
      if (lookahead == 'u') ADVANCE(28);
      END_STATE();
    case 76:
      if (lookahead == 'u') ADVANCE(30);
      END_STATE();
    case 77:
      if (lookahead == 'u') ADVANCE(31);
      END_STATE();
    case 78:
      if (lookahead == 'v') ADVANCE(12);
      END_STATE();
    case 79:
      if (lookahead == 'v') ADVANCE(16);
      END_STATE();
    case 80:
      if (lookahead == 'v') ADVANCE(17);
      END_STATE();
    case 81:
      if (lookahead == 'v') ADVANCE(18);
      END_STATE();
    case 82:
      if (lookahead == 'w') ADVANCE(61);
      END_STATE();
    case 83:
      if (lookahead == 'y') ADVANCE(198);
      END_STATE();
    case 84:
      if (lookahead == '0' ||
          lookahead == '1') ADVANCE(220);
      END_STATE();
    case 85:
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(221);
      END_STATE();
    case 86:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f') ||
          lookahead == '|') ADVANCE(222);
      END_STATE();
    case 87:
      if (eof) ADVANCE(88);
      if (lookahead == '#') ADVANCE(90);
      if (lookahead == 'c') ADVANCE(143);
      if (lookahead == 'i') ADVANCE(136);
      if (lookahead == 't') ADVANCE(176);
      if (lookahead == '\t' ||
          lookahead == '\n' ||
          lookahead == '\r' ||
          lookahead == ' ') SKIP(87)
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 88:
      ACCEPT_TOKEN(ts_builtin_sym_end);
      END_STATE();
    case 89:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead == '\n') ADVANCE(217);
      if (lookahead == '"') ADVANCE(90);
      if (lookahead != 0) ADVANCE(89);
      END_STATE();
    case 90:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(90);
      END_STATE();
    case 91:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (lookahead == '"') ADVANCE(3);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'q') ||
          lookahead == 's' ||
          lookahead == 't' ||
          ('v' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'r') ADVANCE(233);
      if (lookahead == 'u') ADVANCE(167);
      END_STATE();
    case 92:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (lookahead == '"') ADVANCE(3);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'x') ||
          lookahead == 'z') ADVANCE(183);
      if (lookahead == 'y') ADVANCE(168);
      END_STATE();
    case 93:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (lookahead == '"') ADVANCE(3);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 94:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (lookahead == '-') ADVANCE(67);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'h') ||
          ('j' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'i') ADVANCE(169);
      END_STATE();
    case 95:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (lookahead == '-') ADVANCE(22);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 96:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (lookahead == '-') ADVANCE(82);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 97:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (lookahead == '-') ADVANCE(78);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 98:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (lookahead == '-') ADVANCE(47);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 99:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (lookahead == '-') ADVANCE(42);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 100:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (lookahead == '-') ADVANCE(71);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 101:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          lookahead == 'a' ||
          lookahead == 'b' ||
          ('e' <= lookahead && lookahead <= 'm') ||
          ('o' <= lookahead && lookahead <= 's') ||
          ('u' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'c') ADVANCE(111);
      if (lookahead == 'd') ADVANCE(112);
      if (lookahead == 'n') ADVANCE(113);
      if (lookahead == 't') ADVANCE(142);
      END_STATE();
    case 102:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          lookahead == 'a' ||
          lookahead == 'b' ||
          ('e' <= lookahead && lookahead <= 's') ||
          ('u' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'c') ADVANCE(111);
      if (lookahead == 'd') ADVANCE(112);
      if (lookahead == 't') ADVANCE(142);
      END_STATE();
    case 103:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'd') ||
          ('f' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'a') ADVANCE(139);
      if (lookahead == 'e') ADVANCE(108);
      END_STATE();
    case 104:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'k') ||
          lookahead == 'm' ||
          lookahead == 'n' ||
          ('p' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'l') ADVANCE(122);
      if (lookahead == 'o') ADVANCE(141);
      END_STATE();
    case 105:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'l') ||
          ('o' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'm') ADVANCE(146);
      if (lookahead == 'n') ADVANCE(94);
      END_STATE();
    case 106:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'q') ||
          ('s' <= lookahead && lookahead <= 'x') ||
          lookahead == 'z') ADVANCE(183);
      if (lookahead == 'r') ADVANCE(173);
      if (lookahead == 'y') ADVANCE(147);
      END_STATE();
    case 107:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'r') ||
          ('t' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'a') ADVANCE(162);
      if (lookahead == 's') ADVANCE(229);
      END_STATE();
    case 108:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'r') ||
          ('t' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'a') ADVANCE(115);
      if (lookahead == 's') ADVANCE(125);
      END_STATE();
    case 109:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          lookahead == 'a' ||
          ('c' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'b') ADVANCE(135);
      END_STATE();
    case 110:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          lookahead == 'a' ||
          lookahead == 'b' ||
          ('d' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'c') ADVANCE(186);
      END_STATE();
    case 111:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          lookahead == 'a' ||
          lookahead == 'b' ||
          ('d' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'c') ADVANCE(126);
      END_STATE();
    case 112:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'c') ||
          ('e' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'd') ADVANCE(95);
      END_STATE();
    case 113:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'c') ||
          ('e' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'd') ADVANCE(232);
      END_STATE();
    case 114:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'c') ||
          ('e' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'd') ADVANCE(163);
      END_STATE();
    case 115:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'c') ||
          ('e' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'd') ADVANCE(98);
      END_STATE();
    case 116:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'd') ||
          ('f' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'e') ADVANCE(133);
      END_STATE();
    case 117:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'd') ||
          ('f' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'e') ADVANCE(202);
      END_STATE();
    case 118:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'd') ||
          ('f' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'e') ADVANCE(242);
      END_STATE();
    case 119:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'd') ||
          ('f' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'e') ADVANCE(211);
      END_STATE();
    case 120:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'd') ||
          ('f' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'e') ADVANCE(243);
      END_STATE();
    case 121:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'd') ||
          ('f' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'e') ADVANCE(197);
      END_STATE();
    case 122:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'd') ||
          ('f' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'e') ADVANCE(180);
      END_STATE();
    case 123:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'd') ||
          ('f' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'e') ADVANCE(96);
      END_STATE();
    case 124:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'd') ||
          ('f' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'e') ADVANCE(153);
      END_STATE();
    case 125:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'd') ||
          ('f' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'e') ADVANCE(165);
      END_STATE();
    case 126:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'd') ||
          ('f' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'e') ADVANCE(160);
      END_STATE();
    case 127:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'd') ||
          ('f' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'e') ADVANCE(99);
      END_STATE();
    case 128:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'f') ||
          ('h' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'g') ADVANCE(121);
      END_STATE();
    case 129:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'g') ||
          ('i' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'h') ADVANCE(203);
      END_STATE();
    case 130:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'h') ||
          ('j' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'i') ADVANCE(177);
      END_STATE();
    case 131:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'h') ||
          ('j' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'i') ADVANCE(110);
      END_STATE();
    case 132:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'h') ||
          ('j' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'i') ADVANCE(114);
      END_STATE();
    case 133:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'k') ||
          ('m' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'l') ADVANCE(179);
      END_STATE();
    case 134:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'k') ||
          ('m' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'l') ADVANCE(161);
      END_STATE();
    case 135:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'k') ||
          ('m' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'l') ADVANCE(127);
      END_STATE();
    case 136:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'l') ||
          ('n' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'm') ADVANCE(146);
      END_STATE();
    case 137:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'l') ||
          ('n' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'm') ADVANCE(131);
      END_STATE();
    case 138:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'm') ||
          ('o' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'n') ADVANCE(94);
      END_STATE();
    case 139:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'm') ||
          ('o' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'n') ADVANCE(128);
      END_STATE();
    case 140:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'm') ||
          ('o' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'n') ADVANCE(178);
      END_STATE();
    case 141:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'm') ||
          ('o' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'n') ADVANCE(159);
      END_STATE();
    case 142:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'n') ||
          ('p' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'o') ADVANCE(137);
      END_STATE();
    case 143:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'n') ||
          ('p' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'o') ADVANCE(141);
      END_STATE();
    case 144:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'n') ||
          ('p' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'o') ADVANCE(172);
      END_STATE();
    case 145:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'n') ||
          ('p' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'o') ADVANCE(152);
      END_STATE();
    case 146:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'o') ||
          ('q' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'p') ADVANCE(145);
      END_STATE();
    case 147:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'o') ||
          ('q' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'p') ADVANCE(119);
      END_STATE();
    case 148:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'o') ||
          ('q' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'p') ADVANCE(157);
      END_STATE();
    case 149:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'q') ||
          ('s' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'r') ADVANCE(173);
      END_STATE();
    case 150:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'q') ||
          ('s' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'r') ADVANCE(188);
      END_STATE();
    case 151:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'q') ||
          ('s' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'r') ADVANCE(144);
      END_STATE();
    case 152:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'q') ||
          ('s' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'r') ADVANCE(166);
      END_STATE();
    case 153:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'q') ||
          ('s' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'r') ADVANCE(158);
      END_STATE();
    case 154:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'r') ||
          ('t' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 's') ADVANCE(225);
      END_STATE();
    case 155:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'r') ||
          ('t' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 's') ADVANCE(227);
      END_STATE();
    case 156:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'r') ||
          ('t' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 's') ADVANCE(184);
      END_STATE();
    case 157:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'r') ||
          ('t' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 's') ADVANCE(192);
      END_STATE();
    case 158:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'r') ||
          ('t' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 's') ADVANCE(195);
      END_STATE();
    case 159:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'r') ||
          ('t' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 's') ADVANCE(164);
      END_STATE();
    case 160:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'r') ||
          ('t' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 's') ADVANCE(156);
      END_STATE();
    case 161:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'r') ||
          ('t' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 's') ADVANCE(120);
      END_STATE();
    case 162:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'r') ||
          ('t' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 's') ADVANCE(170);
      END_STATE();
    case 163:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 's') ||
          ('u' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 't') ADVANCE(129);
      END_STATE();
    case 164:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 's') ||
          ('u' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 't') ADVANCE(206);
      END_STATE();
    case 165:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 's') ||
          ('u' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 't') ADVANCE(200);
      END_STATE();
    case 166:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 's') ||
          ('u' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 't') ADVANCE(205);
      END_STATE();
    case 167:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 's') ||
          ('u' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 't') ADVANCE(100);
      END_STATE();
    case 168:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 's') ||
          ('u' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 't') ADVANCE(123);
      END_STATE();
    case 169:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 's') ||
          ('u' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 't') ADVANCE(97);
      END_STATE();
    case 170:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 's') ||
          ('u' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 't') ADVANCE(124);
      END_STATE();
    case 171:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 't') ||
          ('v' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'u') ADVANCE(167);
      END_STATE();
    case 172:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 't') ||
          ('v' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'u') ADVANCE(148);
      END_STATE();
    case 173:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 't') ||
          ('v' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'u') ADVANCE(118);
      END_STATE();
    case 174:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'x') ||
          lookahead == 'z') ADVANCE(183);
      if (lookahead == 'y') ADVANCE(168);
      END_STATE();
    case 175:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'x') ||
          lookahead == 'z') ADVANCE(183);
      if (lookahead == 'y') ADVANCE(189);
      END_STATE();
    case 176:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'x') ||
          lookahead == 'z') ADVANCE(183);
      if (lookahead == 'y') ADVANCE(147);
      END_STATE();
    case 177:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'y')) ADVANCE(183);
      if (lookahead == 'z') ADVANCE(117);
      END_STATE();
    case 178:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'a') ADVANCE(109);
      END_STATE();
    case 179:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'a') ADVANCE(175);
      END_STATE();
    case 180:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'a') ADVANCE(150);
      END_STATE();
    case 181:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'a') ADVANCE(134);
      END_STATE();
    case 182:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'a') ADVANCE(162);
      END_STATE();
    case 183:
      ACCEPT_TOKEN(aux_sym_identifier_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 184:
      ACCEPT_TOKEN(anon_sym_access);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 185:
      ACCEPT_TOKEN(anon_sym_add_DASHenable);
      END_STATE();
    case 186:
      ACCEPT_TOKEN(anon_sym_atomic);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 187:
      ACCEPT_TOKEN(anon_sym_byte_DASHwrite_DASHenable);
      END_STATE();
    case 188:
      ACCEPT_TOKEN(anon_sym_clear);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 189:
      ACCEPT_TOKEN(anon_sym_delay);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 190:
      ACCEPT_TOKEN(anon_sym_enable_DASHinit_DASHvalue);
      END_STATE();
    case 191:
      ACCEPT_TOKEN(anon_sym_enable_DASHreset_DASHvalue);
      END_STATE();
    case 192:
      ACCEPT_TOKEN(anon_sym_groups);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 193:
      ACCEPT_TOKEN(anon_sym_init_DASHvalue);
      END_STATE();
    case 194:
      ACCEPT_TOKEN(anon_sym_in_DASHtrigger);
      END_STATE();
    case 195:
      ACCEPT_TOKEN(anon_sym_masters);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 196:
      ACCEPT_TOKEN(anon_sym_out_DASHtrigger);
      END_STATE();
    case 197:
      ACCEPT_TOKEN(anon_sym_range);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 198:
      ACCEPT_TOKEN(anon_sym_read_DASHlatency);
      END_STATE();
    case 199:
      ACCEPT_TOKEN(anon_sym_read_DASHvalue);
      END_STATE();
    case 200:
      ACCEPT_TOKEN(anon_sym_reset);
      if (lookahead == '-') ADVANCE(79);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 201:
      ACCEPT_TOKEN(anon_sym_reset_DASHvalue);
      END_STATE();
    case 202:
      ACCEPT_TOKEN(anon_sym_size);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 203:
      ACCEPT_TOKEN(anon_sym_width);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 204:
      ACCEPT_TOKEN(anon_sym_DOT);
      END_STATE();
    case 205:
      ACCEPT_TOKEN(anon_sym_import);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 206:
      ACCEPT_TOKEN(anon_sym_const);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 207:
      ACCEPT_TOKEN(anon_sym_EQ);
      END_STATE();
    case 208:
      ACCEPT_TOKEN(anon_sym_COMMA);
      END_STATE();
    case 209:
      ACCEPT_TOKEN(anon_sym_LPAREN);
      END_STATE();
    case 210:
      ACCEPT_TOKEN(anon_sym_RPAREN);
      END_STATE();
    case 211:
      ACCEPT_TOKEN(anon_sym_type);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 212:
      ACCEPT_TOKEN(anon_sym_SEMI);
      END_STATE();
    case 213:
      ACCEPT_TOKEN(anon_sym_LBRACK);
      END_STATE();
    case 214:
      ACCEPT_TOKEN(anon_sym_RBRACK);
      END_STATE();
    case 215:
      ACCEPT_TOKEN(anon_sym_DQUOTE);
      END_STATE();
    case 216:
      ACCEPT_TOKEN(aux_sym_string_literal_token1);
      if (lookahead == '#') ADVANCE(89);
      if (lookahead == '\t' ||
          lookahead == '\n' ||
          lookahead == '\r' ||
          lookahead == ' ') ADVANCE(216);
      if (lookahead != 0 &&
          lookahead != '"') ADVANCE(217);
      END_STATE();
    case 217:
      ACCEPT_TOKEN(aux_sym_string_literal_token1);
      if (lookahead != 0 &&
          lookahead != '"') ADVANCE(217);
      END_STATE();
    case 218:
      ACCEPT_TOKEN(sym_zero_literal);
      if (lookahead == 'B' ||
          lookahead == 'b') ADVANCE(84);
      if (lookahead == 'O' ||
          lookahead == 'o') ADVANCE(85);
      if (lookahead == 'X' ||
          lookahead == 'x') ADVANCE(86);
      END_STATE();
    case 219:
      ACCEPT_TOKEN(sym_decimal_literal);
      if (('0' <= lookahead && lookahead <= '9') ||
          lookahead == '_') ADVANCE(219);
      END_STATE();
    case 220:
      ACCEPT_TOKEN(sym_binary_literal);
      if (lookahead == '0' ||
          lookahead == '1' ||
          lookahead == '_') ADVANCE(220);
      END_STATE();
    case 221:
      ACCEPT_TOKEN(sym_octal_literal);
      if (('0' <= lookahead && lookahead <= '7') ||
          lookahead == '_') ADVANCE(221);
      END_STATE();
    case 222:
      ACCEPT_TOKEN(sym_hex_literal);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'f') ||
          lookahead == '|') ADVANCE(222);
      END_STATE();
    case 223:
      ACCEPT_TOKEN(sym_bit_literal);
      END_STATE();
    case 224:
      ACCEPT_TOKEN(anon_sym_ns);
      END_STATE();
    case 225:
      ACCEPT_TOKEN(anon_sym_ns);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 226:
      ACCEPT_TOKEN(anon_sym_us);
      END_STATE();
    case 227:
      ACCEPT_TOKEN(anon_sym_us);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 228:
      ACCEPT_TOKEN(anon_sym_ms);
      END_STATE();
    case 229:
      ACCEPT_TOKEN(anon_sym_ms);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 230:
      ACCEPT_TOKEN(anon_sym_s);
      END_STATE();
    case 231:
      ACCEPT_TOKEN(anon_sym_s);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'h') ||
          ('j' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      if (lookahead == 'i') ADVANCE(177);
      END_STATE();
    case 232:
      ACCEPT_TOKEN(anon_sym_and);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 233:
      ACCEPT_TOKEN(anon_sym_or);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 234:
      ACCEPT_TOKEN(anon_sym_PLUS);
      END_STATE();
    case 235:
      ACCEPT_TOKEN(anon_sym_DASH);
      END_STATE();
    case 236:
      ACCEPT_TOKEN(anon_sym_STAR);
      if (lookahead == '*') ADVANCE(239);
      END_STATE();
    case 237:
      ACCEPT_TOKEN(anon_sym_SLASH);
      END_STATE();
    case 238:
      ACCEPT_TOKEN(anon_sym_PERCENT);
      END_STATE();
    case 239:
      ACCEPT_TOKEN(anon_sym_STAR_STAR);
      END_STATE();
    case 240:
      ACCEPT_TOKEN(anon_sym_LT_LT);
      END_STATE();
    case 241:
      ACCEPT_TOKEN(anon_sym_GT_GT);
      END_STATE();
    case 242:
      ACCEPT_TOKEN(anon_sym_true);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    case 243:
      ACCEPT_TOKEN(anon_sym_false);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(183);
      END_STATE();
    default:
      return false;
  }
}

static const TSLexMode ts_lex_modes[STATE_COUNT] = {
  [0] = {.lex_state = 0, .external_lex_state = 1},
  [1] = {.lex_state = 87},
  [2] = {.lex_state = 5, .external_lex_state = 2},
  [3] = {.lex_state = 5, .external_lex_state = 2},
  [4] = {.lex_state = 5, .external_lex_state = 2},
  [5] = {.lex_state = 5},
  [6] = {.lex_state = 5},
  [7] = {.lex_state = 1},
  [8] = {.lex_state = 1},
  [9] = {.lex_state = 1},
  [10] = {.lex_state = 1},
  [11] = {.lex_state = 1},
  [12] = {.lex_state = 1},
  [13] = {.lex_state = 1},
  [14] = {.lex_state = 1},
  [15] = {.lex_state = 1},
  [16] = {.lex_state = 1},
  [17] = {.lex_state = 1},
  [18] = {.lex_state = 1},
  [19] = {.lex_state = 1},
  [20] = {.lex_state = 1},
  [21] = {.lex_state = 1},
  [22] = {.lex_state = 1},
  [23] = {.lex_state = 1},
  [24] = {.lex_state = 1},
  [25] = {.lex_state = 1},
  [26] = {.lex_state = 1},
  [27] = {.lex_state = 1},
  [28] = {.lex_state = 1},
  [29] = {.lex_state = 1},
  [30] = {.lex_state = 1},
  [31] = {.lex_state = 1},
  [32] = {.lex_state = 1},
  [33] = {.lex_state = 1},
  [34] = {.lex_state = 1},
  [35] = {.lex_state = 5, .external_lex_state = 3},
  [36] = {.lex_state = 5, .external_lex_state = 3},
  [37] = {.lex_state = 5, .external_lex_state = 3},
  [38] = {.lex_state = 5, .external_lex_state = 3},
  [39] = {.lex_state = 5, .external_lex_state = 3},
  [40] = {.lex_state = 5, .external_lex_state = 3},
  [41] = {.lex_state = 5, .external_lex_state = 2},
  [42] = {.lex_state = 5, .external_lex_state = 2},
  [43] = {.lex_state = 5, .external_lex_state = 2},
  [44] = {.lex_state = 5, .external_lex_state = 2},
  [45] = {.lex_state = 5, .external_lex_state = 2},
  [46] = {.lex_state = 5, .external_lex_state = 2},
  [47] = {.lex_state = 5, .external_lex_state = 2},
  [48] = {.lex_state = 5, .external_lex_state = 2},
  [49] = {.lex_state = 5, .external_lex_state = 2},
  [50] = {.lex_state = 5, .external_lex_state = 2},
  [51] = {.lex_state = 5, .external_lex_state = 2},
  [52] = {.lex_state = 5, .external_lex_state = 2},
  [53] = {.lex_state = 5, .external_lex_state = 2},
  [54] = {.lex_state = 5, .external_lex_state = 2},
  [55] = {.lex_state = 5, .external_lex_state = 2},
  [56] = {.lex_state = 5, .external_lex_state = 2},
  [57] = {.lex_state = 5, .external_lex_state = 2},
  [58] = {.lex_state = 5},
  [59] = {.lex_state = 5},
  [60] = {.lex_state = 5},
  [61] = {.lex_state = 5},
  [62] = {.lex_state = 5},
  [63] = {.lex_state = 5},
  [64] = {.lex_state = 5},
  [65] = {.lex_state = 5},
  [66] = {.lex_state = 5},
  [67] = {.lex_state = 5},
  [68] = {.lex_state = 5},
  [69] = {.lex_state = 5},
  [70] = {.lex_state = 5},
  [71] = {.lex_state = 87},
  [72] = {.lex_state = 87},
  [73] = {.lex_state = 4},
  [74] = {.lex_state = 4},
  [75] = {.lex_state = 0},
  [76] = {.lex_state = 4, .external_lex_state = 4},
  [77] = {.lex_state = 4, .external_lex_state = 4},
  [78] = {.lex_state = 0},
  [79] = {.lex_state = 0},
  [80] = {.lex_state = 0},
  [81] = {.lex_state = 0},
  [82] = {.lex_state = 0, .external_lex_state = 4},
  [83] = {.lex_state = 0, .external_lex_state = 4},
  [84] = {.lex_state = 0, .external_lex_state = 4},
  [85] = {.lex_state = 0, .external_lex_state = 4},
  [86] = {.lex_state = 0},
  [87] = {.lex_state = 0, .external_lex_state = 4},
  [88] = {.lex_state = 0},
  [89] = {.lex_state = 0},
  [90] = {.lex_state = 0},
  [91] = {.lex_state = 0},
  [92] = {.lex_state = 0},
  [93] = {.lex_state = 0},
  [94] = {.lex_state = 0},
  [95] = {.lex_state = 0},
  [96] = {.lex_state = 0},
  [97] = {.lex_state = 0},
  [98] = {.lex_state = 0},
  [99] = {.lex_state = 0},
  [100] = {.lex_state = 0},
  [101] = {.lex_state = 0},
  [102] = {.lex_state = 0, .external_lex_state = 4},
  [103] = {.lex_state = 0},
  [104] = {.lex_state = 0},
  [105] = {.lex_state = 0},
  [106] = {.lex_state = 0, .external_lex_state = 4},
  [107] = {.lex_state = 0, .external_lex_state = 4},
  [108] = {.lex_state = 0, .external_lex_state = 4},
  [109] = {.lex_state = 0, .external_lex_state = 4},
  [110] = {.lex_state = 0, .external_lex_state = 4},
  [111] = {.lex_state = 0, .external_lex_state = 4},
  [112] = {.lex_state = 0, .external_lex_state = 4},
  [113] = {.lex_state = 0, .external_lex_state = 4},
  [114] = {.lex_state = 0, .external_lex_state = 4},
  [115] = {.lex_state = 0, .external_lex_state = 4},
  [116] = {.lex_state = 0, .external_lex_state = 4},
  [117] = {.lex_state = 0, .external_lex_state = 4},
  [118] = {.lex_state = 0, .external_lex_state = 4},
  [119] = {.lex_state = 0, .external_lex_state = 4},
  [120] = {.lex_state = 0, .external_lex_state = 4},
  [121] = {.lex_state = 0, .external_lex_state = 4},
  [122] = {.lex_state = 2},
  [123] = {.lex_state = 87, .external_lex_state = 5},
  [124] = {.lex_state = 0, .external_lex_state = 4},
  [125] = {.lex_state = 87, .external_lex_state = 5},
  [126] = {.lex_state = 0, .external_lex_state = 4},
  [127] = {.lex_state = 0, .external_lex_state = 4},
  [128] = {.lex_state = 0, .external_lex_state = 4},
  [129] = {.lex_state = 0, .external_lex_state = 4},
  [130] = {.lex_state = 0, .external_lex_state = 4},
  [131] = {.lex_state = 0, .external_lex_state = 4},
  [132] = {.lex_state = 2, .external_lex_state = 2},
  [133] = {.lex_state = 87, .external_lex_state = 5},
  [134] = {.lex_state = 87, .external_lex_state = 5},
  [135] = {.lex_state = 2, .external_lex_state = 2},
  [136] = {.lex_state = 87, .external_lex_state = 5},
  [137] = {.lex_state = 0, .external_lex_state = 4},
  [138] = {.lex_state = 87, .external_lex_state = 5},
  [139] = {.lex_state = 87},
  [140] = {.lex_state = 87},
  [141] = {.lex_state = 87},
  [142] = {.lex_state = 2},
  [143] = {.lex_state = 2},
  [144] = {.lex_state = 2},
  [145] = {.lex_state = 0, .external_lex_state = 4},
  [146] = {.lex_state = 2},
  [147] = {.lex_state = 87},
  [148] = {.lex_state = 87},
  [149] = {.lex_state = 87},
  [150] = {.lex_state = 87},
  [151] = {.lex_state = 0, .external_lex_state = 4},
  [152] = {.lex_state = 87},
  [153] = {.lex_state = 0, .external_lex_state = 4},
  [154] = {.lex_state = 0, .external_lex_state = 4},
  [155] = {.lex_state = 87},
  [156] = {.lex_state = 2, .external_lex_state = 4},
  [157] = {.lex_state = 2},
  [158] = {.lex_state = 0, .external_lex_state = 4},
  [159] = {.lex_state = 0, .external_lex_state = 4},
  [160] = {.lex_state = 87},
  [161] = {.lex_state = 2},
  [162] = {.lex_state = 87},
  [163] = {.lex_state = 0, .external_lex_state = 4},
  [164] = {.lex_state = 87},
  [165] = {.lex_state = 87},
  [166] = {.lex_state = 87},
  [167] = {.lex_state = 0, .external_lex_state = 4},
  [168] = {.lex_state = 87},
  [169] = {.lex_state = 87},
  [170] = {.lex_state = 87},
  [171] = {.lex_state = 87},
  [172] = {.lex_state = 87},
  [173] = {.lex_state = 87},
  [174] = {.lex_state = 2, .external_lex_state = 2},
  [175] = {.lex_state = 0},
  [176] = {.lex_state = 2, .external_lex_state = 2},
  [177] = {.lex_state = 0},
  [178] = {.lex_state = 0},
  [179] = {.lex_state = 2, .external_lex_state = 2},
  [180] = {.lex_state = 0},
  [181] = {.lex_state = 2},
  [182] = {.lex_state = 2, .external_lex_state = 4},
  [183] = {.lex_state = 0},
  [184] = {.lex_state = 2},
  [185] = {.lex_state = 2, .external_lex_state = 4},
  [186] = {.lex_state = 2, .external_lex_state = 2},
  [187] = {.lex_state = 2},
  [188] = {.lex_state = 2},
  [189] = {.lex_state = 0, .external_lex_state = 4},
  [190] = {.lex_state = 0, .external_lex_state = 4},
  [191] = {.lex_state = 2},
  [192] = {.lex_state = 0, .external_lex_state = 4},
  [193] = {.lex_state = 0, .external_lex_state = 4},
  [194] = {.lex_state = 0},
  [195] = {.lex_state = 0, .external_lex_state = 4},
  [196] = {.lex_state = 0, .external_lex_state = 4},
  [197] = {.lex_state = 0},
  [198] = {.lex_state = 0},
  [199] = {.lex_state = 0},
  [200] = {.lex_state = 0, .external_lex_state = 4},
  [201] = {.lex_state = 0, .external_lex_state = 4},
  [202] = {.lex_state = 0},
  [203] = {.lex_state = 0},
  [204] = {.lex_state = 0},
  [205] = {.lex_state = 2},
  [206] = {.lex_state = 0, .external_lex_state = 4},
  [207] = {.lex_state = 2, .external_lex_state = 2},
  [208] = {.lex_state = 0},
  [209] = {.lex_state = 0, .external_lex_state = 4},
  [210] = {.lex_state = 0, .external_lex_state = 4},
  [211] = {.lex_state = 0, .external_lex_state = 4},
  [212] = {.lex_state = 2},
  [213] = {.lex_state = 0, .external_lex_state = 4},
  [214] = {.lex_state = 0, .external_lex_state = 4},
  [215] = {.lex_state = 0, .external_lex_state = 4},
  [216] = {.lex_state = 0},
  [217] = {.lex_state = 2},
  [218] = {.lex_state = 2},
  [219] = {.lex_state = 2},
  [220] = {.lex_state = 0, .external_lex_state = 4},
  [221] = {.lex_state = 2},
  [222] = {.lex_state = 2, .external_lex_state = 2},
  [223] = {.lex_state = 0, .external_lex_state = 4},
  [224] = {.lex_state = 0},
  [225] = {.lex_state = 0},
  [226] = {.lex_state = 0},
  [227] = {.lex_state = 0, .external_lex_state = 4},
  [228] = {.lex_state = 0},
  [229] = {.lex_state = 0},
  [230] = {.lex_state = 0, .external_lex_state = 4},
  [231] = {.lex_state = 0},
  [232] = {.lex_state = 0},
  [233] = {.lex_state = 0},
  [234] = {.lex_state = 0, .external_lex_state = 4},
  [235] = {.lex_state = 0, .external_lex_state = 4},
  [236] = {.lex_state = 2},
  [237] = {.lex_state = 0},
  [238] = {.lex_state = 0},
  [239] = {.lex_state = 0},
  [240] = {.lex_state = 0},
  [241] = {.lex_state = 0},
  [242] = {.lex_state = 0},
  [243] = {.lex_state = 0, .external_lex_state = 5},
  [244] = {.lex_state = 2},
  [245] = {.lex_state = 0, .external_lex_state = 4},
  [246] = {.lex_state = 0},
  [247] = {.lex_state = 0},
  [248] = {.lex_state = 0},
  [249] = {.lex_state = 0, .external_lex_state = 4},
  [250] = {.lex_state = 0},
  [251] = {.lex_state = 216},
  [252] = {.lex_state = 0, .external_lex_state = 5},
  [253] = {.lex_state = 2},
  [254] = {.lex_state = 216},
  [255] = {.lex_state = 0, .external_lex_state = 4},
  [256] = {.lex_state = 0},
  [257] = {.lex_state = 0},
  [258] = {.lex_state = 0, .external_lex_state = 4},
  [259] = {.lex_state = 0, .external_lex_state = 4},
  [260] = {.lex_state = 0},
  [261] = {.lex_state = 0},
  [262] = {.lex_state = 0},
  [263] = {.lex_state = 0},
  [264] = {.lex_state = 0, .external_lex_state = 5},
  [265] = {.lex_state = 0},
  [266] = {.lex_state = 0},
};

enum {
  ts_external_token__indent = 0,
  ts_external_token__dedent = 1,
  ts_external_token__newline = 2,
};

static const TSSymbol ts_external_scanner_symbol_map[EXTERNAL_TOKEN_COUNT] = {
  [ts_external_token__indent] = sym__indent,
  [ts_external_token__dedent] = sym__dedent,
  [ts_external_token__newline] = sym__newline,
};

static const bool ts_external_scanner_states[6][EXTERNAL_TOKEN_COUNT] = {
  [1] = {
    [ts_external_token__indent] = true,
    [ts_external_token__dedent] = true,
    [ts_external_token__newline] = true,
  },
  [2] = {
    [ts_external_token__dedent] = true,
  },
  [3] = {
    [ts_external_token__indent] = true,
    [ts_external_token__dedent] = true,
  },
  [4] = {
    [ts_external_token__newline] = true,
  },
  [5] = {
    [ts_external_token__indent] = true,
  },
};

static const uint16_t ts_parse_table[LARGE_STATE_COUNT][SYMBOL_COUNT] = {
  [0] = {
    [ts_builtin_sym_end] = ACTIONS(1),
    [sym_comment] = ACTIONS(3),
    [aux_sym_identifier_token1] = ACTIONS(1),
    [anon_sym_access] = ACTIONS(1),
    [anon_sym_add_DASHenable] = ACTIONS(1),
    [anon_sym_atomic] = ACTIONS(1),
    [anon_sym_byte_DASHwrite_DASHenable] = ACTIONS(1),
    [anon_sym_clear] = ACTIONS(1),
    [anon_sym_delay] = ACTIONS(1),
    [anon_sym_enable_DASHinit_DASHvalue] = ACTIONS(1),
    [anon_sym_enable_DASHreset_DASHvalue] = ACTIONS(1),
    [anon_sym_groups] = ACTIONS(1),
    [anon_sym_init_DASHvalue] = ACTIONS(1),
    [anon_sym_in_DASHtrigger] = ACTIONS(1),
    [anon_sym_masters] = ACTIONS(1),
    [anon_sym_out_DASHtrigger] = ACTIONS(1),
    [anon_sym_range] = ACTIONS(1),
    [anon_sym_read_DASHlatency] = ACTIONS(1),
    [anon_sym_read_DASHvalue] = ACTIONS(1),
    [anon_sym_reset] = ACTIONS(1),
    [anon_sym_reset_DASHvalue] = ACTIONS(1),
    [anon_sym_size] = ACTIONS(1),
    [anon_sym_width] = ACTIONS(1),
    [anon_sym_DOT] = ACTIONS(1),
    [anon_sym_import] = ACTIONS(1),
    [anon_sym_const] = ACTIONS(1),
    [anon_sym_EQ] = ACTIONS(1),
    [anon_sym_COMMA] = ACTIONS(1),
    [anon_sym_LPAREN] = ACTIONS(1),
    [anon_sym_RPAREN] = ACTIONS(1),
    [anon_sym_type] = ACTIONS(1),
    [anon_sym_SEMI] = ACTIONS(1),
    [anon_sym_LBRACK] = ACTIONS(1),
    [anon_sym_RBRACK] = ACTIONS(1),
    [anon_sym_DQUOTE] = ACTIONS(1),
    [sym_zero_literal] = ACTIONS(1),
    [sym_decimal_literal] = ACTIONS(1),
    [sym_binary_literal] = ACTIONS(1),
    [sym_octal_literal] = ACTIONS(1),
    [sym_hex_literal] = ACTIONS(1),
    [sym_bit_literal] = ACTIONS(1),
    [anon_sym_ns] = ACTIONS(1),
    [anon_sym_us] = ACTIONS(1),
    [anon_sym_ms] = ACTIONS(1),
    [anon_sym_s] = ACTIONS(1),
    [anon_sym_and] = ACTIONS(1),
    [anon_sym_or] = ACTIONS(1),
    [anon_sym_PLUS] = ACTIONS(1),
    [anon_sym_DASH] = ACTIONS(1),
    [anon_sym_STAR] = ACTIONS(1),
    [anon_sym_SLASH] = ACTIONS(1),
    [anon_sym_PERCENT] = ACTIONS(1),
    [anon_sym_STAR_STAR] = ACTIONS(1),
    [anon_sym_LT_LT] = ACTIONS(1),
    [anon_sym_GT_GT] = ACTIONS(1),
    [anon_sym_true] = ACTIONS(1),
    [anon_sym_false] = ACTIONS(1),
    [sym__indent] = ACTIONS(1),
    [sym__dedent] = ACTIONS(1),
    [sym__newline] = ACTIONS(1),
  },
  [1] = {
    [sym_description] = STATE(260),
    [sym_identifier] = STATE(142),
    [sym__meta_statement] = STATE(71),
    [sym__import_statement] = STATE(71),
    [sym_single_import_statement] = STATE(71),
    [sym_multi_import_statement] = STATE(71),
    [sym__constant_definition] = STATE(71),
    [sym_single_constant_definition] = STATE(71),
    [sym_multi_constant_definition] = STATE(71),
    [sym__type_definition] = STATE(71),
    [sym_single_line_type_definition] = STATE(71),
    [sym_multi_line_type_definition] = STATE(71),
    [sym__instantiation] = STATE(71),
    [sym_single_line_instantiation] = STATE(71),
    [sym_multi_line_instantiation] = STATE(71),
    [aux_sym_description_repeat1] = STATE(71),
    [ts_builtin_sym_end] = ACTIONS(5),
    [sym_comment] = ACTIONS(3),
    [aux_sym_identifier_token1] = ACTIONS(7),
    [anon_sym_import] = ACTIONS(9),
    [anon_sym_const] = ACTIONS(11),
    [anon_sym_type] = ACTIONS(13),
  },
};

static const uint16_t ts_small_parse_table[] = {
  [0] = 11,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(7), 1,
      aux_sym_identifier_token1,
    ACTIONS(19), 1,
      anon_sym_const,
    ACTIONS(21), 1,
      anon_sym_type,
    ACTIONS(23), 1,
      sym__dedent,
    STATE(144), 1,
      sym_identifier,
    STATE(196), 1,
      sym__property_assignment,
    STATE(257), 1,
      sym_property_identifier,
    ACTIONS(15), 10,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
    ACTIONS(17), 10,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    STATE(4), 11,
      sym__constant_definition,
      sym_single_constant_definition,
      sym_multi_constant_definition,
      sym__type_definition,
      sym_single_line_type_definition,
      sym_multi_line_type_definition,
      sym_property_assignments,
      sym__instantiation,
      sym_single_line_instantiation,
      sym_multi_line_instantiation,
      aux_sym_element_body_repeat1,
  [62] = 11,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(7), 1,
      aux_sym_identifier_token1,
    ACTIONS(19), 1,
      anon_sym_const,
    ACTIONS(21), 1,
      anon_sym_type,
    ACTIONS(25), 1,
      sym__dedent,
    STATE(144), 1,
      sym_identifier,
    STATE(196), 1,
      sym__property_assignment,
    STATE(257), 1,
      sym_property_identifier,
    ACTIONS(15), 10,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
    ACTIONS(17), 10,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    STATE(4), 11,
      sym__constant_definition,
      sym_single_constant_definition,
      sym_multi_constant_definition,
      sym__type_definition,
      sym_single_line_type_definition,
      sym_multi_line_type_definition,
      sym_property_assignments,
      sym__instantiation,
      sym_single_line_instantiation,
      sym_multi_line_instantiation,
      aux_sym_element_body_repeat1,
  [124] = 11,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(27), 1,
      aux_sym_identifier_token1,
    ACTIONS(36), 1,
      anon_sym_const,
    ACTIONS(39), 1,
      anon_sym_type,
    ACTIONS(42), 1,
      sym__dedent,
    STATE(144), 1,
      sym_identifier,
    STATE(196), 1,
      sym__property_assignment,
    STATE(257), 1,
      sym_property_identifier,
    ACTIONS(30), 10,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
    ACTIONS(33), 10,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    STATE(4), 11,
      sym__constant_definition,
      sym_single_constant_definition,
      sym_multi_constant_definition,
      sym__type_definition,
      sym_single_line_type_definition,
      sym_multi_line_type_definition,
      sym_property_assignments,
      sym__instantiation,
      sym_single_line_instantiation,
      sym_multi_line_instantiation,
      aux_sym_element_body_repeat1,
  [186] = 10,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(7), 1,
      aux_sym_identifier_token1,
    ACTIONS(19), 1,
      anon_sym_const,
    ACTIONS(21), 1,
      anon_sym_type,
    STATE(144), 1,
      sym_identifier,
    STATE(196), 1,
      sym__property_assignment,
    STATE(257), 1,
      sym_property_identifier,
    ACTIONS(15), 10,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
    ACTIONS(17), 10,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    STATE(3), 11,
      sym__constant_definition,
      sym_single_constant_definition,
      sym_multi_constant_definition,
      sym__type_definition,
      sym_single_line_type_definition,
      sym_multi_line_type_definition,
      sym_property_assignments,
      sym__instantiation,
      sym_single_line_instantiation,
      sym_multi_line_instantiation,
      aux_sym_element_body_repeat1,
  [245] = 10,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(7), 1,
      aux_sym_identifier_token1,
    ACTIONS(19), 1,
      anon_sym_const,
    ACTIONS(21), 1,
      anon_sym_type,
    STATE(144), 1,
      sym_identifier,
    STATE(196), 1,
      sym__property_assignment,
    STATE(257), 1,
      sym_property_identifier,
    ACTIONS(15), 10,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
    ACTIONS(17), 10,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    STATE(2), 11,
      sym__constant_definition,
      sym_single_constant_definition,
      sym_multi_constant_definition,
      sym__type_definition,
      sym_single_line_type_definition,
      sym_multi_line_type_definition,
      sym_property_assignments,
      sym__instantiation,
      sym_single_line_instantiation,
      sym_multi_line_instantiation,
      aux_sym_element_body_repeat1,
  [304] = 21,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(44), 1,
      aux_sym_identifier_token1,
    ACTIONS(46), 1,
      anon_sym_DOT,
    ACTIONS(48), 1,
      anon_sym_LPAREN,
    ACTIONS(50), 1,
      anon_sym_RPAREN,
    ACTIONS(52), 1,
      anon_sym_LBRACK,
    ACTIONS(54), 1,
      anon_sym_DQUOTE,
    ACTIONS(56), 1,
      sym_zero_literal,
    ACTIONS(58), 1,
      sym_decimal_literal,
    ACTIONS(62), 1,
      sym_bit_literal,
    STATE(73), 1,
      sym__integer_literal,
    STATE(80), 1,
      sym_declared_identifier,
    STATE(86), 1,
      sym_qualified_identifier,
    STATE(92), 1,
      sym_primary_expression,
    STATE(250), 1,
      sym__arguments,
    STATE(262), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(64), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(66), 2,
      anon_sym_true,
      anon_sym_false,
    ACTIONS(60), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(194), 3,
      sym__argument,
      sym_expression_list,
      sym__expression,
    STATE(100), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [381] = 21,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(44), 1,
      aux_sym_identifier_token1,
    ACTIONS(46), 1,
      anon_sym_DOT,
    ACTIONS(48), 1,
      anon_sym_LPAREN,
    ACTIONS(52), 1,
      anon_sym_LBRACK,
    ACTIONS(54), 1,
      anon_sym_DQUOTE,
    ACTIONS(56), 1,
      sym_zero_literal,
    ACTIONS(58), 1,
      sym_decimal_literal,
    ACTIONS(62), 1,
      sym_bit_literal,
    ACTIONS(68), 1,
      anon_sym_RPAREN,
    STATE(73), 1,
      sym__integer_literal,
    STATE(81), 1,
      sym_declared_identifier,
    STATE(86), 1,
      sym_qualified_identifier,
    STATE(92), 1,
      sym_primary_expression,
    STATE(247), 1,
      sym__expressions,
    STATE(262), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(64), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(66), 2,
      anon_sym_true,
      anon_sym_false,
    STATE(177), 2,
      sym_expression_list,
      sym__expression,
    ACTIONS(60), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(100), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [457] = 21,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(44), 1,
      aux_sym_identifier_token1,
    ACTIONS(46), 1,
      anon_sym_DOT,
    ACTIONS(48), 1,
      anon_sym_LPAREN,
    ACTIONS(52), 1,
      anon_sym_LBRACK,
    ACTIONS(54), 1,
      anon_sym_DQUOTE,
    ACTIONS(56), 1,
      sym_zero_literal,
    ACTIONS(58), 1,
      sym_decimal_literal,
    ACTIONS(62), 1,
      sym_bit_literal,
    ACTIONS(70), 1,
      anon_sym_RBRACK,
    STATE(73), 1,
      sym__integer_literal,
    STATE(81), 1,
      sym_declared_identifier,
    STATE(86), 1,
      sym_qualified_identifier,
    STATE(92), 1,
      sym_primary_expression,
    STATE(240), 1,
      sym__expressions,
    STATE(262), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(64), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(66), 2,
      anon_sym_true,
      anon_sym_false,
    STATE(177), 2,
      sym_expression_list,
      sym__expression,
    ACTIONS(60), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(100), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [533] = 21,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(44), 1,
      aux_sym_identifier_token1,
    ACTIONS(46), 1,
      anon_sym_DOT,
    ACTIONS(48), 1,
      anon_sym_LPAREN,
    ACTIONS(52), 1,
      anon_sym_LBRACK,
    ACTIONS(54), 1,
      anon_sym_DQUOTE,
    ACTIONS(56), 1,
      sym_zero_literal,
    ACTIONS(58), 1,
      sym_decimal_literal,
    ACTIONS(62), 1,
      sym_bit_literal,
    ACTIONS(72), 1,
      anon_sym_RPAREN,
    STATE(73), 1,
      sym__integer_literal,
    STATE(81), 1,
      sym_declared_identifier,
    STATE(86), 1,
      sym_qualified_identifier,
    STATE(92), 1,
      sym_primary_expression,
    STATE(246), 1,
      sym__expressions,
    STATE(262), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(64), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(66), 2,
      anon_sym_true,
      anon_sym_false,
    STATE(177), 2,
      sym_expression_list,
      sym__expression,
    ACTIONS(60), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(100), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [609] = 21,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(44), 1,
      aux_sym_identifier_token1,
    ACTIONS(46), 1,
      anon_sym_DOT,
    ACTIONS(48), 1,
      anon_sym_LPAREN,
    ACTIONS(52), 1,
      anon_sym_LBRACK,
    ACTIONS(54), 1,
      anon_sym_DQUOTE,
    ACTIONS(56), 1,
      sym_zero_literal,
    ACTIONS(58), 1,
      sym_decimal_literal,
    ACTIONS(62), 1,
      sym_bit_literal,
    ACTIONS(74), 1,
      anon_sym_RBRACK,
    STATE(73), 1,
      sym__integer_literal,
    STATE(81), 1,
      sym_declared_identifier,
    STATE(86), 1,
      sym_qualified_identifier,
    STATE(92), 1,
      sym_primary_expression,
    STATE(242), 1,
      sym__expressions,
    STATE(262), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(64), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(66), 2,
      anon_sym_true,
      anon_sym_false,
    STATE(177), 2,
      sym_expression_list,
      sym__expression,
    ACTIONS(60), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(100), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [685] = 19,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(44), 1,
      aux_sym_identifier_token1,
    ACTIONS(46), 1,
      anon_sym_DOT,
    ACTIONS(48), 1,
      anon_sym_LPAREN,
    ACTIONS(52), 1,
      anon_sym_LBRACK,
    ACTIONS(54), 1,
      anon_sym_DQUOTE,
    ACTIONS(56), 1,
      sym_zero_literal,
    ACTIONS(58), 1,
      sym_decimal_literal,
    ACTIONS(62), 1,
      sym_bit_literal,
    STATE(73), 1,
      sym__integer_literal,
    STATE(80), 1,
      sym_declared_identifier,
    STATE(86), 1,
      sym_qualified_identifier,
    STATE(92), 1,
      sym_primary_expression,
    STATE(262), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(64), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(66), 2,
      anon_sym_true,
      anon_sym_false,
    ACTIONS(60), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(229), 3,
      sym__argument,
      sym_expression_list,
      sym__expression,
    STATE(100), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [756] = 19,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(44), 1,
      aux_sym_identifier_token1,
    ACTIONS(46), 1,
      anon_sym_DOT,
    ACTIONS(48), 1,
      anon_sym_LPAREN,
    ACTIONS(52), 1,
      anon_sym_LBRACK,
    ACTIONS(54), 1,
      anon_sym_DQUOTE,
    ACTIONS(56), 1,
      sym_zero_literal,
    ACTIONS(58), 1,
      sym_decimal_literal,
    ACTIONS(62), 1,
      sym_bit_literal,
    STATE(73), 1,
      sym__integer_literal,
    STATE(81), 1,
      sym_declared_identifier,
    STATE(86), 1,
      sym_qualified_identifier,
    STATE(92), 1,
      sym_primary_expression,
    STATE(262), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(64), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(66), 2,
      anon_sym_true,
      anon_sym_false,
    STATE(199), 2,
      sym_expression_list,
      sym__expression,
    ACTIONS(60), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(100), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [826] = 19,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(76), 1,
      aux_sym_identifier_token1,
    ACTIONS(78), 1,
      anon_sym_DOT,
    ACTIONS(80), 1,
      anon_sym_LPAREN,
    ACTIONS(82), 1,
      anon_sym_LBRACK,
    ACTIONS(84), 1,
      anon_sym_DQUOTE,
    ACTIONS(86), 1,
      sym_zero_literal,
    ACTIONS(88), 1,
      sym_decimal_literal,
    ACTIONS(92), 1,
      sym_bit_literal,
    STATE(77), 1,
      sym__integer_literal,
    STATE(85), 1,
      sym_declared_identifier,
    STATE(102), 1,
      sym_qualified_identifier,
    STATE(109), 1,
      sym_primary_expression,
    STATE(248), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(94), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(96), 2,
      anon_sym_true,
      anon_sym_false,
    STATE(214), 2,
      sym_expression_list,
      sym__expression,
    ACTIONS(90), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(116), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [896] = 19,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(44), 1,
      aux_sym_identifier_token1,
    ACTIONS(46), 1,
      anon_sym_DOT,
    ACTIONS(48), 1,
      anon_sym_LPAREN,
    ACTIONS(52), 1,
      anon_sym_LBRACK,
    ACTIONS(54), 1,
      anon_sym_DQUOTE,
    ACTIONS(56), 1,
      sym_zero_literal,
    ACTIONS(58), 1,
      sym_decimal_literal,
    ACTIONS(62), 1,
      sym_bit_literal,
    STATE(73), 1,
      sym__integer_literal,
    STATE(81), 1,
      sym_declared_identifier,
    STATE(86), 1,
      sym_qualified_identifier,
    STATE(92), 1,
      sym_primary_expression,
    STATE(262), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(64), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(66), 2,
      anon_sym_true,
      anon_sym_false,
    STATE(224), 2,
      sym_expression_list,
      sym__expression,
    ACTIONS(60), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(100), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [966] = 19,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(44), 1,
      aux_sym_identifier_token1,
    ACTIONS(46), 1,
      anon_sym_DOT,
    ACTIONS(48), 1,
      anon_sym_LPAREN,
    ACTIONS(52), 1,
      anon_sym_LBRACK,
    ACTIONS(54), 1,
      anon_sym_DQUOTE,
    ACTIONS(56), 1,
      sym_zero_literal,
    ACTIONS(58), 1,
      sym_decimal_literal,
    ACTIONS(62), 1,
      sym_bit_literal,
    STATE(73), 1,
      sym__integer_literal,
    STATE(81), 1,
      sym_declared_identifier,
    STATE(86), 1,
      sym_qualified_identifier,
    STATE(92), 1,
      sym_primary_expression,
    STATE(262), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(64), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(66), 2,
      anon_sym_true,
      anon_sym_false,
    STATE(228), 2,
      sym_expression_list,
      sym__expression,
    ACTIONS(60), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(100), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [1036] = 19,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(44), 1,
      aux_sym_identifier_token1,
    ACTIONS(46), 1,
      anon_sym_DOT,
    ACTIONS(48), 1,
      anon_sym_LPAREN,
    ACTIONS(52), 1,
      anon_sym_LBRACK,
    ACTIONS(54), 1,
      anon_sym_DQUOTE,
    ACTIONS(56), 1,
      sym_zero_literal,
    ACTIONS(58), 1,
      sym_decimal_literal,
    ACTIONS(62), 1,
      sym_bit_literal,
    STATE(73), 1,
      sym__integer_literal,
    STATE(81), 1,
      sym_declared_identifier,
    STATE(86), 1,
      sym_qualified_identifier,
    STATE(92), 1,
      sym_primary_expression,
    STATE(262), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(64), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(66), 2,
      anon_sym_true,
      anon_sym_false,
    STATE(263), 2,
      sym_expression_list,
      sym__expression,
    ACTIONS(60), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(100), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [1106] = 19,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(76), 1,
      aux_sym_identifier_token1,
    ACTIONS(78), 1,
      anon_sym_DOT,
    ACTIONS(80), 1,
      anon_sym_LPAREN,
    ACTIONS(82), 1,
      anon_sym_LBRACK,
    ACTIONS(84), 1,
      anon_sym_DQUOTE,
    ACTIONS(86), 1,
      sym_zero_literal,
    ACTIONS(88), 1,
      sym_decimal_literal,
    ACTIONS(92), 1,
      sym_bit_literal,
    STATE(77), 1,
      sym__integer_literal,
    STATE(85), 1,
      sym_declared_identifier,
    STATE(102), 1,
      sym_qualified_identifier,
    STATE(109), 1,
      sym_primary_expression,
    STATE(248), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(94), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(96), 2,
      anon_sym_true,
      anon_sym_false,
    STATE(258), 2,
      sym_expression_list,
      sym__expression,
    ACTIONS(90), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(116), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [1176] = 19,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(44), 1,
      aux_sym_identifier_token1,
    ACTIONS(46), 1,
      anon_sym_DOT,
    ACTIONS(48), 1,
      anon_sym_LPAREN,
    ACTIONS(52), 1,
      anon_sym_LBRACK,
    ACTIONS(54), 1,
      anon_sym_DQUOTE,
    ACTIONS(56), 1,
      sym_zero_literal,
    ACTIONS(58), 1,
      sym_decimal_literal,
    ACTIONS(62), 1,
      sym_bit_literal,
    STATE(73), 1,
      sym__integer_literal,
    STATE(81), 1,
      sym_declared_identifier,
    STATE(86), 1,
      sym_qualified_identifier,
    STATE(92), 1,
      sym_primary_expression,
    STATE(262), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(64), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(66), 2,
      anon_sym_true,
      anon_sym_false,
    STATE(238), 2,
      sym_expression_list,
      sym__expression,
    ACTIONS(60), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(100), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [1246] = 19,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(76), 1,
      aux_sym_identifier_token1,
    ACTIONS(78), 1,
      anon_sym_DOT,
    ACTIONS(80), 1,
      anon_sym_LPAREN,
    ACTIONS(82), 1,
      anon_sym_LBRACK,
    ACTIONS(84), 1,
      anon_sym_DQUOTE,
    ACTIONS(86), 1,
      sym_zero_literal,
    ACTIONS(88), 1,
      sym_decimal_literal,
    ACTIONS(92), 1,
      sym_bit_literal,
    STATE(77), 1,
      sym__integer_literal,
    STATE(85), 1,
      sym_declared_identifier,
    STATE(102), 1,
      sym_qualified_identifier,
    STATE(109), 1,
      sym_primary_expression,
    STATE(248), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(94), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(96), 2,
      anon_sym_true,
      anon_sym_false,
    STATE(235), 2,
      sym_expression_list,
      sym__expression,
    ACTIONS(90), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(116), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [1316] = 19,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(44), 1,
      aux_sym_identifier_token1,
    ACTIONS(46), 1,
      anon_sym_DOT,
    ACTIONS(48), 1,
      anon_sym_LPAREN,
    ACTIONS(52), 1,
      anon_sym_LBRACK,
    ACTIONS(54), 1,
      anon_sym_DQUOTE,
    ACTIONS(56), 1,
      sym_zero_literal,
    ACTIONS(58), 1,
      sym_decimal_literal,
    ACTIONS(62), 1,
      sym_bit_literal,
    STATE(73), 1,
      sym__integer_literal,
    STATE(81), 1,
      sym_declared_identifier,
    STATE(86), 1,
      sym_qualified_identifier,
    STATE(92), 1,
      sym_primary_expression,
    STATE(262), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(64), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(66), 2,
      anon_sym_true,
      anon_sym_false,
    STATE(239), 2,
      sym_expression_list,
      sym__expression,
    ACTIONS(60), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(100), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [1386] = 19,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(44), 1,
      aux_sym_identifier_token1,
    ACTIONS(46), 1,
      anon_sym_DOT,
    ACTIONS(48), 1,
      anon_sym_LPAREN,
    ACTIONS(52), 1,
      anon_sym_LBRACK,
    ACTIONS(54), 1,
      anon_sym_DQUOTE,
    ACTIONS(56), 1,
      sym_zero_literal,
    ACTIONS(58), 1,
      sym_decimal_literal,
    ACTIONS(62), 1,
      sym_bit_literal,
    STATE(73), 1,
      sym__integer_literal,
    STATE(81), 1,
      sym_declared_identifier,
    STATE(86), 1,
      sym_qualified_identifier,
    STATE(92), 1,
      sym_primary_expression,
    STATE(262), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(64), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(66), 2,
      anon_sym_true,
      anon_sym_false,
    STATE(232), 2,
      sym_expression_list,
      sym__expression,
    ACTIONS(60), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(100), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [1456] = 19,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(44), 1,
      aux_sym_identifier_token1,
    ACTIONS(46), 1,
      anon_sym_DOT,
    ACTIONS(48), 1,
      anon_sym_LPAREN,
    ACTIONS(52), 1,
      anon_sym_LBRACK,
    ACTIONS(54), 1,
      anon_sym_DQUOTE,
    ACTIONS(56), 1,
      sym_zero_literal,
    ACTIONS(58), 1,
      sym_decimal_literal,
    ACTIONS(62), 1,
      sym_bit_literal,
    STATE(73), 1,
      sym__integer_literal,
    STATE(81), 1,
      sym_declared_identifier,
    STATE(86), 1,
      sym_qualified_identifier,
    STATE(92), 1,
      sym_primary_expression,
    STATE(262), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(64), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(66), 2,
      anon_sym_true,
      anon_sym_false,
    STATE(266), 2,
      sym_expression_list,
      sym__expression,
    ACTIONS(60), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(100), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [1526] = 19,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(76), 1,
      aux_sym_identifier_token1,
    ACTIONS(78), 1,
      anon_sym_DOT,
    ACTIONS(80), 1,
      anon_sym_LPAREN,
    ACTIONS(82), 1,
      anon_sym_LBRACK,
    ACTIONS(84), 1,
      anon_sym_DQUOTE,
    ACTIONS(86), 1,
      sym_zero_literal,
    ACTIONS(88), 1,
      sym_decimal_literal,
    ACTIONS(92), 1,
      sym_bit_literal,
    STATE(77), 1,
      sym__integer_literal,
    STATE(85), 1,
      sym_declared_identifier,
    STATE(102), 1,
      sym_qualified_identifier,
    STATE(109), 1,
      sym_primary_expression,
    STATE(248), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(94), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(96), 2,
      anon_sym_true,
      anon_sym_false,
    STATE(234), 2,
      sym_expression_list,
      sym__expression,
    ACTIONS(90), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(116), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [1596] = 17,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(76), 1,
      aux_sym_identifier_token1,
    ACTIONS(78), 1,
      anon_sym_DOT,
    ACTIONS(80), 1,
      anon_sym_LPAREN,
    ACTIONS(84), 1,
      anon_sym_DQUOTE,
    ACTIONS(86), 1,
      sym_zero_literal,
    ACTIONS(88), 1,
      sym_decimal_literal,
    ACTIONS(92), 1,
      sym_bit_literal,
    STATE(77), 1,
      sym__integer_literal,
    STATE(85), 1,
      sym_declared_identifier,
    STATE(102), 1,
      sym_qualified_identifier,
    STATE(106), 1,
      sym_primary_expression,
    STATE(248), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(94), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(96), 2,
      anon_sym_true,
      anon_sym_false,
    ACTIONS(90), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(116), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [1659] = 17,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(44), 1,
      aux_sym_identifier_token1,
    ACTIONS(46), 1,
      anon_sym_DOT,
    ACTIONS(48), 1,
      anon_sym_LPAREN,
    ACTIONS(54), 1,
      anon_sym_DQUOTE,
    ACTIONS(56), 1,
      sym_zero_literal,
    ACTIONS(58), 1,
      sym_decimal_literal,
    ACTIONS(62), 1,
      sym_bit_literal,
    STATE(73), 1,
      sym__integer_literal,
    STATE(81), 1,
      sym_declared_identifier,
    STATE(86), 1,
      sym_qualified_identifier,
    STATE(94), 1,
      sym_primary_expression,
    STATE(262), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(64), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(66), 2,
      anon_sym_true,
      anon_sym_false,
    ACTIONS(60), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(100), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [1722] = 17,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(76), 1,
      aux_sym_identifier_token1,
    ACTIONS(78), 1,
      anon_sym_DOT,
    ACTIONS(80), 1,
      anon_sym_LPAREN,
    ACTIONS(84), 1,
      anon_sym_DQUOTE,
    ACTIONS(86), 1,
      sym_zero_literal,
    ACTIONS(88), 1,
      sym_decimal_literal,
    ACTIONS(92), 1,
      sym_bit_literal,
    STATE(77), 1,
      sym__integer_literal,
    STATE(85), 1,
      sym_declared_identifier,
    STATE(102), 1,
      sym_qualified_identifier,
    STATE(117), 1,
      sym_primary_expression,
    STATE(248), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(94), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(96), 2,
      anon_sym_true,
      anon_sym_false,
    ACTIONS(90), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(116), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [1785] = 17,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(44), 1,
      aux_sym_identifier_token1,
    ACTIONS(46), 1,
      anon_sym_DOT,
    ACTIONS(48), 1,
      anon_sym_LPAREN,
    ACTIONS(54), 1,
      anon_sym_DQUOTE,
    ACTIONS(56), 1,
      sym_zero_literal,
    ACTIONS(58), 1,
      sym_decimal_literal,
    ACTIONS(62), 1,
      sym_bit_literal,
    STATE(73), 1,
      sym__integer_literal,
    STATE(81), 1,
      sym_declared_identifier,
    STATE(86), 1,
      sym_qualified_identifier,
    STATE(95), 1,
      sym_primary_expression,
    STATE(262), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(64), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(66), 2,
      anon_sym_true,
      anon_sym_false,
    ACTIONS(60), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(100), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [1848] = 17,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(44), 1,
      aux_sym_identifier_token1,
    ACTIONS(46), 1,
      anon_sym_DOT,
    ACTIONS(48), 1,
      anon_sym_LPAREN,
    ACTIONS(54), 1,
      anon_sym_DQUOTE,
    ACTIONS(56), 1,
      sym_zero_literal,
    ACTIONS(58), 1,
      sym_decimal_literal,
    ACTIONS(62), 1,
      sym_bit_literal,
    STATE(73), 1,
      sym__integer_literal,
    STATE(81), 1,
      sym_declared_identifier,
    STATE(86), 1,
      sym_qualified_identifier,
    STATE(90), 1,
      sym_primary_expression,
    STATE(262), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(64), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(66), 2,
      anon_sym_true,
      anon_sym_false,
    ACTIONS(60), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(100), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [1911] = 17,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(44), 1,
      aux_sym_identifier_token1,
    ACTIONS(46), 1,
      anon_sym_DOT,
    ACTIONS(48), 1,
      anon_sym_LPAREN,
    ACTIONS(54), 1,
      anon_sym_DQUOTE,
    ACTIONS(56), 1,
      sym_zero_literal,
    ACTIONS(58), 1,
      sym_decimal_literal,
    ACTIONS(62), 1,
      sym_bit_literal,
    STATE(73), 1,
      sym__integer_literal,
    STATE(81), 1,
      sym_declared_identifier,
    STATE(86), 1,
      sym_qualified_identifier,
    STATE(96), 1,
      sym_primary_expression,
    STATE(262), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(64), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(66), 2,
      anon_sym_true,
      anon_sym_false,
    ACTIONS(60), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(100), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [1974] = 17,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(76), 1,
      aux_sym_identifier_token1,
    ACTIONS(78), 1,
      anon_sym_DOT,
    ACTIONS(80), 1,
      anon_sym_LPAREN,
    ACTIONS(84), 1,
      anon_sym_DQUOTE,
    ACTIONS(86), 1,
      sym_zero_literal,
    ACTIONS(88), 1,
      sym_decimal_literal,
    ACTIONS(92), 1,
      sym_bit_literal,
    STATE(77), 1,
      sym__integer_literal,
    STATE(85), 1,
      sym_declared_identifier,
    STATE(102), 1,
      sym_qualified_identifier,
    STATE(107), 1,
      sym_primary_expression,
    STATE(248), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(94), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(96), 2,
      anon_sym_true,
      anon_sym_false,
    ACTIONS(90), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(116), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [2037] = 17,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(76), 1,
      aux_sym_identifier_token1,
    ACTIONS(78), 1,
      anon_sym_DOT,
    ACTIONS(80), 1,
      anon_sym_LPAREN,
    ACTIONS(84), 1,
      anon_sym_DQUOTE,
    ACTIONS(86), 1,
      sym_zero_literal,
    ACTIONS(88), 1,
      sym_decimal_literal,
    ACTIONS(92), 1,
      sym_bit_literal,
    STATE(77), 1,
      sym__integer_literal,
    STATE(85), 1,
      sym_declared_identifier,
    STATE(102), 1,
      sym_qualified_identifier,
    STATE(112), 1,
      sym_primary_expression,
    STATE(248), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(94), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(96), 2,
      anon_sym_true,
      anon_sym_false,
    ACTIONS(90), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(116), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [2100] = 17,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(76), 1,
      aux_sym_identifier_token1,
    ACTIONS(78), 1,
      anon_sym_DOT,
    ACTIONS(80), 1,
      anon_sym_LPAREN,
    ACTIONS(84), 1,
      anon_sym_DQUOTE,
    ACTIONS(86), 1,
      sym_zero_literal,
    ACTIONS(88), 1,
      sym_decimal_literal,
    ACTIONS(92), 1,
      sym_bit_literal,
    STATE(77), 1,
      sym__integer_literal,
    STATE(85), 1,
      sym_declared_identifier,
    STATE(102), 1,
      sym_qualified_identifier,
    STATE(114), 1,
      sym_primary_expression,
    STATE(248), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(94), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(96), 2,
      anon_sym_true,
      anon_sym_false,
    ACTIONS(90), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(116), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [2163] = 17,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(44), 1,
      aux_sym_identifier_token1,
    ACTIONS(46), 1,
      anon_sym_DOT,
    ACTIONS(48), 1,
      anon_sym_LPAREN,
    ACTIONS(54), 1,
      anon_sym_DQUOTE,
    ACTIONS(56), 1,
      sym_zero_literal,
    ACTIONS(58), 1,
      sym_decimal_literal,
    ACTIONS(62), 1,
      sym_bit_literal,
    STATE(73), 1,
      sym__integer_literal,
    STATE(81), 1,
      sym_declared_identifier,
    STATE(86), 1,
      sym_qualified_identifier,
    STATE(103), 1,
      sym_primary_expression,
    STATE(262), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(64), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(66), 2,
      anon_sym_true,
      anon_sym_false,
    ACTIONS(60), 3,
      sym_binary_literal,
      sym_octal_literal,
      sym_hex_literal,
    STATE(100), 8,
      sym_string_literal,
      sym_float_literal,
      sym_time_literal,
      sym_unary_operation,
      sym_binary_operation,
      sym_call,
      sym_parenthesized_expression,
      sym_subscript,
  [2226] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(102), 1,
      sym__indent,
    ACTIONS(100), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(98), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2261] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(102), 1,
      sym__indent,
    ACTIONS(106), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(104), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2296] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(102), 1,
      sym__indent,
    ACTIONS(110), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(108), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2331] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(102), 1,
      sym__indent,
    ACTIONS(114), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(112), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2366] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(102), 1,
      sym__indent,
    ACTIONS(118), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(116), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2401] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(102), 1,
      sym__indent,
    ACTIONS(122), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(120), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2436] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(110), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(108), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2468] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(126), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(124), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2500] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(130), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(128), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2532] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(134), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(132), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2564] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(138), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(136), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2596] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(142), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(140), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2628] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(106), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(104), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2660] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(146), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(144), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2692] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(150), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(148), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2724] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(154), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(152), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2756] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(122), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(120), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2788] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(158), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(156), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2820] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(162), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(160), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2852] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(114), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(112), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2884] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(166), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(164), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2916] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(170), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(168), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2948] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(174), 11,
      sym__dedent,
      anon_sym_add_DASHenable,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_out_DASHtrigger,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
    ACTIONS(172), 13,
      aux_sym_identifier_token1,
      anon_sym_access,
      anon_sym_atomic,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_groups,
      anon_sym_masters,
      anon_sym_range,
      anon_sym_reset,
      anon_sym_size,
      anon_sym_width,
      anon_sym_const,
      anon_sym_type,
  [2980] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(15), 1,
      anon_sym_reset,
    STATE(41), 1,
      sym_property_assignments,
    STATE(196), 1,
      sym__property_assignment,
    STATE(257), 1,
      sym_property_identifier,
    ACTIONS(17), 19,
      anon_sym_access,
      anon_sym_add_DASHenable,
      anon_sym_atomic,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_groups,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_masters,
      anon_sym_out_DASHtrigger,
      anon_sym_range,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
      anon_sym_size,
      anon_sym_width,
  [3017] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(15), 1,
      anon_sym_reset,
    STATE(164), 1,
      sym_property_assignments,
    STATE(190), 1,
      sym__property_assignment,
    STATE(257), 1,
      sym_property_identifier,
    ACTIONS(17), 19,
      anon_sym_access,
      anon_sym_add_DASHenable,
      anon_sym_atomic,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_groups,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_masters,
      anon_sym_out_DASHtrigger,
      anon_sym_range,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
      anon_sym_size,
      anon_sym_width,
  [3054] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(15), 1,
      anon_sym_reset,
    STATE(45), 1,
      sym_property_assignments,
    STATE(196), 1,
      sym__property_assignment,
    STATE(257), 1,
      sym_property_identifier,
    ACTIONS(17), 19,
      anon_sym_access,
      anon_sym_add_DASHenable,
      anon_sym_atomic,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_groups,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_masters,
      anon_sym_out_DASHtrigger,
      anon_sym_range,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
      anon_sym_size,
      anon_sym_width,
  [3091] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(15), 1,
      anon_sym_reset,
    STATE(46), 1,
      sym_property_assignments,
    STATE(196), 1,
      sym__property_assignment,
    STATE(257), 1,
      sym_property_identifier,
    ACTIONS(17), 19,
      anon_sym_access,
      anon_sym_add_DASHenable,
      anon_sym_atomic,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_groups,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_masters,
      anon_sym_out_DASHtrigger,
      anon_sym_range,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
      anon_sym_size,
      anon_sym_width,
  [3128] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(15), 1,
      anon_sym_reset,
    STATE(51), 1,
      sym_property_assignments,
    STATE(196), 1,
      sym__property_assignment,
    STATE(257), 1,
      sym_property_identifier,
    ACTIONS(17), 19,
      anon_sym_access,
      anon_sym_add_DASHenable,
      anon_sym_atomic,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_groups,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_masters,
      anon_sym_out_DASHtrigger,
      anon_sym_range,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
      anon_sym_size,
      anon_sym_width,
  [3165] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(15), 1,
      anon_sym_reset,
    STATE(54), 1,
      sym_property_assignments,
    STATE(196), 1,
      sym__property_assignment,
    STATE(257), 1,
      sym_property_identifier,
    ACTIONS(17), 19,
      anon_sym_access,
      anon_sym_add_DASHenable,
      anon_sym_atomic,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_groups,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_masters,
      anon_sym_out_DASHtrigger,
      anon_sym_range,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
      anon_sym_size,
      anon_sym_width,
  [3202] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(15), 1,
      anon_sym_reset,
    STATE(148), 1,
      sym_property_assignments,
    STATE(190), 1,
      sym__property_assignment,
    STATE(257), 1,
      sym_property_identifier,
    ACTIONS(17), 19,
      anon_sym_access,
      anon_sym_add_DASHenable,
      anon_sym_atomic,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_groups,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_masters,
      anon_sym_out_DASHtrigger,
      anon_sym_range,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
      anon_sym_size,
      anon_sym_width,
  [3239] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(15), 1,
      anon_sym_reset,
    STATE(173), 1,
      sym_property_assignments,
    STATE(190), 1,
      sym__property_assignment,
    STATE(257), 1,
      sym_property_identifier,
    ACTIONS(17), 19,
      anon_sym_access,
      anon_sym_add_DASHenable,
      anon_sym_atomic,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_groups,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_masters,
      anon_sym_out_DASHtrigger,
      anon_sym_range,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
      anon_sym_size,
      anon_sym_width,
  [3276] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(15), 1,
      anon_sym_reset,
    STATE(147), 1,
      sym_property_assignments,
    STATE(190), 1,
      sym__property_assignment,
    STATE(257), 1,
      sym_property_identifier,
    ACTIONS(17), 19,
      anon_sym_access,
      anon_sym_add_DASHenable,
      anon_sym_atomic,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_groups,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_masters,
      anon_sym_out_DASHtrigger,
      anon_sym_range,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
      anon_sym_size,
      anon_sym_width,
  [3313] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(15), 1,
      anon_sym_reset,
    STATE(166), 1,
      sym_property_assignments,
    STATE(190), 1,
      sym__property_assignment,
    STATE(257), 1,
      sym_property_identifier,
    ACTIONS(17), 19,
      anon_sym_access,
      anon_sym_add_DASHenable,
      anon_sym_atomic,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_groups,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_masters,
      anon_sym_out_DASHtrigger,
      anon_sym_range,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
      anon_sym_size,
      anon_sym_width,
  [3350] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(15), 1,
      anon_sym_reset,
    STATE(169), 1,
      sym_property_assignments,
    STATE(190), 1,
      sym__property_assignment,
    STATE(257), 1,
      sym_property_identifier,
    ACTIONS(17), 19,
      anon_sym_access,
      anon_sym_add_DASHenable,
      anon_sym_atomic,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_groups,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_masters,
      anon_sym_out_DASHtrigger,
      anon_sym_range,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
      anon_sym_size,
      anon_sym_width,
  [3387] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(15), 1,
      anon_sym_reset,
    STATE(47), 1,
      sym_property_assignments,
    STATE(196), 1,
      sym__property_assignment,
    STATE(257), 1,
      sym_property_identifier,
    ACTIONS(17), 19,
      anon_sym_access,
      anon_sym_add_DASHenable,
      anon_sym_atomic,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_groups,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_masters,
      anon_sym_out_DASHtrigger,
      anon_sym_range,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
      anon_sym_size,
      anon_sym_width,
  [3424] = 5,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(15), 1,
      anon_sym_reset,
    STATE(223), 1,
      sym__property_assignment,
    STATE(257), 1,
      sym_property_identifier,
    ACTIONS(17), 19,
      anon_sym_access,
      anon_sym_add_DASHenable,
      anon_sym_atomic,
      anon_sym_byte_DASHwrite_DASHenable,
      anon_sym_clear,
      anon_sym_delay,
      anon_sym_enable_DASHinit_DASHvalue,
      anon_sym_enable_DASHreset_DASHvalue,
      anon_sym_groups,
      anon_sym_init_DASHvalue,
      anon_sym_in_DASHtrigger,
      anon_sym_masters,
      anon_sym_out_DASHtrigger,
      anon_sym_range,
      anon_sym_read_DASHlatency,
      anon_sym_read_DASHvalue,
      anon_sym_reset_DASHvalue,
      anon_sym_size,
      anon_sym_width,
  [3458] = 8,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(7), 1,
      aux_sym_identifier_token1,
    ACTIONS(9), 1,
      anon_sym_import,
    ACTIONS(11), 1,
      anon_sym_const,
    ACTIONS(13), 1,
      anon_sym_type,
    ACTIONS(176), 1,
      ts_builtin_sym_end,
    STATE(142), 1,
      sym_identifier,
    STATE(72), 14,
      sym__meta_statement,
      sym__import_statement,
      sym_single_import_statement,
      sym_multi_import_statement,
      sym__constant_definition,
      sym_single_constant_definition,
      sym_multi_constant_definition,
      sym__type_definition,
      sym_single_line_type_definition,
      sym_multi_line_type_definition,
      sym__instantiation,
      sym_single_line_instantiation,
      sym_multi_line_instantiation,
      aux_sym_description_repeat1,
  [3496] = 8,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(178), 1,
      ts_builtin_sym_end,
    ACTIONS(180), 1,
      aux_sym_identifier_token1,
    ACTIONS(183), 1,
      anon_sym_import,
    ACTIONS(186), 1,
      anon_sym_const,
    ACTIONS(189), 1,
      anon_sym_type,
    STATE(142), 1,
      sym_identifier,
    STATE(72), 14,
      sym__meta_statement,
      sym__import_statement,
      sym_single_import_statement,
      sym_multi_import_statement,
      sym__constant_definition,
      sym_single_constant_definition,
      sym_multi_constant_definition,
      sym__type_definition,
      sym_single_line_type_definition,
      sym_multi_line_type_definition,
      sym__instantiation,
      sym_single_line_instantiation,
      sym_multi_line_instantiation,
      aux_sym_description_repeat1,
  [3534] = 5,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(196), 1,
      anon_sym_STAR,
    STATE(104), 1,
      sym_time_unit,
    ACTIONS(194), 4,
      anon_sym_ns,
      anon_sym_us,
      anon_sym_ms,
      anon_sym_s,
    ACTIONS(192), 10,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [3562] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(198), 1,
      anon_sym_DOT,
    ACTIONS(202), 1,
      anon_sym_STAR,
    ACTIONS(200), 14,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_ns,
      anon_sym_us,
      anon_sym_ms,
      anon_sym_s,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [3588] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(206), 1,
      anon_sym_STAR,
    ACTIONS(204), 14,
      anon_sym_DOT,
      anon_sym_EQ,
      anon_sym_COMMA,
      anon_sym_LPAREN,
      anon_sym_RPAREN,
      anon_sym_LBRACK,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [3611] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(198), 1,
      anon_sym_DOT,
    ACTIONS(202), 1,
      anon_sym_STAR,
    ACTIONS(200), 13,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_ns,
      anon_sym_us,
      anon_sym_ms,
      anon_sym_s,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [3636] = 5,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(196), 1,
      anon_sym_STAR,
    STATE(119), 1,
      sym_time_unit,
    ACTIONS(208), 4,
      anon_sym_ns,
      anon_sym_us,
      anon_sym_ms,
      anon_sym_s,
    ACTIONS(192), 9,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [3663] = 5,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(214), 1,
      anon_sym_STAR,
    STATE(97), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(212), 2,
      sym_zero_literal,
      sym_decimal_literal,
    ACTIONS(210), 10,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [3689] = 5,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(220), 1,
      anon_sym_STAR,
    STATE(101), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(218), 2,
      sym_zero_literal,
      sym_decimal_literal,
    ACTIONS(216), 10,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [3715] = 7,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(196), 1,
      anon_sym_STAR,
    ACTIONS(222), 1,
      anon_sym_DOT,
    ACTIONS(224), 1,
      anon_sym_EQ,
    ACTIONS(226), 1,
      anon_sym_LPAREN,
    ACTIONS(228), 1,
      anon_sym_LBRACK,
    ACTIONS(192), 9,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [3745] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(196), 1,
      anon_sym_STAR,
    ACTIONS(222), 1,
      anon_sym_DOT,
    ACTIONS(226), 1,
      anon_sym_LPAREN,
    ACTIONS(228), 1,
      anon_sym_LBRACK,
    ACTIONS(192), 10,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [3773] = 5,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(220), 1,
      anon_sym_STAR,
    STATE(115), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(230), 2,
      sym_zero_literal,
      sym_decimal_literal,
    ACTIONS(216), 9,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [3798] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(206), 1,
      anon_sym_STAR,
    ACTIONS(204), 12,
      sym__newline,
      anon_sym_DOT,
      anon_sym_LPAREN,
      anon_sym_SEMI,
      anon_sym_LBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [3819] = 5,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(214), 1,
      anon_sym_STAR,
    STATE(118), 1,
      sym__zero_or_decimal_literal,
    ACTIONS(232), 2,
      sym_zero_literal,
      sym_decimal_literal,
    ACTIONS(210), 9,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [3844] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(196), 1,
      anon_sym_STAR,
    ACTIONS(234), 1,
      anon_sym_DOT,
    ACTIONS(236), 1,
      anon_sym_LPAREN,
    ACTIONS(238), 1,
      anon_sym_LBRACK,
    ACTIONS(192), 9,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [3871] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(196), 1,
      anon_sym_STAR,
    ACTIONS(228), 1,
      anon_sym_LBRACK,
    ACTIONS(192), 10,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [3893] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(242), 1,
      anon_sym_STAR,
    ACTIONS(240), 11,
      sym__newline,
      anon_sym_LPAREN,
      anon_sym_SEMI,
      anon_sym_LBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [3913] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(242), 1,
      anon_sym_STAR,
    ACTIONS(240), 11,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_LBRACK,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [3933] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(246), 1,
      anon_sym_STAR,
    ACTIONS(244), 10,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [3952] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(250), 1,
      anon_sym_STAR,
    ACTIONS(252), 1,
      anon_sym_STAR_STAR,
    ACTIONS(248), 9,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [3973] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(256), 1,
      anon_sym_STAR,
    ACTIONS(254), 10,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [3992] = 7,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(252), 1,
      anon_sym_STAR_STAR,
    ACTIONS(262), 1,
      anon_sym_STAR,
    ACTIONS(260), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(264), 2,
      anon_sym_SLASH,
      anon_sym_PERCENT,
    ACTIONS(266), 2,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
    ACTIONS(258), 3,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
  [4019] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(270), 1,
      anon_sym_STAR,
    ACTIONS(268), 10,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4038] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(252), 1,
      anon_sym_STAR_STAR,
    ACTIONS(262), 1,
      anon_sym_STAR,
    ACTIONS(260), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(264), 2,
      anon_sym_SLASH,
      anon_sym_PERCENT,
    ACTIONS(248), 5,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4063] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(250), 1,
      anon_sym_STAR,
    ACTIONS(252), 1,
      anon_sym_STAR_STAR,
    ACTIONS(248), 9,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4084] = 5,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(252), 1,
      anon_sym_STAR_STAR,
    ACTIONS(262), 1,
      anon_sym_STAR,
    ACTIONS(264), 2,
      anon_sym_SLASH,
      anon_sym_PERCENT,
    ACTIONS(248), 7,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4107] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(274), 1,
      anon_sym_STAR,
    ACTIONS(272), 10,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4126] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(278), 1,
      anon_sym_STAR,
    ACTIONS(276), 10,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4145] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(282), 1,
      anon_sym_STAR,
    ACTIONS(280), 10,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4164] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(196), 1,
      anon_sym_STAR,
    ACTIONS(192), 10,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4183] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(214), 1,
      anon_sym_STAR,
    ACTIONS(210), 10,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4202] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(196), 1,
      anon_sym_STAR,
    ACTIONS(238), 1,
      anon_sym_LBRACK,
    ACTIONS(192), 9,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4223] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(252), 1,
      anon_sym_STAR_STAR,
    ACTIONS(286), 1,
      anon_sym_STAR,
    ACTIONS(284), 9,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4244] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(290), 1,
      anon_sym_STAR,
    ACTIONS(288), 10,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4263] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(294), 1,
      anon_sym_STAR,
    ACTIONS(292), 10,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4282] = 5,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(296), 1,
      anon_sym_STAR,
    ACTIONS(300), 1,
      anon_sym_STAR_STAR,
    ACTIONS(298), 2,
      anon_sym_SLASH,
      anon_sym_PERCENT,
    ACTIONS(248), 6,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4304] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(296), 1,
      anon_sym_STAR,
    ACTIONS(300), 1,
      anon_sym_STAR_STAR,
    ACTIONS(298), 2,
      anon_sym_SLASH,
      anon_sym_PERCENT,
    ACTIONS(302), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(248), 4,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4328] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(294), 1,
      anon_sym_STAR,
    ACTIONS(292), 9,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4346] = 7,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(296), 1,
      anon_sym_STAR,
    ACTIONS(300), 1,
      anon_sym_STAR_STAR,
    ACTIONS(258), 2,
      sym__newline,
      anon_sym_SEMI,
    ACTIONS(298), 2,
      anon_sym_SLASH,
      anon_sym_PERCENT,
    ACTIONS(302), 2,
      anon_sym_PLUS,
      anon_sym_DASH,
    ACTIONS(304), 2,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4372] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(278), 1,
      anon_sym_STAR,
    ACTIONS(276), 9,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4390] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(270), 1,
      anon_sym_STAR,
    ACTIONS(268), 9,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4408] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(250), 1,
      anon_sym_STAR,
    ACTIONS(300), 1,
      anon_sym_STAR_STAR,
    ACTIONS(248), 8,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4428] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(282), 1,
      anon_sym_STAR,
    ACTIONS(280), 9,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4446] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(250), 1,
      anon_sym_STAR,
    ACTIONS(300), 1,
      anon_sym_STAR_STAR,
    ACTIONS(248), 8,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4466] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(214), 1,
      anon_sym_STAR,
    ACTIONS(210), 9,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4484] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(196), 1,
      anon_sym_STAR,
    ACTIONS(192), 9,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4502] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(286), 1,
      anon_sym_STAR,
    ACTIONS(300), 1,
      anon_sym_STAR_STAR,
    ACTIONS(284), 8,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4522] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(274), 1,
      anon_sym_STAR,
    ACTIONS(272), 9,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4540] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(290), 1,
      anon_sym_STAR,
    ACTIONS(288), 9,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4558] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(256), 1,
      anon_sym_STAR,
    ACTIONS(254), 9,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4576] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(246), 1,
      anon_sym_STAR,
    ACTIONS(244), 9,
      sym__newline,
      anon_sym_SEMI,
      anon_sym_PLUS,
      anon_sym_DASH,
      anon_sym_SLASH,
      anon_sym_PERCENT,
      anon_sym_STAR_STAR,
      anon_sym_LT_LT,
      anon_sym_GT_GT,
  [4594] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(306), 7,
      aux_sym_identifier_token1,
      anon_sym_EQ,
      anon_sym_COMMA,
      anon_sym_LPAREN,
      anon_sym_RPAREN,
      anon_sym_LBRACK,
      anon_sym_DQUOTE,
  [4607] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(110), 1,
      ts_builtin_sym_end,
    ACTIONS(308), 1,
      sym__indent,
    ACTIONS(108), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [4623] = 7,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(234), 1,
      anon_sym_DOT,
    ACTIONS(310), 1,
      anon_sym_LPAREN,
    ACTIONS(312), 1,
      anon_sym_SEMI,
    ACTIONS(314), 1,
      sym__newline,
    STATE(165), 1,
      sym_element_body,
    STATE(213), 1,
      sym_argument_list,
  [4645] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(100), 1,
      ts_builtin_sym_end,
    ACTIONS(308), 1,
      sym__indent,
    ACTIONS(98), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [4661] = 7,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(234), 1,
      anon_sym_DOT,
    ACTIONS(310), 1,
      anon_sym_LPAREN,
    ACTIONS(316), 1,
      anon_sym_SEMI,
    ACTIONS(318), 1,
      sym__newline,
    STATE(57), 1,
      sym_element_body,
    STATE(193), 1,
      sym_argument_list,
  [4683] = 7,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(234), 1,
      anon_sym_DOT,
    ACTIONS(310), 1,
      anon_sym_LPAREN,
    ACTIONS(320), 1,
      anon_sym_SEMI,
    ACTIONS(322), 1,
      sym__newline,
    STATE(139), 1,
      sym_element_body,
    STATE(200), 1,
      sym_argument_list,
  [4705] = 7,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(234), 1,
      anon_sym_DOT,
    ACTIONS(310), 1,
      anon_sym_LPAREN,
    ACTIONS(324), 1,
      anon_sym_SEMI,
    ACTIONS(326), 1,
      sym__newline,
    STATE(44), 1,
      sym_element_body,
    STATE(201), 1,
      sym_argument_list,
  [4727] = 7,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(234), 1,
      anon_sym_DOT,
    ACTIONS(310), 1,
      anon_sym_LPAREN,
    ACTIONS(328), 1,
      anon_sym_SEMI,
    ACTIONS(330), 1,
      sym__newline,
    STATE(43), 1,
      sym_element_body,
    STATE(206), 1,
      sym_argument_list,
  [4749] = 7,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(234), 1,
      anon_sym_DOT,
    ACTIONS(310), 1,
      anon_sym_LPAREN,
    ACTIONS(332), 1,
      anon_sym_SEMI,
    ACTIONS(334), 1,
      sym__newline,
    STATE(172), 1,
      sym_element_body,
    STATE(210), 1,
      sym_argument_list,
  [4771] = 7,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(234), 1,
      anon_sym_DOT,
    ACTIONS(310), 1,
      anon_sym_LPAREN,
    ACTIONS(336), 1,
      anon_sym_SEMI,
    ACTIONS(338), 1,
      sym__newline,
    STATE(53), 1,
      sym_element_body,
    STATE(195), 1,
      sym_argument_list,
  [4793] = 7,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(84), 1,
      anon_sym_DQUOTE,
    ACTIONS(340), 1,
      aux_sym_identifier_token1,
    ACTIONS(342), 1,
      sym__dedent,
    STATE(135), 1,
      aux_sym_multi_import_statement_repeat1,
    STATE(216), 1,
      sym_identifier,
    STATE(249), 1,
      sym_string_literal,
  [4815] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(106), 1,
      ts_builtin_sym_end,
    ACTIONS(308), 1,
      sym__indent,
    ACTIONS(104), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [4831] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(118), 1,
      ts_builtin_sym_end,
    ACTIONS(308), 1,
      sym__indent,
    ACTIONS(116), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [4847] = 7,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(344), 1,
      aux_sym_identifier_token1,
    ACTIONS(347), 1,
      anon_sym_DQUOTE,
    ACTIONS(350), 1,
      sym__dedent,
    STATE(135), 1,
      aux_sym_multi_import_statement_repeat1,
    STATE(216), 1,
      sym_identifier,
    STATE(249), 1,
      sym_string_literal,
  [4869] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(114), 1,
      ts_builtin_sym_end,
    ACTIONS(308), 1,
      sym__indent,
    ACTIONS(112), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [4885] = 7,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(234), 1,
      anon_sym_DOT,
    ACTIONS(310), 1,
      anon_sym_LPAREN,
    ACTIONS(352), 1,
      anon_sym_SEMI,
    ACTIONS(354), 1,
      sym__newline,
    STATE(170), 1,
      sym_element_body,
    STATE(189), 1,
      sym_argument_list,
  [4907] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(122), 1,
      ts_builtin_sym_end,
    ACTIONS(308), 1,
      sym__indent,
    ACTIONS(120), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [4923] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(134), 1,
      ts_builtin_sym_end,
    ACTIONS(132), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [4936] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(356), 1,
      ts_builtin_sym_end,
    ACTIONS(358), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [4949] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(158), 1,
      ts_builtin_sym_end,
    ACTIONS(156), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [4962] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(360), 1,
      aux_sym_identifier_token1,
    ACTIONS(362), 1,
      anon_sym_LBRACK,
    STATE(124), 1,
      sym_declared_identifier,
    STATE(151), 1,
      sym_qualified_identifier,
    STATE(205), 1,
      sym_array_marker,
  [4981] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(360), 1,
      aux_sym_identifier_token1,
    ACTIONS(364), 1,
      anon_sym_LPAREN,
    STATE(128), 1,
      sym_declared_identifier,
    STATE(158), 1,
      sym_qualified_identifier,
    STATE(187), 1,
      sym_parameter_list,
  [5000] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(360), 1,
      aux_sym_identifier_token1,
    ACTIONS(362), 1,
      anon_sym_LBRACK,
    STATE(129), 1,
      sym_declared_identifier,
    STATE(159), 1,
      sym_qualified_identifier,
    STATE(191), 1,
      sym_array_marker,
  [5019] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(310), 1,
      anon_sym_LPAREN,
    ACTIONS(316), 1,
      anon_sym_SEMI,
    ACTIONS(318), 1,
      sym__newline,
    STATE(57), 1,
      sym_element_body,
    STATE(193), 1,
      sym_argument_list,
  [5038] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(360), 1,
      aux_sym_identifier_token1,
    ACTIONS(364), 1,
      anon_sym_LPAREN,
    STATE(127), 1,
      sym_declared_identifier,
    STATE(163), 1,
      sym_qualified_identifier,
    STATE(181), 1,
      sym_parameter_list,
  [5057] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(122), 1,
      ts_builtin_sym_end,
    ACTIONS(120), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [5070] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(114), 1,
      ts_builtin_sym_end,
    ACTIONS(112), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [5083] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(166), 1,
      ts_builtin_sym_end,
    ACTIONS(164), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [5096] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(146), 1,
      ts_builtin_sym_end,
    ACTIONS(144), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [5109] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(310), 1,
      anon_sym_LPAREN,
    ACTIONS(312), 1,
      anon_sym_SEMI,
    ACTIONS(314), 1,
      sym__newline,
    STATE(165), 1,
      sym_element_body,
    STATE(213), 1,
      sym_argument_list,
  [5128] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(170), 1,
      ts_builtin_sym_end,
    ACTIONS(168), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [5141] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(310), 1,
      anon_sym_LPAREN,
    ACTIONS(332), 1,
      anon_sym_SEMI,
    ACTIONS(334), 1,
      sym__newline,
    STATE(172), 1,
      sym_element_body,
    STATE(210), 1,
      sym_argument_list,
  [5160] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(310), 1,
      anon_sym_LPAREN,
    ACTIONS(336), 1,
      anon_sym_SEMI,
    ACTIONS(338), 1,
      sym__newline,
    STATE(53), 1,
      sym_element_body,
    STATE(195), 1,
      sym_argument_list,
  [5179] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(154), 1,
      ts_builtin_sym_end,
    ACTIONS(152), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [5192] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(84), 1,
      anon_sym_DQUOTE,
    ACTIONS(340), 1,
      aux_sym_identifier_token1,
    ACTIONS(366), 1,
      sym__newline,
    STATE(226), 1,
      sym_identifier,
    STATE(245), 1,
      sym_string_literal,
  [5211] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(84), 1,
      anon_sym_DQUOTE,
    ACTIONS(340), 1,
      aux_sym_identifier_token1,
    STATE(132), 1,
      aux_sym_multi_import_statement_repeat1,
    STATE(216), 1,
      sym_identifier,
    STATE(249), 1,
      sym_string_literal,
  [5230] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(310), 1,
      anon_sym_LPAREN,
    ACTIONS(324), 1,
      anon_sym_SEMI,
    ACTIONS(326), 1,
      sym__newline,
    STATE(44), 1,
      sym_element_body,
    STATE(201), 1,
      sym_argument_list,
  [5249] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(310), 1,
      anon_sym_LPAREN,
    ACTIONS(328), 1,
      anon_sym_SEMI,
    ACTIONS(330), 1,
      sym__newline,
    STATE(43), 1,
      sym_element_body,
    STATE(206), 1,
      sym_argument_list,
  [5268] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(126), 1,
      ts_builtin_sym_end,
    ACTIONS(124), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [5281] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(340), 1,
      aux_sym_identifier_token1,
    ACTIONS(368), 1,
      anon_sym_RPAREN,
    STATE(203), 1,
      sym__parameter,
    STATE(204), 1,
      sym_identifier,
    STATE(237), 1,
      sym__parameters,
  [5300] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(150), 1,
      ts_builtin_sym_end,
    ACTIONS(148), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [5313] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(310), 1,
      anon_sym_LPAREN,
    ACTIONS(320), 1,
      anon_sym_SEMI,
    ACTIONS(322), 1,
      sym__newline,
    STATE(139), 1,
      sym_element_body,
    STATE(200), 1,
      sym_argument_list,
  [5332] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(142), 1,
      ts_builtin_sym_end,
    ACTIONS(140), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [5345] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(130), 1,
      ts_builtin_sym_end,
    ACTIONS(128), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [5358] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(138), 1,
      ts_builtin_sym_end,
    ACTIONS(136), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [5371] = 6,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(310), 1,
      anon_sym_LPAREN,
    ACTIONS(352), 1,
      anon_sym_SEMI,
    ACTIONS(354), 1,
      sym__newline,
    STATE(170), 1,
      sym_element_body,
    STATE(189), 1,
      sym_argument_list,
  [5390] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(370), 1,
      ts_builtin_sym_end,
    ACTIONS(372), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [5403] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(106), 1,
      ts_builtin_sym_end,
    ACTIONS(104), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [5416] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(162), 1,
      ts_builtin_sym_end,
    ACTIONS(160), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [5429] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(374), 1,
      ts_builtin_sym_end,
    ACTIONS(376), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [5442] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(174), 1,
      ts_builtin_sym_end,
    ACTIONS(172), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [5455] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(110), 1,
      ts_builtin_sym_end,
    ACTIONS(108), 4,
      aux_sym_identifier_token1,
      anon_sym_import,
      anon_sym_const,
      anon_sym_type,
  [5468] = 5,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(340), 1,
      aux_sym_identifier_token1,
    ACTIONS(378), 1,
      sym__dedent,
    STATE(179), 1,
      aux_sym_multi_constant_definition_repeat1,
    STATE(233), 1,
      sym_identifier,
  [5484] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(380), 1,
      anon_sym_COMMA,
    STATE(178), 1,
      aux_sym__expressions_repeat1,
    ACTIONS(382), 2,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
  [5498] = 5,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(340), 1,
      aux_sym_identifier_token1,
    ACTIONS(384), 1,
      sym__dedent,
    STATE(179), 1,
      aux_sym_multi_constant_definition_repeat1,
    STATE(233), 1,
      sym_identifier,
  [5514] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(380), 1,
      anon_sym_COMMA,
    STATE(175), 1,
      aux_sym__expressions_repeat1,
    ACTIONS(386), 2,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
  [5528] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(388), 1,
      anon_sym_COMMA,
    STATE(178), 1,
      aux_sym__expressions_repeat1,
    ACTIONS(391), 2,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
  [5542] = 5,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(393), 1,
      aux_sym_identifier_token1,
    ACTIONS(396), 1,
      sym__dedent,
    STATE(179), 1,
      aux_sym_multi_constant_definition_repeat1,
    STATE(233), 1,
      sym_identifier,
  [5558] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(398), 3,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
  [5567] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(360), 1,
      aux_sym_identifier_token1,
    STATE(130), 1,
      sym_declared_identifier,
    STATE(153), 1,
      sym_qualified_identifier,
  [5580] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(340), 1,
      aux_sym_identifier_token1,
    ACTIONS(400), 1,
      sym__newline,
    STATE(265), 1,
      sym_identifier,
  [5593] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(402), 1,
      anon_sym_COMMA,
    ACTIONS(404), 1,
      anon_sym_RPAREN,
    STATE(197), 1,
      aux_sym__parameters_repeat1,
  [5606] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(340), 1,
      aux_sym_identifier_token1,
    STATE(204), 1,
      sym_identifier,
    STATE(225), 1,
      sym__parameter,
  [5619] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(340), 1,
      aux_sym_identifier_token1,
    ACTIONS(406), 1,
      sym__newline,
    STATE(241), 1,
      sym_identifier,
  [5632] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(350), 3,
      sym__dedent,
      aux_sym_identifier_token1,
      anon_sym_DQUOTE,
  [5641] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(360), 1,
      aux_sym_identifier_token1,
    STATE(126), 1,
      sym_declared_identifier,
    STATE(145), 1,
      sym_qualified_identifier,
  [5654] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(340), 1,
      aux_sym_identifier_token1,
    STATE(176), 1,
      aux_sym_multi_constant_definition_repeat1,
    STATE(233), 1,
      sym_identifier,
  [5667] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(408), 1,
      anon_sym_SEMI,
    ACTIONS(410), 1,
      sym__newline,
    STATE(141), 1,
      sym_element_body,
  [5680] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(412), 1,
      anon_sym_SEMI,
    ACTIONS(414), 1,
      sym__newline,
    STATE(209), 1,
      aux_sym_property_assignments_repeat1,
  [5693] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(360), 1,
      aux_sym_identifier_token1,
    STATE(131), 1,
      sym_declared_identifier,
    STATE(154), 1,
      sym_qualified_identifier,
  [5706] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(412), 1,
      anon_sym_SEMI,
    ACTIONS(416), 1,
      sym__newline,
    STATE(211), 1,
      aux_sym_property_assignments_repeat1,
  [5719] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(418), 1,
      anon_sym_SEMI,
    ACTIONS(420), 1,
      sym__newline,
    STATE(48), 1,
      sym_element_body,
  [5732] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(422), 1,
      anon_sym_COMMA,
    ACTIONS(424), 1,
      anon_sym_RPAREN,
    STATE(202), 1,
      aux_sym__arguments_repeat1,
  [5745] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(426), 1,
      anon_sym_SEMI,
    ACTIONS(428), 1,
      sym__newline,
    STATE(52), 1,
      sym_element_body,
  [5758] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(412), 1,
      anon_sym_SEMI,
    ACTIONS(430), 1,
      sym__newline,
    STATE(192), 1,
      aux_sym_property_assignments_repeat1,
  [5771] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(432), 1,
      anon_sym_COMMA,
    ACTIONS(435), 1,
      anon_sym_RPAREN,
    STATE(197), 1,
      aux_sym__parameters_repeat1,
  [5784] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(437), 3,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
  [5793] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(391), 3,
      anon_sym_COMMA,
      anon_sym_RPAREN,
      anon_sym_RBRACK,
  [5802] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(332), 1,
      anon_sym_SEMI,
    ACTIONS(334), 1,
      sym__newline,
    STATE(172), 1,
      sym_element_body,
  [5815] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(316), 1,
      anon_sym_SEMI,
    ACTIONS(318), 1,
      sym__newline,
    STATE(57), 1,
      sym_element_body,
  [5828] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(422), 1,
      anon_sym_COMMA,
    ACTIONS(439), 1,
      anon_sym_RPAREN,
    STATE(208), 1,
      aux_sym__arguments_repeat1,
  [5841] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(402), 1,
      anon_sym_COMMA,
    ACTIONS(441), 1,
      anon_sym_RPAREN,
    STATE(183), 1,
      aux_sym__parameters_repeat1,
  [5854] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(443), 1,
      anon_sym_EQ,
    ACTIONS(445), 2,
      anon_sym_COMMA,
      anon_sym_RPAREN,
  [5865] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(360), 1,
      aux_sym_identifier_token1,
    STATE(137), 1,
      sym_declared_identifier,
    STATE(167), 1,
      sym_qualified_identifier,
  [5878] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(336), 1,
      anon_sym_SEMI,
    ACTIONS(338), 1,
      sym__newline,
    STATE(53), 1,
      sym_element_body,
  [5891] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(447), 3,
      sym__dedent,
      aux_sym_identifier_token1,
      anon_sym_DQUOTE,
  [5900] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(449), 1,
      anon_sym_COMMA,
    ACTIONS(452), 1,
      anon_sym_RPAREN,
    STATE(208), 1,
      aux_sym__arguments_repeat1,
  [5913] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(412), 1,
      anon_sym_SEMI,
    ACTIONS(454), 1,
      sym__newline,
    STATE(211), 1,
      aux_sym_property_assignments_repeat1,
  [5926] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(456), 1,
      anon_sym_SEMI,
    ACTIONS(458), 1,
      sym__newline,
    STATE(150), 1,
      sym_element_body,
  [5939] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(460), 1,
      anon_sym_SEMI,
    ACTIONS(463), 1,
      sym__newline,
    STATE(211), 1,
      aux_sym_property_assignments_repeat1,
  [5952] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(340), 1,
      aux_sym_identifier_token1,
    STATE(174), 1,
      aux_sym_multi_constant_definition_repeat1,
    STATE(233), 1,
      sym_identifier,
  [5965] = 4,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(352), 1,
      anon_sym_SEMI,
    ACTIONS(354), 1,
      sym__newline,
    STATE(170), 1,
      sym_element_body,
  [5978] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(465), 2,
      sym__newline,
      anon_sym_SEMI,
  [5986] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(398), 2,
      sym__newline,
      anon_sym_SEMI,
  [5994] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(84), 1,
      anon_sym_DQUOTE,
    STATE(259), 1,
      sym_string_literal,
  [6004] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(467), 1,
      aux_sym_identifier_token1,
    STATE(88), 1,
      sym_declared_identifier,
  [6014] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(340), 1,
      aux_sym_identifier_token1,
    STATE(143), 1,
      sym_identifier,
  [6024] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(360), 1,
      aux_sym_identifier_token1,
    STATE(87), 1,
      sym_declared_identifier,
  [6034] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(437), 2,
      sym__newline,
      anon_sym_SEMI,
  [6042] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(340), 1,
      aux_sym_identifier_token1,
    STATE(146), 1,
      sym_identifier,
  [6052] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(469), 2,
      sym__dedent,
      aux_sym_identifier_token1,
  [6060] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(463), 2,
      sym__newline,
      anon_sym_SEMI,
  [6068] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(471), 2,
      anon_sym_COMMA,
      anon_sym_RPAREN,
  [6076] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(435), 2,
      anon_sym_COMMA,
      anon_sym_RPAREN,
  [6084] = 3,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(84), 1,
      anon_sym_DQUOTE,
    STATE(255), 1,
      sym_string_literal,
  [6094] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(473), 2,
      sym__newline,
      anon_sym_SEMI,
  [6102] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(475), 2,
      anon_sym_COMMA,
      anon_sym_RPAREN,
  [6110] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(452), 2,
      anon_sym_COMMA,
      anon_sym_RPAREN,
  [6118] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(477), 2,
      sym__newline,
      anon_sym_SEMI,
  [6126] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(479), 1,
      anon_sym_DQUOTE,
  [6133] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(481), 1,
      anon_sym_RBRACK,
  [6140] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(483), 1,
      anon_sym_EQ,
  [6147] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(485), 1,
      sym__newline,
  [6154] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(487), 1,
      sym__newline,
  [6161] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(489), 1,
      aux_sym_identifier_token1,
  [6168] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(491), 1,
      anon_sym_RPAREN,
  [6175] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(493), 1,
      anon_sym_RPAREN,
  [6182] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(495), 1,
      anon_sym_RPAREN,
  [6189] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(497), 1,
      anon_sym_RBRACK,
  [6196] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(499), 1,
      anon_sym_EQ,
  [6203] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(501), 1,
      anon_sym_RBRACK,
  [6210] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(503), 1,
      sym__indent,
  [6217] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(505), 1,
      aux_sym_identifier_token1,
  [6224] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(507), 1,
      sym__newline,
  [6231] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(509), 1,
      anon_sym_RPAREN,
  [6238] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(511), 1,
      anon_sym_RPAREN,
  [6245] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(513), 1,
      anon_sym_DOT,
  [6252] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(515), 1,
      sym__newline,
  [6259] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(517), 1,
      anon_sym_RPAREN,
  [6266] = 2,
    ACTIONS(519), 1,
      sym_comment,
    ACTIONS(521), 1,
      aux_sym_string_literal_token1,
  [6273] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(523), 1,
      sym__indent,
  [6280] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(525), 1,
      aux_sym_identifier_token1,
  [6287] = 2,
    ACTIONS(519), 1,
      sym_comment,
    ACTIONS(527), 1,
      aux_sym_string_literal_token1,
  [6294] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(529), 1,
      sym__newline,
  [6301] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(531), 1,
      anon_sym_EQ,
  [6308] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(533), 1,
      anon_sym_EQ,
  [6315] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(535), 1,
      sym__newline,
  [6322] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(537), 1,
      sym__newline,
  [6329] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(539), 1,
      ts_builtin_sym_end,
  [6336] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(541), 1,
      anon_sym_DQUOTE,
  [6343] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(543), 1,
      anon_sym_DOT,
  [6350] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(545), 1,
      anon_sym_RBRACK,
  [6357] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(547), 1,
      sym__indent,
  [6364] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(549), 1,
      anon_sym_EQ,
  [6371] = 2,
    ACTIONS(3), 1,
      sym_comment,
    ACTIONS(551), 1,
      anon_sym_RBRACK,
};

static const uint32_t ts_small_parse_table_map[] = {
  [SMALL_STATE(2)] = 0,
  [SMALL_STATE(3)] = 62,
  [SMALL_STATE(4)] = 124,
  [SMALL_STATE(5)] = 186,
  [SMALL_STATE(6)] = 245,
  [SMALL_STATE(7)] = 304,
  [SMALL_STATE(8)] = 381,
  [SMALL_STATE(9)] = 457,
  [SMALL_STATE(10)] = 533,
  [SMALL_STATE(11)] = 609,
  [SMALL_STATE(12)] = 685,
  [SMALL_STATE(13)] = 756,
  [SMALL_STATE(14)] = 826,
  [SMALL_STATE(15)] = 896,
  [SMALL_STATE(16)] = 966,
  [SMALL_STATE(17)] = 1036,
  [SMALL_STATE(18)] = 1106,
  [SMALL_STATE(19)] = 1176,
  [SMALL_STATE(20)] = 1246,
  [SMALL_STATE(21)] = 1316,
  [SMALL_STATE(22)] = 1386,
  [SMALL_STATE(23)] = 1456,
  [SMALL_STATE(24)] = 1526,
  [SMALL_STATE(25)] = 1596,
  [SMALL_STATE(26)] = 1659,
  [SMALL_STATE(27)] = 1722,
  [SMALL_STATE(28)] = 1785,
  [SMALL_STATE(29)] = 1848,
  [SMALL_STATE(30)] = 1911,
  [SMALL_STATE(31)] = 1974,
  [SMALL_STATE(32)] = 2037,
  [SMALL_STATE(33)] = 2100,
  [SMALL_STATE(34)] = 2163,
  [SMALL_STATE(35)] = 2226,
  [SMALL_STATE(36)] = 2261,
  [SMALL_STATE(37)] = 2296,
  [SMALL_STATE(38)] = 2331,
  [SMALL_STATE(39)] = 2366,
  [SMALL_STATE(40)] = 2401,
  [SMALL_STATE(41)] = 2436,
  [SMALL_STATE(42)] = 2468,
  [SMALL_STATE(43)] = 2500,
  [SMALL_STATE(44)] = 2532,
  [SMALL_STATE(45)] = 2564,
  [SMALL_STATE(46)] = 2596,
  [SMALL_STATE(47)] = 2628,
  [SMALL_STATE(48)] = 2660,
  [SMALL_STATE(49)] = 2692,
  [SMALL_STATE(50)] = 2724,
  [SMALL_STATE(51)] = 2756,
  [SMALL_STATE(52)] = 2788,
  [SMALL_STATE(53)] = 2820,
  [SMALL_STATE(54)] = 2852,
  [SMALL_STATE(55)] = 2884,
  [SMALL_STATE(56)] = 2916,
  [SMALL_STATE(57)] = 2948,
  [SMALL_STATE(58)] = 2980,
  [SMALL_STATE(59)] = 3017,
  [SMALL_STATE(60)] = 3054,
  [SMALL_STATE(61)] = 3091,
  [SMALL_STATE(62)] = 3128,
  [SMALL_STATE(63)] = 3165,
  [SMALL_STATE(64)] = 3202,
  [SMALL_STATE(65)] = 3239,
  [SMALL_STATE(66)] = 3276,
  [SMALL_STATE(67)] = 3313,
  [SMALL_STATE(68)] = 3350,
  [SMALL_STATE(69)] = 3387,
  [SMALL_STATE(70)] = 3424,
  [SMALL_STATE(71)] = 3458,
  [SMALL_STATE(72)] = 3496,
  [SMALL_STATE(73)] = 3534,
  [SMALL_STATE(74)] = 3562,
  [SMALL_STATE(75)] = 3588,
  [SMALL_STATE(76)] = 3611,
  [SMALL_STATE(77)] = 3636,
  [SMALL_STATE(78)] = 3663,
  [SMALL_STATE(79)] = 3689,
  [SMALL_STATE(80)] = 3715,
  [SMALL_STATE(81)] = 3745,
  [SMALL_STATE(82)] = 3773,
  [SMALL_STATE(83)] = 3798,
  [SMALL_STATE(84)] = 3819,
  [SMALL_STATE(85)] = 3844,
  [SMALL_STATE(86)] = 3871,
  [SMALL_STATE(87)] = 3893,
  [SMALL_STATE(88)] = 3913,
  [SMALL_STATE(89)] = 3933,
  [SMALL_STATE(90)] = 3952,
  [SMALL_STATE(91)] = 3973,
  [SMALL_STATE(92)] = 3992,
  [SMALL_STATE(93)] = 4019,
  [SMALL_STATE(94)] = 4038,
  [SMALL_STATE(95)] = 4063,
  [SMALL_STATE(96)] = 4084,
  [SMALL_STATE(97)] = 4107,
  [SMALL_STATE(98)] = 4126,
  [SMALL_STATE(99)] = 4145,
  [SMALL_STATE(100)] = 4164,
  [SMALL_STATE(101)] = 4183,
  [SMALL_STATE(102)] = 4202,
  [SMALL_STATE(103)] = 4223,
  [SMALL_STATE(104)] = 4244,
  [SMALL_STATE(105)] = 4263,
  [SMALL_STATE(106)] = 4282,
  [SMALL_STATE(107)] = 4304,
  [SMALL_STATE(108)] = 4328,
  [SMALL_STATE(109)] = 4346,
  [SMALL_STATE(110)] = 4372,
  [SMALL_STATE(111)] = 4390,
  [SMALL_STATE(112)] = 4408,
  [SMALL_STATE(113)] = 4428,
  [SMALL_STATE(114)] = 4446,
  [SMALL_STATE(115)] = 4466,
  [SMALL_STATE(116)] = 4484,
  [SMALL_STATE(117)] = 4502,
  [SMALL_STATE(118)] = 4522,
  [SMALL_STATE(119)] = 4540,
  [SMALL_STATE(120)] = 4558,
  [SMALL_STATE(121)] = 4576,
  [SMALL_STATE(122)] = 4594,
  [SMALL_STATE(123)] = 4607,
  [SMALL_STATE(124)] = 4623,
  [SMALL_STATE(125)] = 4645,
  [SMALL_STATE(126)] = 4661,
  [SMALL_STATE(127)] = 4683,
  [SMALL_STATE(128)] = 4705,
  [SMALL_STATE(129)] = 4727,
  [SMALL_STATE(130)] = 4749,
  [SMALL_STATE(131)] = 4771,
  [SMALL_STATE(132)] = 4793,
  [SMALL_STATE(133)] = 4815,
  [SMALL_STATE(134)] = 4831,
  [SMALL_STATE(135)] = 4847,
  [SMALL_STATE(136)] = 4869,
  [SMALL_STATE(137)] = 4885,
  [SMALL_STATE(138)] = 4907,
  [SMALL_STATE(139)] = 4923,
  [SMALL_STATE(140)] = 4936,
  [SMALL_STATE(141)] = 4949,
  [SMALL_STATE(142)] = 4962,
  [SMALL_STATE(143)] = 4981,
  [SMALL_STATE(144)] = 5000,
  [SMALL_STATE(145)] = 5019,
  [SMALL_STATE(146)] = 5038,
  [SMALL_STATE(147)] = 5057,
  [SMALL_STATE(148)] = 5070,
  [SMALL_STATE(149)] = 5083,
  [SMALL_STATE(150)] = 5096,
  [SMALL_STATE(151)] = 5109,
  [SMALL_STATE(152)] = 5128,
  [SMALL_STATE(153)] = 5141,
  [SMALL_STATE(154)] = 5160,
  [SMALL_STATE(155)] = 5179,
  [SMALL_STATE(156)] = 5192,
  [SMALL_STATE(157)] = 5211,
  [SMALL_STATE(158)] = 5230,
  [SMALL_STATE(159)] = 5249,
  [SMALL_STATE(160)] = 5268,
  [SMALL_STATE(161)] = 5281,
  [SMALL_STATE(162)] = 5300,
  [SMALL_STATE(163)] = 5313,
  [SMALL_STATE(164)] = 5332,
  [SMALL_STATE(165)] = 5345,
  [SMALL_STATE(166)] = 5358,
  [SMALL_STATE(167)] = 5371,
  [SMALL_STATE(168)] = 5390,
  [SMALL_STATE(169)] = 5403,
  [SMALL_STATE(170)] = 5416,
  [SMALL_STATE(171)] = 5429,
  [SMALL_STATE(172)] = 5442,
  [SMALL_STATE(173)] = 5455,
  [SMALL_STATE(174)] = 5468,
  [SMALL_STATE(175)] = 5484,
  [SMALL_STATE(176)] = 5498,
  [SMALL_STATE(177)] = 5514,
  [SMALL_STATE(178)] = 5528,
  [SMALL_STATE(179)] = 5542,
  [SMALL_STATE(180)] = 5558,
  [SMALL_STATE(181)] = 5567,
  [SMALL_STATE(182)] = 5580,
  [SMALL_STATE(183)] = 5593,
  [SMALL_STATE(184)] = 5606,
  [SMALL_STATE(185)] = 5619,
  [SMALL_STATE(186)] = 5632,
  [SMALL_STATE(187)] = 5641,
  [SMALL_STATE(188)] = 5654,
  [SMALL_STATE(189)] = 5667,
  [SMALL_STATE(190)] = 5680,
  [SMALL_STATE(191)] = 5693,
  [SMALL_STATE(192)] = 5706,
  [SMALL_STATE(193)] = 5719,
  [SMALL_STATE(194)] = 5732,
  [SMALL_STATE(195)] = 5745,
  [SMALL_STATE(196)] = 5758,
  [SMALL_STATE(197)] = 5771,
  [SMALL_STATE(198)] = 5784,
  [SMALL_STATE(199)] = 5793,
  [SMALL_STATE(200)] = 5802,
  [SMALL_STATE(201)] = 5815,
  [SMALL_STATE(202)] = 5828,
  [SMALL_STATE(203)] = 5841,
  [SMALL_STATE(204)] = 5854,
  [SMALL_STATE(205)] = 5865,
  [SMALL_STATE(206)] = 5878,
  [SMALL_STATE(207)] = 5891,
  [SMALL_STATE(208)] = 5900,
  [SMALL_STATE(209)] = 5913,
  [SMALL_STATE(210)] = 5926,
  [SMALL_STATE(211)] = 5939,
  [SMALL_STATE(212)] = 5952,
  [SMALL_STATE(213)] = 5965,
  [SMALL_STATE(214)] = 5978,
  [SMALL_STATE(215)] = 5986,
  [SMALL_STATE(216)] = 5994,
  [SMALL_STATE(217)] = 6004,
  [SMALL_STATE(218)] = 6014,
  [SMALL_STATE(219)] = 6024,
  [SMALL_STATE(220)] = 6034,
  [SMALL_STATE(221)] = 6042,
  [SMALL_STATE(222)] = 6052,
  [SMALL_STATE(223)] = 6060,
  [SMALL_STATE(224)] = 6068,
  [SMALL_STATE(225)] = 6076,
  [SMALL_STATE(226)] = 6084,
  [SMALL_STATE(227)] = 6094,
  [SMALL_STATE(228)] = 6102,
  [SMALL_STATE(229)] = 6110,
  [SMALL_STATE(230)] = 6118,
  [SMALL_STATE(231)] = 6126,
  [SMALL_STATE(232)] = 6133,
  [SMALL_STATE(233)] = 6140,
  [SMALL_STATE(234)] = 6147,
  [SMALL_STATE(235)] = 6154,
  [SMALL_STATE(236)] = 6161,
  [SMALL_STATE(237)] = 6168,
  [SMALL_STATE(238)] = 6175,
  [SMALL_STATE(239)] = 6182,
  [SMALL_STATE(240)] = 6189,
  [SMALL_STATE(241)] = 6196,
  [SMALL_STATE(242)] = 6203,
  [SMALL_STATE(243)] = 6210,
  [SMALL_STATE(244)] = 6217,
  [SMALL_STATE(245)] = 6224,
  [SMALL_STATE(246)] = 6231,
  [SMALL_STATE(247)] = 6238,
  [SMALL_STATE(248)] = 6245,
  [SMALL_STATE(249)] = 6252,
  [SMALL_STATE(250)] = 6259,
  [SMALL_STATE(251)] = 6266,
  [SMALL_STATE(252)] = 6273,
  [SMALL_STATE(253)] = 6280,
  [SMALL_STATE(254)] = 6287,
  [SMALL_STATE(255)] = 6294,
  [SMALL_STATE(256)] = 6301,
  [SMALL_STATE(257)] = 6308,
  [SMALL_STATE(258)] = 6315,
  [SMALL_STATE(259)] = 6322,
  [SMALL_STATE(260)] = 6329,
  [SMALL_STATE(261)] = 6336,
  [SMALL_STATE(262)] = 6343,
  [SMALL_STATE(263)] = 6350,
  [SMALL_STATE(264)] = 6357,
  [SMALL_STATE(265)] = 6364,
  [SMALL_STATE(266)] = 6371,
};

static const TSParseActionEntry ts_parse_actions[] = {
  [0] = {.entry = {.count = 0, .reusable = false}},
  [1] = {.entry = {.count = 1, .reusable = false}}, RECOVER(),
  [3] = {.entry = {.count = 1, .reusable = true}}, SHIFT_EXTRA(),
  [5] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_description, 0),
  [7] = {.entry = {.count = 1, .reusable = false}}, SHIFT(122),
  [9] = {.entry = {.count = 1, .reusable = false}}, SHIFT(156),
  [11] = {.entry = {.count = 1, .reusable = false}}, SHIFT(185),
  [13] = {.entry = {.count = 1, .reusable = false}}, SHIFT(221),
  [15] = {.entry = {.count = 1, .reusable = false}}, SHIFT(256),
  [17] = {.entry = {.count = 1, .reusable = true}}, SHIFT(256),
  [19] = {.entry = {.count = 1, .reusable = false}}, SHIFT(182),
  [21] = {.entry = {.count = 1, .reusable = false}}, SHIFT(218),
  [23] = {.entry = {.count = 1, .reusable = true}}, SHIFT(155),
  [25] = {.entry = {.count = 1, .reusable = true}}, SHIFT(50),
  [27] = {.entry = {.count = 2, .reusable = false}}, REDUCE(aux_sym_element_body_repeat1, 2), SHIFT_REPEAT(122),
  [30] = {.entry = {.count = 2, .reusable = false}}, REDUCE(aux_sym_element_body_repeat1, 2), SHIFT_REPEAT(256),
  [33] = {.entry = {.count = 2, .reusable = true}}, REDUCE(aux_sym_element_body_repeat1, 2), SHIFT_REPEAT(256),
  [36] = {.entry = {.count = 2, .reusable = false}}, REDUCE(aux_sym_element_body_repeat1, 2), SHIFT_REPEAT(182),
  [39] = {.entry = {.count = 2, .reusable = false}}, REDUCE(aux_sym_element_body_repeat1, 2), SHIFT_REPEAT(218),
  [42] = {.entry = {.count = 1, .reusable = true}}, REDUCE(aux_sym_element_body_repeat1, 2),
  [44] = {.entry = {.count = 1, .reusable = false}}, SHIFT(75),
  [46] = {.entry = {.count = 1, .reusable = true}}, SHIFT(79),
  [48] = {.entry = {.count = 1, .reusable = true}}, SHIFT(19),
  [50] = {.entry = {.count = 1, .reusable = true}}, SHIFT(227),
  [52] = {.entry = {.count = 1, .reusable = true}}, SHIFT(11),
  [54] = {.entry = {.count = 1, .reusable = true}}, SHIFT(254),
  [56] = {.entry = {.count = 1, .reusable = false}}, SHIFT(74),
  [58] = {.entry = {.count = 1, .reusable = true}}, SHIFT(74),
  [60] = {.entry = {.count = 1, .reusable = true}}, SHIFT(73),
  [62] = {.entry = {.count = 1, .reusable = true}}, SHIFT(100),
  [64] = {.entry = {.count = 1, .reusable = true}}, SHIFT(34),
  [66] = {.entry = {.count = 1, .reusable = false}}, SHIFT(100),
  [68] = {.entry = {.count = 1, .reusable = true}}, SHIFT(110),
  [70] = {.entry = {.count = 1, .reusable = true}}, SHIFT(220),
  [72] = {.entry = {.count = 1, .reusable = true}}, SHIFT(98),
  [74] = {.entry = {.count = 1, .reusable = true}}, SHIFT(198),
  [76] = {.entry = {.count = 1, .reusable = false}}, SHIFT(83),
  [78] = {.entry = {.count = 1, .reusable = true}}, SHIFT(82),
  [80] = {.entry = {.count = 1, .reusable = true}}, SHIFT(21),
  [82] = {.entry = {.count = 1, .reusable = true}}, SHIFT(9),
  [84] = {.entry = {.count = 1, .reusable = true}}, SHIFT(251),
  [86] = {.entry = {.count = 1, .reusable = false}}, SHIFT(76),
  [88] = {.entry = {.count = 1, .reusable = true}}, SHIFT(76),
  [90] = {.entry = {.count = 1, .reusable = true}}, SHIFT(77),
  [92] = {.entry = {.count = 1, .reusable = true}}, SHIFT(116),
  [94] = {.entry = {.count = 1, .reusable = true}}, SHIFT(27),
  [96] = {.entry = {.count = 1, .reusable = false}}, SHIFT(116),
  [98] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_single_line_type_definition, 4),
  [100] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_single_line_type_definition, 4),
  [102] = {.entry = {.count = 1, .reusable = true}}, SHIFT(5),
  [104] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_single_line_instantiation, 4),
  [106] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_single_line_instantiation, 4),
  [108] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_single_line_type_definition, 5),
  [110] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_single_line_type_definition, 5),
  [112] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_single_line_instantiation, 5),
  [114] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_single_line_instantiation, 5),
  [116] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_single_line_instantiation, 3),
  [118] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_single_line_instantiation, 3),
  [120] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_single_line_type_definition, 6),
  [122] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_single_line_type_definition, 6),
  [124] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_multi_constant_definition, 5),
  [126] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_multi_constant_definition, 5),
  [128] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_multi_line_instantiation, 3),
  [130] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_multi_line_instantiation, 3),
  [132] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_multi_line_type_definition, 4),
  [134] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_multi_line_type_definition, 4),
  [136] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_single_line_type_definition, 7),
  [138] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_single_line_type_definition, 7),
  [140] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_single_line_instantiation, 6),
  [142] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_single_line_instantiation, 6),
  [144] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_multi_line_type_definition, 6),
  [146] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_multi_line_type_definition, 6),
  [148] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_property_assignments, 3),
  [150] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_property_assignments, 3),
  [152] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_element_body, 4),
  [154] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_element_body, 4),
  [156] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_multi_line_instantiation, 5),
  [158] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_multi_line_instantiation, 5),
  [160] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_multi_line_instantiation, 4),
  [162] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_multi_line_instantiation, 4),
  [164] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_property_assignments, 2),
  [166] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_property_assignments, 2),
  [168] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_single_constant_definition, 5),
  [170] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_single_constant_definition, 5),
  [172] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_multi_line_type_definition, 5),
  [174] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_multi_line_type_definition, 5),
  [176] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_description, 1),
  [178] = {.entry = {.count = 1, .reusable = true}}, REDUCE(aux_sym_description_repeat1, 2),
  [180] = {.entry = {.count = 2, .reusable = false}}, REDUCE(aux_sym_description_repeat1, 2), SHIFT_REPEAT(122),
  [183] = {.entry = {.count = 2, .reusable = false}}, REDUCE(aux_sym_description_repeat1, 2), SHIFT_REPEAT(156),
  [186] = {.entry = {.count = 2, .reusable = false}}, REDUCE(aux_sym_description_repeat1, 2), SHIFT_REPEAT(185),
  [189] = {.entry = {.count = 2, .reusable = false}}, REDUCE(aux_sym_description_repeat1, 2), SHIFT_REPEAT(221),
  [192] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_primary_expression, 1),
  [194] = {.entry = {.count = 1, .reusable = true}}, SHIFT(89),
  [196] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_primary_expression, 1),
  [198] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym__zero_or_decimal_literal, 1),
  [200] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym__integer_literal, 1),
  [202] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym__integer_literal, 1),
  [204] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_declared_identifier, 1),
  [206] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_declared_identifier, 1),
  [208] = {.entry = {.count = 1, .reusable = true}}, SHIFT(121),
  [210] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_float_literal, 2),
  [212] = {.entry = {.count = 1, .reusable = true}}, SHIFT(97),
  [214] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_float_literal, 2),
  [216] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_float_literal, 1),
  [218] = {.entry = {.count = 1, .reusable = true}}, SHIFT(101),
  [220] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_float_literal, 1),
  [222] = {.entry = {.count = 1, .reusable = true}}, SHIFT(217),
  [224] = {.entry = {.count = 1, .reusable = true}}, SHIFT(16),
  [226] = {.entry = {.count = 1, .reusable = true}}, SHIFT(10),
  [228] = {.entry = {.count = 1, .reusable = true}}, SHIFT(23),
  [230] = {.entry = {.count = 1, .reusable = true}}, SHIFT(115),
  [232] = {.entry = {.count = 1, .reusable = true}}, SHIFT(118),
  [234] = {.entry = {.count = 1, .reusable = true}}, SHIFT(219),
  [236] = {.entry = {.count = 1, .reusable = true}}, SHIFT(8),
  [238] = {.entry = {.count = 1, .reusable = true}}, SHIFT(22),
  [240] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_qualified_identifier, 3),
  [242] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_qualified_identifier, 3),
  [244] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_time_unit, 1),
  [246] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_time_unit, 1),
  [248] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_binary_operation, 3),
  [250] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_binary_operation, 3),
  [252] = {.entry = {.count = 1, .reusable = true}}, SHIFT(28),
  [254] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_call, 4),
  [256] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_call, 4),
  [258] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym__expression, 1),
  [260] = {.entry = {.count = 1, .reusable = true}}, SHIFT(30),
  [262] = {.entry = {.count = 1, .reusable = false}}, SHIFT(29),
  [264] = {.entry = {.count = 1, .reusable = true}}, SHIFT(29),
  [266] = {.entry = {.count = 1, .reusable = true}}, SHIFT(26),
  [268] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_subscript, 4),
  [270] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_subscript, 4),
  [272] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_float_literal, 3),
  [274] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_float_literal, 3),
  [276] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_call, 3),
  [278] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_call, 3),
  [280] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_parenthesized_expression, 3),
  [282] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_parenthesized_expression, 3),
  [284] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_unary_operation, 2),
  [286] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_unary_operation, 2),
  [288] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_time_literal, 2),
  [290] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_time_literal, 2),
  [292] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_string_literal, 3),
  [294] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_string_literal, 3),
  [296] = {.entry = {.count = 1, .reusable = false}}, SHIFT(33),
  [298] = {.entry = {.count = 1, .reusable = true}}, SHIFT(33),
  [300] = {.entry = {.count = 1, .reusable = true}}, SHIFT(32),
  [302] = {.entry = {.count = 1, .reusable = true}}, SHIFT(25),
  [304] = {.entry = {.count = 1, .reusable = true}}, SHIFT(31),
  [306] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_identifier, 1),
  [308] = {.entry = {.count = 1, .reusable = true}}, SHIFT(6),
  [310] = {.entry = {.count = 1, .reusable = true}}, SHIFT(7),
  [312] = {.entry = {.count = 1, .reusable = true}}, SHIFT(68),
  [314] = {.entry = {.count = 1, .reusable = true}}, SHIFT(134),
  [316] = {.entry = {.count = 1, .reusable = true}}, SHIFT(62),
  [318] = {.entry = {.count = 1, .reusable = true}}, SHIFT(37),
  [320] = {.entry = {.count = 1, .reusable = true}}, SHIFT(65),
  [322] = {.entry = {.count = 1, .reusable = true}}, SHIFT(125),
  [324] = {.entry = {.count = 1, .reusable = true}}, SHIFT(58),
  [326] = {.entry = {.count = 1, .reusable = true}}, SHIFT(35),
  [328] = {.entry = {.count = 1, .reusable = true}}, SHIFT(69),
  [330] = {.entry = {.count = 1, .reusable = true}}, SHIFT(39),
  [332] = {.entry = {.count = 1, .reusable = true}}, SHIFT(66),
  [334] = {.entry = {.count = 1, .reusable = true}}, SHIFT(123),
  [336] = {.entry = {.count = 1, .reusable = true}}, SHIFT(63),
  [338] = {.entry = {.count = 1, .reusable = true}}, SHIFT(36),
  [340] = {.entry = {.count = 1, .reusable = true}}, SHIFT(122),
  [342] = {.entry = {.count = 1, .reusable = true}}, SHIFT(171),
  [344] = {.entry = {.count = 2, .reusable = true}}, REDUCE(aux_sym_multi_import_statement_repeat1, 2), SHIFT_REPEAT(122),
  [347] = {.entry = {.count = 2, .reusable = true}}, REDUCE(aux_sym_multi_import_statement_repeat1, 2), SHIFT_REPEAT(251),
  [350] = {.entry = {.count = 1, .reusable = true}}, REDUCE(aux_sym_multi_import_statement_repeat1, 2),
  [352] = {.entry = {.count = 1, .reusable = true}}, SHIFT(64),
  [354] = {.entry = {.count = 1, .reusable = true}}, SHIFT(133),
  [356] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_single_import_statement, 3),
  [358] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_single_import_statement, 3),
  [360] = {.entry = {.count = 1, .reusable = true}}, SHIFT(83),
  [362] = {.entry = {.count = 1, .reusable = true}}, SHIFT(17),
  [364] = {.entry = {.count = 1, .reusable = true}}, SHIFT(161),
  [366] = {.entry = {.count = 1, .reusable = true}}, SHIFT(252),
  [368] = {.entry = {.count = 1, .reusable = true}}, SHIFT(236),
  [370] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_single_import_statement, 4),
  [372] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_single_import_statement, 4),
  [374] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_multi_import_statement, 5),
  [376] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_multi_import_statement, 5),
  [378] = {.entry = {.count = 1, .reusable = true}}, SHIFT(160),
  [380] = {.entry = {.count = 1, .reusable = true}}, SHIFT(13),
  [382] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym__expressions, 2),
  [384] = {.entry = {.count = 1, .reusable = true}}, SHIFT(42),
  [386] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym__expressions, 1),
  [388] = {.entry = {.count = 2, .reusable = true}}, REDUCE(aux_sym__expressions_repeat1, 2), SHIFT_REPEAT(13),
  [391] = {.entry = {.count = 1, .reusable = true}}, REDUCE(aux_sym__expressions_repeat1, 2),
  [393] = {.entry = {.count = 2, .reusable = true}}, REDUCE(aux_sym_multi_constant_definition_repeat1, 2), SHIFT_REPEAT(122),
  [396] = {.entry = {.count = 1, .reusable = true}}, REDUCE(aux_sym_multi_constant_definition_repeat1, 2),
  [398] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_expression_list, 3),
  [400] = {.entry = {.count = 1, .reusable = true}}, SHIFT(264),
  [402] = {.entry = {.count = 1, .reusable = true}}, SHIFT(184),
  [404] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym__parameters, 2),
  [406] = {.entry = {.count = 1, .reusable = true}}, SHIFT(243),
  [408] = {.entry = {.count = 1, .reusable = true}}, SHIFT(59),
  [410] = {.entry = {.count = 1, .reusable = true}}, SHIFT(136),
  [412] = {.entry = {.count = 1, .reusable = true}}, SHIFT(70),
  [414] = {.entry = {.count = 1, .reusable = true}}, SHIFT(149),
  [416] = {.entry = {.count = 1, .reusable = true}}, SHIFT(49),
  [418] = {.entry = {.count = 1, .reusable = true}}, SHIFT(60),
  [420] = {.entry = {.count = 1, .reusable = true}}, SHIFT(40),
  [422] = {.entry = {.count = 1, .reusable = true}}, SHIFT(12),
  [424] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym__arguments, 1),
  [426] = {.entry = {.count = 1, .reusable = true}}, SHIFT(61),
  [428] = {.entry = {.count = 1, .reusable = true}}, SHIFT(38),
  [430] = {.entry = {.count = 1, .reusable = true}}, SHIFT(55),
  [432] = {.entry = {.count = 2, .reusable = true}}, REDUCE(aux_sym__parameters_repeat1, 2), SHIFT_REPEAT(184),
  [435] = {.entry = {.count = 1, .reusable = true}}, REDUCE(aux_sym__parameters_repeat1, 2),
  [437] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_expression_list, 2),
  [439] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym__arguments, 2),
  [441] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym__parameters, 1),
  [443] = {.entry = {.count = 1, .reusable = true}}, SHIFT(15),
  [445] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym__parameter, 1),
  [447] = {.entry = {.count = 1, .reusable = true}}, REDUCE(aux_sym_multi_import_statement_repeat1, 3),
  [449] = {.entry = {.count = 2, .reusable = true}}, REDUCE(aux_sym__arguments_repeat1, 2), SHIFT_REPEAT(12),
  [452] = {.entry = {.count = 1, .reusable = true}}, REDUCE(aux_sym__arguments_repeat1, 2),
  [454] = {.entry = {.count = 1, .reusable = true}}, SHIFT(162),
  [456] = {.entry = {.count = 1, .reusable = true}}, SHIFT(67),
  [458] = {.entry = {.count = 1, .reusable = true}}, SHIFT(138),
  [460] = {.entry = {.count = 2, .reusable = true}}, REDUCE(aux_sym_property_assignments_repeat1, 2), SHIFT_REPEAT(70),
  [463] = {.entry = {.count = 1, .reusable = true}}, REDUCE(aux_sym_property_assignments_repeat1, 2),
  [465] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym__property_assignment, 3),
  [467] = {.entry = {.count = 1, .reusable = true}}, SHIFT(75),
  [469] = {.entry = {.count = 1, .reusable = true}}, REDUCE(aux_sym_multi_constant_definition_repeat1, 4),
  [471] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym__parameter, 3),
  [473] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_argument_list, 2),
  [475] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym__argument, 3),
  [477] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_argument_list, 3),
  [479] = {.entry = {.count = 1, .reusable = true}}, SHIFT(108),
  [481] = {.entry = {.count = 1, .reusable = true}}, SHIFT(111),
  [483] = {.entry = {.count = 1, .reusable = true}}, SHIFT(18),
  [485] = {.entry = {.count = 1, .reusable = true}}, SHIFT(152),
  [487] = {.entry = {.count = 1, .reusable = true}}, SHIFT(56),
  [489] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_parameter_list, 2),
  [491] = {.entry = {.count = 1, .reusable = true}}, SHIFT(253),
  [493] = {.entry = {.count = 1, .reusable = true}}, SHIFT(99),
  [495] = {.entry = {.count = 1, .reusable = true}}, SHIFT(113),
  [497] = {.entry = {.count = 1, .reusable = true}}, SHIFT(215),
  [499] = {.entry = {.count = 1, .reusable = true}}, SHIFT(24),
  [501] = {.entry = {.count = 1, .reusable = true}}, SHIFT(180),
  [503] = {.entry = {.count = 1, .reusable = true}}, SHIFT(212),
  [505] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_array_marker, 3),
  [507] = {.entry = {.count = 1, .reusable = true}}, SHIFT(140),
  [509] = {.entry = {.count = 1, .reusable = true}}, SHIFT(91),
  [511] = {.entry = {.count = 1, .reusable = true}}, SHIFT(120),
  [513] = {.entry = {.count = 1, .reusable = true}}, SHIFT(84),
  [515] = {.entry = {.count = 1, .reusable = true}}, SHIFT(186),
  [517] = {.entry = {.count = 1, .reusable = true}}, SHIFT(230),
  [519] = {.entry = {.count = 1, .reusable = false}}, SHIFT_EXTRA(),
  [521] = {.entry = {.count = 1, .reusable = false}}, SHIFT(231),
  [523] = {.entry = {.count = 1, .reusable = true}}, SHIFT(157),
  [525] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_parameter_list, 3),
  [527] = {.entry = {.count = 1, .reusable = false}}, SHIFT(261),
  [529] = {.entry = {.count = 1, .reusable = true}}, SHIFT(168),
  [531] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_property_identifier, 1),
  [533] = {.entry = {.count = 1, .reusable = true}}, SHIFT(14),
  [535] = {.entry = {.count = 1, .reusable = true}}, SHIFT(222),
  [537] = {.entry = {.count = 1, .reusable = true}}, SHIFT(207),
  [539] = {.entry = {.count = 1, .reusable = true}},  ACCEPT_INPUT(),
  [541] = {.entry = {.count = 1, .reusable = true}}, SHIFT(105),
  [543] = {.entry = {.count = 1, .reusable = true}}, SHIFT(78),
  [545] = {.entry = {.count = 1, .reusable = true}}, SHIFT(244),
  [547] = {.entry = {.count = 1, .reusable = true}}, SHIFT(188),
  [549] = {.entry = {.count = 1, .reusable = true}}, SHIFT(20),
  [551] = {.entry = {.count = 1, .reusable = true}}, SHIFT(93),
};

#ifdef __cplusplus
extern "C" {
#endif
void *tree_sitter_fbdl_external_scanner_create(void);
void tree_sitter_fbdl_external_scanner_destroy(void *);
bool tree_sitter_fbdl_external_scanner_scan(void *, TSLexer *, const bool *);
unsigned tree_sitter_fbdl_external_scanner_serialize(void *, char *);
void tree_sitter_fbdl_external_scanner_deserialize(void *, const char *, unsigned);

#ifdef _WIN32
#define extern __declspec(dllexport)
#endif

extern const TSLanguage *tree_sitter_fbdl(void) {
  static const TSLanguage language = {
    .version = LANGUAGE_VERSION,
    .symbol_count = SYMBOL_COUNT,
    .alias_count = ALIAS_COUNT,
    .token_count = TOKEN_COUNT,
    .external_token_count = EXTERNAL_TOKEN_COUNT,
    .state_count = STATE_COUNT,
    .large_state_count = LARGE_STATE_COUNT,
    .production_id_count = PRODUCTION_ID_COUNT,
    .field_count = FIELD_COUNT,
    .max_alias_sequence_length = MAX_ALIAS_SEQUENCE_LENGTH,
    .parse_table = &ts_parse_table[0][0],
    .small_parse_table = ts_small_parse_table,
    .small_parse_table_map = ts_small_parse_table_map,
    .parse_actions = ts_parse_actions,
    .symbol_names = ts_symbol_names,
    .symbol_metadata = ts_symbol_metadata,
    .public_symbol_map = ts_symbol_map,
    .alias_map = ts_non_terminal_alias_map,
    .alias_sequences = &ts_alias_sequences[0][0],
    .lex_modes = ts_lex_modes,
    .lex_fn = ts_lex,
    .external_scanner = {
      &ts_external_scanner_states[0][0],
      ts_external_scanner_symbol_map,
      tree_sitter_fbdl_external_scanner_create,
      tree_sitter_fbdl_external_scanner_destroy,
      tree_sitter_fbdl_external_scanner_scan,
      tree_sitter_fbdl_external_scanner_serialize,
      tree_sitter_fbdl_external_scanner_deserialize,
    },
  };
  return &language;
}
#ifdef __cplusplus
}
#endif
