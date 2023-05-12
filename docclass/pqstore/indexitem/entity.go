package indexitem

import (
	"database/sql"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-batis/sqlutil"
)

const (
	ClassIdFieldIndex      = 0
	NdxIdFieldIndex        = 1
	NameFieldIndex         = 2
	NdxTypeFieldIndex      = 3
	DataTypeFieldIndex     = 4
	FormatFieldIndex       = 5
	ValueFieldIndex        = 6
	SourceFormatFieldIndex = 7
	RequiredFieldIndex     = 8
)

type IndexItemEntity struct {
	ClassId      Max30Text      `db:"class_id"`
	NdxId        int32          `db:"ndx_id"`
	Name         Max80Text      `db:"name"`
	NdxType      sql.NullString `db:"type"`
	DataType     sql.NullString `db:"data_type"`
	Format       sql.NullString `db:"format"`
	Value        Max80Text      `db:"value"`
	SourceFormat sql.NullString `db:"source_format"`
	Required     bool           `db:"required"`
}

type PrimaryKey struct {
	ClassId Max30Text `db:"class_id"`
	NdxId   int32     `db:"ndx_id"`
}

// isLengthRestrictionValid utility function for Max??Text types
func isLengthRestrictionValid(s string, length, minLength, maxLength int) bool {
	if length > 0 && len(s) != length {
		return false
	}

	if minLength > 0 && len(s) < minLength {
		return false
	}

	if maxLength > 0 && len(s) > maxLength {
		return false
	}

	return true
}

/*
 * Max20Text Type Def
 */

type Max20Text string

const (
	Max20TextZero      = ""
	Max20TextLength    = 0
	Max20TextMinLength = 1
	Max20TextMaxLength = 20
)

// IsValid checks if Max105Text of type String is valid
func (t Max20Text) IsValid() bool {
	return isLengthRestrictionValid(t.ToString(), Max20TextLength, Max20TextMinLength, Max20TextMaxLength)
}

// ToString method for easy conversion
func (t Max20Text) ToString() string {
	return string(t)
}

// ToMax20Text  method for easy conversion with application of restrictions
func ToMax20Text(i interface{}) (Max20Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max20TextLength, Max20TextMinLength, Max20TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max20Text", s)
	}

	return Max20Text(s), nil
}

// MustToMax20Text  method for easy conversion with application of restrictions. Panics on error.
func MustToMax20Text(s interface{}) Max20Text {
	v, err := ToMax20Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max0Text Type Def
 */

type Max0Text string

const (
	Max0TextZero      = ""
	Max0TextLength    = 0
	Max0TextMinLength = 1
	Max0TextMaxLength = 0
)

// IsValid checks if Max105Text of type String is valid
func (t Max0Text) IsValid() bool {
	return isLengthRestrictionValid(t.ToString(), Max0TextLength, Max0TextMinLength, Max0TextMaxLength)
}

// ToString method for easy conversion
func (t Max0Text) ToString() string {
	return string(t)
}

// ToMax0Text  method for easy conversion with application of restrictions
func ToMax0Text(i interface{}) (Max0Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max0TextLength, Max0TextMinLength, Max0TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max0Text", s)
	}

	return Max0Text(s), nil
}

// MustToMax0Text  method for easy conversion with application of restrictions. Panics on error.
func MustToMax0Text(s interface{}) Max0Text {
	v, err := ToMax0Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max30Text Type Def
 */

type Max30Text string

const (
	Max30TextZero      = ""
	Max30TextLength    = 0
	Max30TextMinLength = 1
	Max30TextMaxLength = 30
)

// IsValid checks if Max105Text of type String is valid
func (t Max30Text) IsValid() bool {
	return isLengthRestrictionValid(t.ToString(), Max30TextLength, Max30TextMinLength, Max30TextMaxLength)
}

// ToString method for easy conversion
func (t Max30Text) ToString() string {
	return string(t)
}

// ToMax30Text  method for easy conversion with application of restrictions
func ToMax30Text(i interface{}) (Max30Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max30TextLength, Max30TextMinLength, Max30TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max30Text", s)
	}

	return Max30Text(s), nil
}

// MustToMax30Text  method for easy conversion with application of restrictions. Panics on error.
func MustToMax30Text(s interface{}) Max30Text {
	v, err := ToMax30Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max80Text Type Def
 */

type Max80Text string

const (
	Max80TextZero      = ""
	Max80TextLength    = 0
	Max80TextMinLength = 1
	Max80TextMaxLength = 80
)

// IsValid checks if Max105Text of type String is valid
func (t Max80Text) IsValid() bool {
	return isLengthRestrictionValid(t.ToString(), Max80TextLength, Max80TextMinLength, Max80TextMaxLength)
}

// ToString method for easy conversion
func (t Max80Text) ToString() string {
	return string(t)
}

// ToMax80Text  method for easy conversion with application of restrictions
func ToMax80Text(i interface{}) (Max80Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max80TextLength, Max80TextMinLength, Max80TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max80Text", s)
	}

	return Max80Text(s), nil
}

// MustToMax80Text  method for easy conversion with application of restrictions. Panics on error.
func MustToMax80Text(s interface{}) Max80Text {
	v, err := ToMax80Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max10Text Type Def
 */

type Max10Text string

const (
	Max10TextZero      = ""
	Max10TextLength    = 0
	Max10TextMinLength = 1
	Max10TextMaxLength = 10
)

// IsValid checks if Max105Text of type String is valid
func (t Max10Text) IsValid() bool {
	return isLengthRestrictionValid(t.ToString(), Max10TextLength, Max10TextMinLength, Max10TextMaxLength)
}

// ToString method for easy conversion
func (t Max10Text) ToString() string {
	return string(t)
}

// ToMax10Text  method for easy conversion with application of restrictions
func ToMax10Text(i interface{}) (Max10Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max10TextLength, Max10TextMinLength, Max10TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max10Text", s)
	}

	return Max10Text(s), nil
}

// MustToMax10Text  method for easy conversion with application of restrictions. Panics on error.
func MustToMax10Text(s interface{}) Max10Text {
	v, err := ToMax10Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * NullableMax20Text Type Def
 */

type NullableMax20Text sql.NullString

const (
	NullableMax20TextZero      = ""
	NullableMax20TextLength    = 0
	NullableMax20TextMinLength = 0
	NullableMax20TextMaxLength = 20
)

// IsValid checks if Max105Text of type String is valid
func (t NullableMax20Text) IsValid() bool {
	return isLengthRestrictionValid(t.ToString(), NullableMax20TextLength, NullableMax20TextMinLength, NullableMax20TextMaxLength)
}

// ToString method for easy conversion
func (t NullableMax20Text) ToString() string {
	if t.Valid {
		return t.String
	}
	return ""
}

// ToNullableMax20Text  method for easy conversion with application of restrictions
func ToNullableMax20Text(i interface{}) (NullableMax20Text, error) {

	s := ""
	switch ti := i.(type) {
	case sql.NullString:
		if ti.Valid {
			s = ti.String
		}

	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return NullableMax20Text(sql.NullString{}), fmt.Errorf("")
	}

	if !isLengthRestrictionValid(s, NullableMax20TextLength, NullableMax20TextMinLength, NullableMax20TextMaxLength) {
		return NullableMax20Text(sql.NullString{}), fmt.Errorf("cannot satisfy length restriction for %s of type NullableMax20Text", s)
	}

	return NullableMax20Text(sqlutil.ToSqlNullString(s)), nil
}

// MustToNullableMax20Text  method for easy conversion with application of restrictions. Panics on error.
func MustToNullableMax20Text(s interface{}) NullableMax20Text {
	v, err := ToNullableMax20Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * NullableMax0Text Type Def
 */

type NullableMax0Text sql.NullString

const (
	NullableMax0TextZero      = ""
	NullableMax0TextLength    = 0
	NullableMax0TextMinLength = 0
	NullableMax0TextMaxLength = 0
)

// IsValid checks if Max105Text of type String is valid
func (t NullableMax0Text) IsValid() bool {
	return isLengthRestrictionValid(t.ToString(), NullableMax0TextLength, NullableMax0TextMinLength, NullableMax0TextMaxLength)
}

// ToString method for easy conversion
func (t NullableMax0Text) ToString() string {
	if t.Valid {
		return t.String
	}
	return ""
}

// ToNullableMax0Text  method for easy conversion with application of restrictions
func ToNullableMax0Text(i interface{}) (NullableMax0Text, error) {

	s := ""
	switch ti := i.(type) {
	case sql.NullString:
		if ti.Valid {
			s = ti.String
		}

	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return NullableMax0Text(sql.NullString{}), fmt.Errorf("")
	}

	if !isLengthRestrictionValid(s, NullableMax0TextLength, NullableMax0TextMinLength, NullableMax0TextMaxLength) {
		return NullableMax0Text(sql.NullString{}), fmt.Errorf("cannot satisfy length restriction for %s of type NullableMax0Text", s)
	}

	return NullableMax0Text(sqlutil.ToSqlNullString(s)), nil
}

// MustToNullableMax0Text  method for easy conversion with application of restrictions. Panics on error.
func MustToNullableMax0Text(s interface{}) NullableMax0Text {
	v, err := ToNullableMax0Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * NullableMax30Text Type Def
 */

type NullableMax30Text sql.NullString

const (
	NullableMax30TextZero      = ""
	NullableMax30TextLength    = 0
	NullableMax30TextMinLength = 0
	NullableMax30TextMaxLength = 30
)

// IsValid checks if Max105Text of type String is valid
func (t NullableMax30Text) IsValid() bool {
	return isLengthRestrictionValid(t.ToString(), NullableMax30TextLength, NullableMax30TextMinLength, NullableMax30TextMaxLength)
}

// ToString method for easy conversion
func (t NullableMax30Text) ToString() string {
	if t.Valid {
		return t.String
	}
	return ""
}

// ToNullableMax30Text  method for easy conversion with application of restrictions
func ToNullableMax30Text(i interface{}) (NullableMax30Text, error) {

	s := ""
	switch ti := i.(type) {
	case sql.NullString:
		if ti.Valid {
			s = ti.String
		}

	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return NullableMax30Text(sql.NullString{}), fmt.Errorf("")
	}

	if !isLengthRestrictionValid(s, NullableMax30TextLength, NullableMax30TextMinLength, NullableMax30TextMaxLength) {
		return NullableMax30Text(sql.NullString{}), fmt.Errorf("cannot satisfy length restriction for %s of type NullableMax30Text", s)
	}

	return NullableMax30Text(sqlutil.ToSqlNullString(s)), nil
}

// MustToNullableMax30Text  method for easy conversion with application of restrictions. Panics on error.
func MustToNullableMax30Text(s interface{}) NullableMax30Text {
	v, err := ToNullableMax30Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * NullableMax80Text Type Def
 */

type NullableMax80Text sql.NullString

const (
	NullableMax80TextZero      = ""
	NullableMax80TextLength    = 0
	NullableMax80TextMinLength = 0
	NullableMax80TextMaxLength = 80
)

// IsValid checks if Max105Text of type String is valid
func (t NullableMax80Text) IsValid() bool {
	return isLengthRestrictionValid(t.ToString(), NullableMax80TextLength, NullableMax80TextMinLength, NullableMax80TextMaxLength)
}

// ToString method for easy conversion
func (t NullableMax80Text) ToString() string {
	if t.Valid {
		return t.String
	}
	return ""
}

// ToNullableMax80Text  method for easy conversion with application of restrictions
func ToNullableMax80Text(i interface{}) (NullableMax80Text, error) {

	s := ""
	switch ti := i.(type) {
	case sql.NullString:
		if ti.Valid {
			s = ti.String
		}

	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return NullableMax80Text(sql.NullString{}), fmt.Errorf("")
	}

	if !isLengthRestrictionValid(s, NullableMax80TextLength, NullableMax80TextMinLength, NullableMax80TextMaxLength) {
		return NullableMax80Text(sql.NullString{}), fmt.Errorf("cannot satisfy length restriction for %s of type NullableMax80Text", s)
	}

	return NullableMax80Text(sqlutil.ToSqlNullString(s)), nil
}

// MustToNullableMax80Text  method for easy conversion with application of restrictions. Panics on error.
func MustToNullableMax80Text(s interface{}) NullableMax80Text {
	v, err := ToNullableMax80Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * NullableMax10Text Type Def
 */

type NullableMax10Text sql.NullString

const (
	NullableMax10TextZero      = ""
	NullableMax10TextLength    = 0
	NullableMax10TextMinLength = 0
	NullableMax10TextMaxLength = 10
)

// IsValid checks if Max105Text of type String is valid
func (t NullableMax10Text) IsValid() bool {
	return isLengthRestrictionValid(t.ToString(), NullableMax10TextLength, NullableMax10TextMinLength, NullableMax10TextMaxLength)
}

// ToString method for easy conversion
func (t NullableMax10Text) ToString() string {
	if t.Valid {
		return t.String
	}
	return ""
}

// ToNullableMax10Text  method for easy conversion with application of restrictions
func ToNullableMax10Text(i interface{}) (NullableMax10Text, error) {

	s := ""
	switch ti := i.(type) {
	case sql.NullString:
		if ti.Valid {
			s = ti.String
		}

	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return NullableMax10Text(sql.NullString{}), fmt.Errorf("")
	}

	if !isLengthRestrictionValid(s, NullableMax10TextLength, NullableMax10TextMinLength, NullableMax10TextMaxLength) {
		return NullableMax10Text(sql.NullString{}), fmt.Errorf("cannot satisfy length restriction for %s of type NullableMax10Text", s)
	}

	return NullableMax10Text(sqlutil.ToSqlNullString(s)), nil
}

// MustToNullableMax10Text  method for easy conversion with application of restrictions. Panics on error.
func MustToNullableMax10Text(s interface{}) NullableMax10Text {
	v, err := ToNullableMax10Text(s)
	if err != nil {
		panic(err)
	}

	return v
}
