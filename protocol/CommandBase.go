package protocol

// Header is prefixed to every body.
type Header struct {
	ProtocolVersion uint8
	CommandID       uint8
	BodyLength      uint16
}

// GetBodyLength get body length
func (h *Header) GetBodyLength() int {
	return 0
}

// EncodeBody can encode client message to binary
func (h *Header) EncodeBody() int {
	return 0
}

// DecodeBody can decode binary code to struct
func (h *Header) DecodeBody([]byte) int {
	return 0
}
