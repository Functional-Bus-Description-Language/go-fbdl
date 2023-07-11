package token

// Token is the set of lexical tokens of the FBDL.
type Kind int

const (
	// Special tokens
	INVALID Kind = iota
	COMMENT
	INDENT_INC // Indent increase
	INDENT_DEC // Indent decrease
	NEWLINE
	// EOF - probably there will be no need for EOF token

	literal_start
	IDENT
	BOOL
	INT
	REAL
	STRING
	BIT_STRING
	TIME
	literal_end

	//operator_start
	// Unary operators
	NEG // !

	// Binary arithmetic operators
	ADD // +
	SUB // -
	MUL // *
	DIV // /
	REM // %
	EXP // **

	// Binary comparison operators
	EQ  // ==
	NEQ // !=
	LSS // <
	LEQ // <=
	GTR // >
	GEQ // >=

	// Binary logical operators
	LAND // &&
	LOR  // ||

	// Binary bitwise operators
	SHL // <<
	SHR // >>
	AND // &
	OR  // |
	XOR // ^
	//operator_end

	LPAREN // (
	RPAREN // )

	LBRACK // [
	RBRACK // ]

	LBRACE // {
	RBRACE // }

	COMMA     // ,
	SEMICOLON // ;

	//keyword_start
	BLOCK
	BUS
	CONST
	IMPORT
	IRQ
	MASK
	MEMORY
	PARAM
	PROC
	RETURN
	STATIC
	STREAM
	TYPE
	//keyword_end

	property_start
	ACCESS
	ADD_ENABLE
	ATOMIC
	BYTE_WRITE_ENABLE
	CLEAR
	DELAY
	ENABLE_INIT_VALUE
	ENABLE_RESET_VALUE
	GROUPS
	INIT_VALUE
	IN_TRIGGER
	MASTERS
	OUT_TRIGGER
	RANGE
	READ_LATENCY
	READ_VALUE
	RESET
	RESET_VALUE
	SIZE
	WIDTH
	property_end

	//unused_start
	// Tokens currently not used by the language specification.
	PERIOD // .
	COLON  // :
	//unused_end
)

type Token struct {
	Kind Kind
	Pos  Position
}
