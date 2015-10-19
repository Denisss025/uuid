// Package uuid provides a simple binding to the C-library
// libuuid.
package uuid

// #cgo pkg-config: uuid
// #cgo LDFLAGS: -luuid
// #include <string.h>
// #include <uuid/uuid.h>
import "C"

import (
	"crypto/rand"
	"encoding/hex"
	"unsafe"
)

// Uuid holds the binary representation of a universally unique
// identifier (UUID).
type Uuid struct {
	uuid C.uuid_t
}

// Clear sets the value of the Uuid variable uu to the nil value.
func (uu *Uuid) Clear() {
	C.uuid_clear(&uu.uuid[0])
}

// CompareTo compares the two supplied Uuid variables uu and uu2 to
// each other.
// Returns an integer less than, equal to, or greater than zero if uu is found,
// respectively, to be lexicographically less than, equal, or greater than uu2.
func (uu Uuid) CompareTo(uu2 Uuid) int {
	return int(C.uuid_compare(&uu.uuid[0], &uu2.uuid[0]))
}

// CopyTo copies the uu into uu2.
func (uu Uuid) CopyTo(uu2 *Uuid) {
	C.uuid_copy(&uu2.uuid[0], &uu.uuid[0])
}

// Copy returns a copy of the Uuid variable uu.
func (uu Uuid) Copy() (result Uuid) {
	C.uuid_copy(&result.uuid[0], &uu.uuid[0])
	return
}

// IsNil compares the value of the supplied Uuid variable uu to the Nil value.
func (uu Uuid) IsNil() bool {
	return C.uuid_is_null(&uu.uuid[0]) == 1
}

// Generate creates a new universally unique identifier (UUID).
// The UUID will be generated based on high-quality randomness generator.
// (i.e.from /dev/urandom), if available. If it is not available, then
// it will use an alternative algorithm which uses the current time, the
// local ethernet MAC address (if available), and random data generated using
// a pseudo-random generator.
func (uu *Uuid) Generate() {
	C.uuid_generate(&uu.uuid[0])
}

// GenerateRandom forces the use of the all-random UUID format, even if a
// high-quality generator (i.e. /dev/urandom) is not available, in which
// case a pseudo-random generator will be substituted. Note that the use
// of a pseudo-random generator may compromise the uniqueness of UUIDs
// generated in this fasion.
func (uu *Uuid) GenerateRandom() error {
	//C.uuid_generate_random(&uu.uuid[0])
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)

	if n != len(uuid) || err != nil {
		return err
	}

	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40

	for i, u := range uuid {
		uu.uuid[i] = C.uchar(u)
	}

	return nil
}

// GenerateTime furces the use of the alternative algorithm which uses the
// current time and the local ethernet MAC address, it can leak information
// about when and where the UUID was generated.
func (uu *Uuid) GenerateTime() {
	C.uuid_generate_time(&uu.uuid[0])
}

// GenerateTimeSafe is similar to GenerateTime, except that it returns a value
// which denotes whether any of the synchronization mechanisms has been used.
func (uu *Uuid) GenerateTimeSafe() {
	C.uuid_generate_time_safe(&uu.uuid[0])
}

// String converts the suplied Uuid uu from the binary representation into
// a 32-byte string value of the form '1b4e28ba2fa111d2883f0016d3cca427'.
func (uu Uuid) String() string {
	if uu.IsNil() {
		return ""
	}

	result := make([]byte, 32)
	hex.Encode(result, C.GoBytes(unsafe.Pointer(&uu.uuid[0]), 16))
	return string(result)
}

// ToGuid converts the suplied Uuid uu from the binary representation into
// a 36-byte string value of the form '1b4e28ba-2fa1-11d2-883f-0016d3cca427'.
func (uu Uuid) ToGuid() string {
	if uu.IsNil() {
		return ""
	}
	uuid := C.GoBytes(unsafe.Pointer(&uu.uuid[0]), 16)

	result := make([]byte, 36)
	hex.Encode(result[:8], uuid[:4])
	result[8] = '-'
	hex.Encode(result[9:], uuid[4:6])
	result[13] = '-'
	hex.Encode(result[14:], uuid[6:8])
	result[18] = '-'
	hex.Encode(result[19:], uuid[8:10])
	result[23] = '-'
	hex.Encode(result[24:], uuid[10:])

	return string(result)
}
