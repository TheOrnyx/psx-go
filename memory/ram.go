package memory

import "github.com/TheOrnyx/psx-go/utils"

type Ram struct {
	data []uint8
}

// NewRam create and return a new ram object and fill it with garbage
func NewRam() Ram {
	data := make([]uint8, 0x200000)
	data[0] = 0xca

	// thank youu https://gist.github.com/taylorza/df2f89d5f9ab3ffd06865062a4cf015d
	for i := 1; i < len(data); i *=2 {
		copy(data[i:], data[:i])
	}

	return Ram{data: data}
}

// load32 load 32 bit little endian word at offset in data
func (r *Ram) load32(offset uint32) uint32 {
	b0 := r.data[offset + 0]
	b1 := r.data[offset + 1]
	b2 := r.data[offset + 2]
	b3 := r.data[offset + 3]

	return utils.BytesToUint32(b0, b1, b2, b3)
}

// store32 store value val into data at offset location
func (r *Ram) store32(offset, val uint32)  {
	b0, b1, b2, b3 := utils.Uint32ToBytes(val)

	r.data[offset + 0] = b0
	r.data[offset + 1] = b1
	r.data[offset + 2] = b2
	r.data[offset + 3] = b3
}
