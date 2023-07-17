package token

import "fmt"

// Token is the set of lexical tokens of the FBDL.
type Kind int

const (
	// Special tokens
	INVALID Kind = iota
	COMMENT
	INDENT_INC // Indent increase
	INDENT_DEC // Indent decrease
	NEWLINE
	EOF

	literal_start
	IDENT
	BOOL
	INT
	REAL
	STRING
	BIT_STRING
	TIME
	literal_end

	operator_start
	// Unary operators
	NEG // !

	// Binary arithmetic operators
	ASS // =
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
	operator_end

	LPAREN // (
	RPAREN // )

	LBRACK // [
	RBRACK // ]

	COMMA     // ,
	SEMICOLON // ;

	keyword_start
	CONST
	IMPORT
	TYPE
	functionality_start
	BLOCK
	BUS
	CONFIG
	IRQ
	MASK
	MEMORY
	PARAM
	PROC
	RETURN
	STATIC
	STREAM
	functionality_end
	keyword_end

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
	LBRACE // {
	RBRACE // }
	//unused_end
)

func (k Kind) String() string {
	switch k {
	case INVALID:
		return "invalid"
	case COMMENT:
		return "comment"
	case INDENT_INC:
		return "indent increase"
	case INDENT_DEC:
		return "indent decrease"
	case NEWLINE:
		return "newline"
	case EOF:
		return "EOF"
	case IDENT:
		return "identifier"
	case BOOL:
		return "bool"
	case INT:
		return "integer"
	case REAL:
		return "real"
	case STRING:
		return "string"
	case BIT_STRING:
		return "bit string"
	case TIME:
		return "time"
	case NEG:
		return "!"
	case ASS:
		return "="
	case ADD:
		return "+"
	case SUB:
		return "-"
	case MUL:
		return "*"
	case DIV:
		return "/"
	case REM:
		return "%"
	case EXP:
		return "**"
	case EQ:
		return "=="
	case NEQ:
		return "!="
	case LSS:
		return "<"
	case LEQ:
		return "<="
	case GTR:
		return ">"
	case GEQ:
		return ">="
	case LAND:
		return "&&"
	case LOR:
		return "||"
	case SHL:
		return "<<"
	case SHR:
		return ">>"
	case AND:
		return "&"
	case OR:
		return "|"
	case XOR:
		return "^"
	case LPAREN:
		return "("
	case RPAREN:
		return ")"
	case LBRACK:
		return "["
	case RBRACK:
		return "]"
	case COMMA:
		return ","
	case SEMICOLON:
		return ";"
	case CONST:
		return "const"
	case IMPORT:
		return "import"
	case TYPE:
		return "type"
	case BLOCK:
		return "block"
	case BUS:
		return "bus"
	case CONFIG:
		return "config"
	case IRQ:
		return "irq"
	case MASK:
		return "mask"
	case MEMORY:
		return "memory"
	case PARAM:
		return "param"
	case RETURN:
		return "return"
	case STATIC:
		return "static"
	case STREAM:
		return "stream"
	case ACCESS:
		return "access"
	case ADD_ENABLE:
		return "add-enable"
	case ATOMIC:
		return "atomic"
	case BYTE_WRITE_ENABLE:
		return "byte-write-enable"
	case CLEAR:
		return "clear"
	case DELAY:
		return "delay"
	case ENABLE_INIT_VALUE:
		return "enable-init-value"
	case ENABLE_RESET_VALUE:
		return "enable-reset-value"
	case GROUPS:
		return "groups"
	case INIT_VALUE:
		return "init-value"
	case IN_TRIGGER:
		return "in-trigger"
	case MASTERS:
		return "masters"
	case OUT_TRIGGER:
		return "out-trigger"
	case RANGE:
		return "range"
	case READ_LATENCY:
		return "read-latency"
	case READ_VALUE:
		return "read-value"
	case RESET:
		return "reset"
	case RESET_VALUE:
		return "reset-value"
	case SIZE:
		return "size"
	case WIDTH:
		return "width"
	case PERIOD:
		return "."
	case COLON:
		return ":"
	case LBRACE:
		return "{"
	case RBRACE:
		return "}"
	default:
		panic(fmt.Sprintf("unhandled %d kind", k))
	}
}

func isOperator(k Kind) bool {
	return operator_start < k && k < operator_end
}

func isFunctionality(k Kind) bool {
	return functionality_start < k && k < functionality_end
}
