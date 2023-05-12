package docclass

import (
	"database/sql"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-batis/sqlutil"
)

const (
	IdFieldIndex               = 0
	NameFieldIndex             = 1
	ProducedClassIdsFieldIndex = 2
	ExtFieldIndex              = 3
	CodClienteFieldIndex       = 4
	MaxCpxFieldIndex           = 5
	MaxDocFieldIndex           = 6
	MaxSizeFieldIndex          = 7
	PlatformFieldIndex         = 8
	ServizioFieldIndex         = 9
	ProceduraFieldIndex        = 10
	VersionFieldIndex          = 11
	PackageLayoutFieldIndex    = 12
	SqlQueryFieldIndex         = 13
	DistintaGEDFieldIndex      = 14
)

type DocClassEntity struct {
	Id               Max30Text      `db:"class_id"`
	Name             Max80Text      `db:"name"`
	ProducedClassIds sql.NullString `db:"child_ids"`
	Ext              Max3Text       `db:"extension"`
	CodCliente       Max10Text      `db:"cod_cliente"`
	MaxCpx           int32          `db:"max_cpx"`
	MaxDoc           int32          `db:"max_docs"`
	MaxSize          int32          `db:"max_size"`
	Platform         Max10Text      `db:"platform"`
	Servizio         Max10Text      `db:"servizio"`
	Procedura        Max10Text      `db:"procedura"`
	Version          Max10Text      `db:"version"`
	PackageLayout    sql.NullString `db:"pkg_layout"`
	SqlQuery         Max1024Text    `db:"sql_query"`
	DistintaGED      bool           `db:"distinta_ged"`
}

type PrimaryKey struct {
	Id Max30Text `db:"class_id"`
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
 * Max2Text Type Def
 */

type Max2Text string

const (
	Max2TextZero      = ""
	Max2TextLength    = 0
	Max2TextMinLength = 1
	Max2TextMaxLength = 2
)

// IsValid checks if Max105Text of type String is valid
func (t Max2Text) IsValid() bool {
	return isLengthRestrictionValid(t.ToString(), Max2TextLength, Max2TextMinLength, Max2TextMaxLength)
}

// ToString method for easy conversion
func (t Max2Text) ToString() string {
	return string(t)
}

// ToMax2Text  method for easy conversion with application of restrictions
func ToMax2Text(i interface{}) (Max2Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max2TextLength, Max2TextMinLength, Max2TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max2Text", s)
	}

	return Max2Text(s), nil
}

// MustToMax2Text  method for easy conversion with application of restrictions. Panics on error.
func MustToMax2Text(s interface{}) Max2Text {
	v, err := ToMax2Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * Max1024Text Type Def
 */

type Max1024Text string

const (
	Max1024TextZero      = ""
	Max1024TextLength    = 0
	Max1024TextMinLength = 1
	Max1024TextMaxLength = 1024
)

// IsValid checks if Max105Text of type String is valid
func (t Max1024Text) IsValid() bool {
	return isLengthRestrictionValid(t.ToString(), Max1024TextLength, Max1024TextMinLength, Max1024TextMaxLength)
}

// ToString method for easy conversion
func (t Max1024Text) ToString() string {
	return string(t)
}

// ToMax1024Text  method for easy conversion with application of restrictions
func ToMax1024Text(i interface{}) (Max1024Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max1024TextLength, Max1024TextMinLength, Max1024TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max1024Text", s)
	}

	return Max1024Text(s), nil
}

// MustToMax1024Text  method for easy conversion with application of restrictions. Panics on error.
func MustToMax1024Text(s interface{}) Max1024Text {
	v, err := ToMax1024Text(s)
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
 * Max3Text Type Def
 */

type Max3Text string

const (
	Max3TextZero      = ""
	Max3TextLength    = 0
	Max3TextMinLength = 1
	Max3TextMaxLength = 3
)

// IsValid checks if Max105Text of type String is valid
func (t Max3Text) IsValid() bool {
	return isLengthRestrictionValid(t.ToString(), Max3TextLength, Max3TextMinLength, Max3TextMaxLength)
}

// ToString method for easy conversion
func (t Max3Text) ToString() string {
	return string(t)
}

// ToMax3Text  method for easy conversion with application of restrictions
func ToMax3Text(i interface{}) (Max3Text, error) {

	s := ""
	switch ti := i.(type) {
	case fmt.Stringer:
		s = ti.String()
	case string:
		s = ti
	default:
		return "", fmt.Errorf("")
	}
	if !isLengthRestrictionValid(s, Max3TextLength, Max3TextMinLength, Max3TextMaxLength) {
		return "", fmt.Errorf("cannot satisfy length restriction for %s of type Max3Text", s)
	}

	return Max3Text(s), nil
}

// MustToMax3Text  method for easy conversion with application of restrictions. Panics on error.
func MustToMax3Text(s interface{}) Max3Text {
	v, err := ToMax3Text(s)
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
 * NullableMax3Text Type Def
 */

type NullableMax3Text sql.NullString

const (
	NullableMax3TextZero      = ""
	NullableMax3TextLength    = 0
	NullableMax3TextMinLength = 0
	NullableMax3TextMaxLength = 3
)

// IsValid checks if Max105Text of type String is valid
func (t NullableMax3Text) IsValid() bool {
	return isLengthRestrictionValid(t.ToString(), NullableMax3TextLength, NullableMax3TextMinLength, NullableMax3TextMaxLength)
}

// ToString method for easy conversion
func (t NullableMax3Text) ToString() string {
	if t.Valid {
		return t.String
	}
	return ""
}

// ToNullableMax3Text  method for easy conversion with application of restrictions
func ToNullableMax3Text(i interface{}) (NullableMax3Text, error) {

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
		return NullableMax3Text(sql.NullString{}), fmt.Errorf("")
	}

	if !isLengthRestrictionValid(s, NullableMax3TextLength, NullableMax3TextMinLength, NullableMax3TextMaxLength) {
		return NullableMax3Text(sql.NullString{}), fmt.Errorf("cannot satisfy length restriction for %s of type NullableMax3Text", s)
	}

	return NullableMax3Text(sqlutil.ToSqlNullString(s)), nil
}

// MustToNullableMax3Text  method for easy conversion with application of restrictions. Panics on error.
func MustToNullableMax3Text(s interface{}) NullableMax3Text {
	v, err := ToNullableMax3Text(s)
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

/*
 * NullableMax2Text Type Def
 */

type NullableMax2Text sql.NullString

const (
	NullableMax2TextZero      = ""
	NullableMax2TextLength    = 0
	NullableMax2TextMinLength = 0
	NullableMax2TextMaxLength = 2
)

// IsValid checks if Max105Text of type String is valid
func (t NullableMax2Text) IsValid() bool {
	return isLengthRestrictionValid(t.ToString(), NullableMax2TextLength, NullableMax2TextMinLength, NullableMax2TextMaxLength)
}

// ToString method for easy conversion
func (t NullableMax2Text) ToString() string {
	if t.Valid {
		return t.String
	}
	return ""
}

// ToNullableMax2Text  method for easy conversion with application of restrictions
func ToNullableMax2Text(i interface{}) (NullableMax2Text, error) {

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
		return NullableMax2Text(sql.NullString{}), fmt.Errorf("")
	}

	if !isLengthRestrictionValid(s, NullableMax2TextLength, NullableMax2TextMinLength, NullableMax2TextMaxLength) {
		return NullableMax2Text(sql.NullString{}), fmt.Errorf("cannot satisfy length restriction for %s of type NullableMax2Text", s)
	}

	return NullableMax2Text(sqlutil.ToSqlNullString(s)), nil
}

// MustToNullableMax2Text  method for easy conversion with application of restrictions. Panics on error.
func MustToNullableMax2Text(s interface{}) NullableMax2Text {
	v, err := ToNullableMax2Text(s)
	if err != nil {
		panic(err)
	}

	return v
}

/*
 * NullableMax1024Text Type Def
 */

type NullableMax1024Text sql.NullString

const (
	NullableMax1024TextZero      = ""
	NullableMax1024TextLength    = 0
	NullableMax1024TextMinLength = 0
	NullableMax1024TextMaxLength = 1024
)

// IsValid checks if Max105Text of type String is valid
func (t NullableMax1024Text) IsValid() bool {
	return isLengthRestrictionValid(t.ToString(), NullableMax1024TextLength, NullableMax1024TextMinLength, NullableMax1024TextMaxLength)
}

// ToString method for easy conversion
func (t NullableMax1024Text) ToString() string {
	if t.Valid {
		return t.String
	}
	return ""
}

// ToNullableMax1024Text  method for easy conversion with application of restrictions
func ToNullableMax1024Text(i interface{}) (NullableMax1024Text, error) {

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
		return NullableMax1024Text(sql.NullString{}), fmt.Errorf("")
	}

	if !isLengthRestrictionValid(s, NullableMax1024TextLength, NullableMax1024TextMinLength, NullableMax1024TextMaxLength) {
		return NullableMax1024Text(sql.NullString{}), fmt.Errorf("cannot satisfy length restriction for %s of type NullableMax1024Text", s)
	}

	return NullableMax1024Text(sqlutil.ToSqlNullString(s)), nil
}

// MustToNullableMax1024Text  method for easy conversion with application of restrictions. Panics on error.
func MustToNullableMax1024Text(s interface{}) NullableMax1024Text {
	v, err := ToNullableMax1024Text(s)
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
