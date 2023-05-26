package idl

// FieldType represents a field of a struct or variant.
type FieldType struct {
	// Name is the name of the field.
	Name string
	// Type is the type of the field.
	Type Type

	// index is the index of the field in de type definition table.
	index int64
}
