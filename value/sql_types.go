package value

import (
	"reflect"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/guregu/null"
)

// Value represents go data types which carta supports for loading as well as what data types arrive from the sql driver
type Value int

// NOTE, carta does NOT support loading []uint8, any data that arrives from sql database as []uint8
//is converted to bytes and expected field type is a string or *string
const (
	Invalid Value = iota
	Time
	Timestamp
	NullBool
	NullFloat64
	NullInt32
	NullInt64
	NullString
	NullTime
	Float64
	Float32
	Int
	Int8
	Int16
	Int32
	Uint
	Uint8
	Uint16
	Uint32
	Int64
	Uint64
	Bool
	String //  note, []uint8get converted to string, this is because mysql returns []uint8 for varchar while pg returns string
)

var BasicKinds = map[reflect.Kind]Value{
	reflect.Bool:    Bool,
	reflect.Int:     Int,
	reflect.Int8:    Int8,
	reflect.Int16:   Int16,
	reflect.Int32:   Int32,
	reflect.Int64:   Int64,
	reflect.Uint:    Uint,
	reflect.Uint8:   Uint8,
	reflect.Uint16:  Uint16,
	reflect.Uint32:  Uint32,
	reflect.Uint64:  Uint64,
	reflect.Float32: Float32,
	reflect.Float64: Float64,
	reflect.String:  String,
}

var BasicTypes = map[reflect.Type]Value{
	reflect.TypeOf(time.Time{}):           Time,
	reflect.TypeOf(timestamp.Timestamp{}): Timestamp,
	reflect.TypeOf(null.Bool{}):           NullBool,
	reflect.TypeOf(null.Float{}):          NullFloat64,
	reflect.TypeOf(null.Int{}):            NullInt32,
	reflect.TypeOf(null.String{}):         NullString,
	reflect.TypeOf(null.Time{}):           NullTime,
}

var NullableTypes = map[reflect.Type]Value{
	reflect.TypeOf(null.Bool{}):   NullBool,
	reflect.TypeOf(null.Float{}):  NullFloat64,
	reflect.TypeOf(null.Int{}):    NullInt32,
	reflect.TypeOf(null.String{}): NullString,
	reflect.TypeOf(null.Time{}):   NullTime,
}

// Map of database data types to go types
// var SQLTypes = map[string]Value{
// "VARCHAR":  String,
// "TEXT":     String,
// "NVARCHAR": String,
// "NUMERIC":  String,
// "UUID":     String,
// "BPCHAR":   String,
// "BIT":      String,
// "CIDR":     String,
// "XML":      String,
// "OID":      String,
//
// "DECIMAL": Float64,
// "FLOAT8":  Float64,
// "FLOAT4":  Float64,
//
// "BOOL": Bool,
//
// "INT":  Int64,
// "INT2": Int64,
// "INT4": Int64,
// "INT8": Int64,
//
// "TIME":        Time,
// "DATE":        Time,
// "TIMESTAMP":   Time,
// "TIMESTAMPZ":  Time,
// "TIMETZ":      Time,
// "TIMESTAMPTZ": Time,
// }
