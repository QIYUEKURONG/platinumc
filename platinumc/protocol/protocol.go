package protocol

// ProtocolVersion Version.
const ProtocolVersion byte = 1

const (
	// CommandUnknown is command id for unknown type
	CommandUnknown = 0x00
	// CommandPing is command id for message Ping
	CommandPing = 0x05
	// CommandPingReply is command id for unknown type
	CommandPingReply = 0x06
	// CommandFin is command id for unknown type
	CommandFin = 0x07
	// CommandBlockRequest is command id for block type
	CommandBlockRequest = 0x30
	// CommandBlockResponse is command id for block type
	CommandBlockResponse = 0x31
	// CommandPieceRequest is command id for piece type
	CommandPieceRequest = 0x32
	// CommandPieceResponse is command id for piece type
	CommandPieceResponse = 0x33
)

const (
	// ClientTypeUnknown is unknown type
	ClientTypeUnknown = 0
	// ClientTypeFlash is flash type
	ClientTypeFlash = 1
	// ClientTypeAndroid is Android type
	ClientTypeAndroid = 2
	// ClientTypeiOS is ios type
	ClientTypeiOS = 3
	// ClientTypeWindows is windous type
	ClientTypeWindows = 4
	// ClientTypeMiner is mine type
	ClientTypeMiner = 5
	// ClientTypemacOS is macos type
	ClientTypemacOS = 6
	// ClientTypeHTML5 is HTML5 type
	ClientTypeHTML5 = 7
)

// Header is prefixed to every body.
type Header struct {
	ProtocolVersion uint8
	CommandID       uint8
	BodyLength      uint16
}

// BlockRequest record client_type,client_id and file_offset
type BlockRequest struct {
	Head       Header
	ClientType uint8
	ClientID   string
	FileIndex  string
	FileOffset uint64
}

// BlockResponse receive message from server
type BlockResponse struct {
	Head             Header
	FileIndex        string
	FileOffset       uint64
	FileSize         uint64
	FilelastModified uint64
}
