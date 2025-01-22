// Package candid is autogenerated by https://github.com/0x51-dev/upeg. DO NOT EDIT.
package candid

import (
	"github.com/0x51-dev/upeg/parser"
	"github.com/0x51-dev/upeg/parser/op"
)

var (
	Prog        = op.Capture{Name: "Prog", Value: op.And{op.Optional{Value: op.And{Ws, Def, op.ZeroOrMore{Value: op.And{';', Ws, Def}}}}, op.Optional{Value: ';'}, Ws, op.Optional{Value: op.And{Ws, Actor, op.ZeroOrMore{Value: op.And{';', Ws, Actor}}}}, op.Optional{Value: ';'}, Ws}}
	Def         = op.Or{Type, Import}
	Type        = op.Capture{Name: "Type", Value: op.And{"type", Sp, Id, Sp, '=', Sp, DataType}}
	Import      = op.Capture{Name: "Import", Value: op.And{"import", Sp, Text}}
	Actor       = op.Capture{Name: "Actor", Value: op.And{"service", op.Optional{Value: Sp}, op.Optional{Value: op.And{Id, Sp}}, ':', Sp, op.Optional{Value: op.And{TupType, Sp, "->", Ws}}, op.Or{ActorType, Id}}}
	ActorType   = op.Capture{Name: "ActorType", Value: op.And{'{', Ws, op.Optional{Value: op.And{MethType, op.ZeroOrMore{Value: op.And{';', Ws, MethType}}, op.Optional{Value: ';'}, Ws}}, '}'}}
	MethType    = op.Capture{Name: "MethType", Value: op.And{Name, op.Optional{Value: Sp}, ':', Ws, op.Or{FuncType, Id}}}
	FuncType    = op.Capture{Name: "FuncType", Value: op.And{TupType, op.Optional{Value: op.And{Sp, "->", Ws, TupType, op.Optional{Value: op.And{Sp, FuncAnn}}}}}}
	FuncAnn     = op.Capture{Name: "FuncAnn", Value: op.Or{"oneway", "query"}}
	TupType     = op.Capture{Name: "TupType", Value: op.Or{op.And{'(', Ws, op.Optional{Value: op.And{ArgType, op.ZeroOrMore{Value: op.And{',', Sp, ArgType}}, op.Optional{Value: op.And{',', Ws}}}}, Ws, ')'}, ArgType}}
	ArgType     = op.Capture{Name: "ArgType", Value: op.And{op.Optional{Value: op.And{Name, op.Optional{Value: Sp}, ':', Sp}}, DataType}}
	FieldType   = op.Capture{Name: "FieldType", Value: op.Or{op.And{op.Optional{Value: op.And{op.Or{Nat, Name}, op.Optional{Value: Sp}, ':', Sp}}, DataType}, Nat, Name}}
	DataType    = op.Or{ConsType, RefType, PrimType, Id}
	PrimType    = op.Capture{Name: "PrimType", Value: op.Or{NumType, "bool", "text", "null", "reserved", "empty"}}
	NumType     = op.Or{"nat8", "nat16", "nat32", "nat64", "nat", "int8", "int16", "int32", "int64", "int", "float32", "float64"}
	ConsType    = op.Or{Blob, Opt, Vec, Record, Variant}
	Blob        = op.Capture{Name: "Blob", Value: "blob"}
	Opt         = op.Capture{Name: "Opt", Value: op.And{"opt", Sp, op.Reference{Name: "DataType"}}}
	Vec         = op.Capture{Name: "Vec", Value: op.And{"vec", Sp, op.Reference{Name: "DataType"}}}
	Record      = op.Capture{Name: "Record", Value: op.And{"record", op.Optional{Value: Sp}, '{', Ws, op.Optional{Value: Fields}, Ws, '}'}}
	Variant     = op.Capture{Name: "Variant", Value: op.And{"variant", Sp, '{', Ws, op.Optional{Value: Fields}, Ws, '}'}}
	Fields      = op.And{op.Reference{Name: "FieldType"}, op.ZeroOrMore{Value: op.And{';', Ws, op.Reference{Name: "FieldType"}}}, op.Optional{Value: ';'}}
	RefType     = op.Or{Func, Service, Principal}
	Func        = op.Capture{Name: "Func", Value: op.And{"func", op.Optional{Value: Sp}, op.Reference{Name: "FuncType"}}}
	Service     = op.Capture{Name: "Service", Value: op.And{"service", Sp, op.Reference{Name: "ActorType"}}}
	Principal   = op.Capture{Name: "Principal", Value: "principal"}
	Name        = op.Or{Id, Text}
	Id          = op.Capture{Name: "Id", Value: op.And{op.Or{Letter, '_'}, op.ZeroOrMore{Value: op.Or{Letter, Digit, '_'}}}}
	Text        = op.Capture{Name: "Text", Value: op.And{rune(0x22), op.ZeroOrMore{Value: Char}, rune(0x22)}}
	Char        = op.Or{Utf, op.And{ESC, op.Repeat{Min: 2, Max: 2, Value: Hex}}, op.And{ESC, Escape}, op.And{"\\u{", HexNum, '}'}}
	Num         = op.And{Digit, op.ZeroOrMore{Value: op.And{op.Optional{Value: '_'}, Digit}}}
	HexNum      = op.And{Hex, op.ZeroOrMore{Value: op.And{op.Optional{Value: '_'}, Hex}}}
	Nat         = op.Capture{Name: "Nat", Value: op.Or{op.And{"0x", HexNum}, Num}}
	Utf         = op.Or{Ascii, UtfEnc}
	UtfEnc      = op.Or{op.And{op.RuneRange{Min: 0xC2, Max: 0xDF}, Utfcont}, op.And{rune(0xE0), op.RuneRange{Min: 0xA0, Max: 0xBF}, Utfcont}, op.And{rune(0xED), op.RuneRange{Min: 0x80, Max: 0x9F}, Utfcont}, op.And{op.RuneRange{Min: 0xE1, Max: 0xEC}, op.Repeat{Min: 2, Max: 2, Value: Utfcont}}, op.And{op.RuneRange{Min: 0xEE, Max: 0xEF}, op.Repeat{Min: 2, Max: 2, Value: Utfcont}}, op.And{rune(0xF0), op.RuneRange{Min: 0x90, Max: 0xBF}, op.Repeat{Min: 2, Max: 2, Value: Utfcont}}, op.And{rune(0xF4), op.RuneRange{Min: 0x80, Max: 0x8F}, op.Repeat{Min: 2, Max: 2, Value: Utfcont}}, op.And{op.RuneRange{Min: 0xF1, Max: 0xF3}, op.Repeat{Min: 3, Max: 3, Value: Utfcont}}}
	Utfcont     = op.RuneRange{Min: 0x80, Max: 0xBF}
	Ascii       = op.Or{op.RuneRange{Min: 0x20, Max: 0x21}, op.RuneRange{Min: 0x23, Max: 0x5B}, op.RuneRange{Min: 0x5D, Max: 0x7E}}
	Escape      = op.Or{'n', 'r', 't', ESC, rune(0x22), rune(0x27)}
	Letter      = op.Or{op.RuneRange{Min: 0x41, Max: 0x5A}, op.RuneRange{Min: 0x61, Max: 0x7A}}
	Digit       = op.RuneRange{Min: 0x30, Max: 0x39}
	Hex         = op.Or{Digit, op.RuneRange{Min: 0x41, Max: 0x46}, op.RuneRange{Min: 0x61, Max: 0x66}}
	CommentText = op.Capture{Name: "CommentText", Value: op.ZeroOrMore{Value: op.Or{Ascii, rune(0x22), rune(0x27), rune(0x60)}}}
	Comment     = op.And{"//", CommentText, Nl}
	Nl          = op.Or{rune(0x0A), rune(0x0D), op.And{rune(0x0D), rune(0x0A)}}
	Ws          = op.ZeroOrMore{Value: op.Or{Sp, rune(0x09), Comment, Nl}}
	Sp          = op.OneOrMore{Value: ' '}
	ESC         = rune(0x5C)
)

func NewParser(input []rune) (*parser.Parser, error) {
	p, err := parser.New(input)
	if err != nil {
		return nil, err
	}
	p.Rules["ActorType"] = ActorType
	p.Rules["FuncType"] = FuncType
	p.Rules["FieldType"] = FieldType
	p.Rules["DataType"] = DataType
	return p, nil
}
