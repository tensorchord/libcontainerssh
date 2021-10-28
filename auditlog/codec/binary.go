package codec

import (
	"github.com/containerssh/libcontainerssh/internal/auditlog/codec/binary"
)

// NewBinaryDecoder returns a decoder for the ContainerSSH binary audit log protocol.
func NewBinaryDecoder() Decoder {
	return binary.NewDecoder()
}
