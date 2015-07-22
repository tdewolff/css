package svg

// generated by hasher -type=Hash -file=hash.go; DO NOT EDIT, except for adding more constants to the list and rerun go generate

// uses github.com/tdewolff/hasher
//go:generate hasher -type=Hash -file=hash.go
type Hash uint32

const (
	D        Hash = 0x401
	G        Hash = 0x1301
	Metadata Hash = 0x8
	Path     Hash = 0x804
	Style    Hash = 0xc05
	Svg      Hash = 0x1103
	Version  Hash = 0x1407
)

// String returns the hash' name.
func (i Hash) String() string {
	start := uint32(i >> 8)
	n := uint32(i & 0xff)
	if start+n > uint32(len(_Hash_text)) {
		return ""
	}
	return _Hash_text[start : start+n]
}

// Hash returns the hash whose name is s. It returns zero if there is no
// such hash. It is case sensitive.
func ToHash(s []byte) Hash {
	if len(s) == 0 || len(s) > _Hash_maxLen {
		return 0
	}
	h := _Hash_fnv(s)
	if i := _Hash_table[h&uint32(len(_Hash_table)-1)]; int(i&0xff) == len(s) && _Hash_match(_Hash_string(i), s) {
		return i
	}
	if i := _Hash_table[(h>>16)&uint32(len(_Hash_table)-1)]; int(i&0xff) == len(s) && _Hash_match(_Hash_string(i), s) {
		return i
	}
	return 0
}

// _Hash_fnv computes the FNV hash with an arbitrary starting value h.
func _Hash_fnv(s []byte) uint32 {
	h := uint32(_Hash_hash0)
	for i := range s {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return h
}

func _Hash_match(s string, t []byte) bool {
	for i, c := range t {
		if s[i] != c {
			return false
		}
	}
	return true
}

func _Hash_string(i Hash) string {
	return _Hash_text[i>>8 : i>>8+i&0xff]
}

const _Hash_hash0 = 0x9acb0442
const _Hash_maxLen = 8
const _Hash_text = "metadatapathstylesvgversion"

var _Hash_table = [1 << 3]Hash{
	0x0: 0x1103, // svg
	0x1: 0x804,  // path
	0x2: 0x401,  // d
	0x4: 0x1407, // version
	0x5: 0x8,    // metadata
	0x6: 0xc05,  // style
	0x7: 0x1301, // g
}
