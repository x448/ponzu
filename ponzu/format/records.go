package format

import (
	"time"
)

// There are a handful of record types. These are CBOR types.

/*
*
Record base that all other things are based on.
This base structure contains the relevant information that goes into or comes out of
the preamble, plus a Reader that is used in context.
*/

type RecordBase struct {
	// Preamble
	Preamble Preamble `cbor:"-"`
}

// All archives start with a Start of Archive header
type StartOfArchive struct {
	RecordBase
	// Version of the archive
	Version uint8 `cbor:"0,keyasint"`
	// Host OS the archive was made on
	Host string `cbor:"1,keyasint"`
	// Prefix to write all files to
	Prefix string `cbor:"2,keyasint"`
	// Comment for the archive (open text field)
	Comment string `cbor:"3,keyasint"`
}

type File struct {
	RecordBase
	Name       string          `cbor:"0, keyasint"`
	Mode       uint16          `cbor:"1, keyasint"`
	Owner      string          `cbor:"2, keyasint"`
	Group      string          `cbor:"3, keyasint"`
	ModTime    time.Time       `cbor:"4, keyasint"`
	Compressor CompressionType `cbor:"5, keyasint"`
	Metadata   map[string]any  `cbor:"6, keyasint"`
}

type Link struct {
	Target string `cbor:"-1, keyasint"`
}

type Symlink struct{ Link }
type Hardlink struct{ Link }
type Directory struct{ File }
type ZstdDictionary struct{ RecordBase }

type OSSpecial struct {
	File
	SpecialType string `cbor:"-1, keyasint"`
	Mode        uint32 `cbor:"mknodMode"`
	Device      uint32 `cbor:"mknodDev"`
}
