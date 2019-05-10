package protocol

// ProtocolVersion Version.
const ProtocolVersion = 1

// Command Type
const (
	// CommandUnknown is command id for unknown type
	CommandUnknown = 0x00
	// CommandPings is command id for message Ping
	CommandPings = 0x05
	// CommandPingReply is command id for unknown type
	CommandPingReply = 0x06
	// CommandFin1 is command id for unknown type
	CommandFin1 = 0x07
	// CommandBlockRequest is command id for block type
	CommandBlockRequest = 0x30
	// CommandBlockResponse is command id for block type
	CommandBlockResponse = 0x31
	// CommandPieceRequest is command id for piece type
	CommandPieceRequest = 0x32
	// CommandPieceResponse is command id for piece type
	CommandPieceResponse = 0x33
)

// Client Type
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

// Fin Code  ==  Server => SDK
const (
	// Fin_FileNotExist file no exit
	FinFileNotExist = 1
	// Fin_LastModifiedDiffer last modify file differ
	FinLastModifiedDiffer = 2
	// Fin_ConnectionLimit fin link limit
	FinConnectionLimit = 3
	// FinNoPieceAtOffset no piece at offset
	FinNoPieceAtOffset = 4
	//FinLongTimeNoRequest long time no request
	FinLongTimeNoRequest = 5
	//FinFileSizeDiffer file size differ
	FinFileSizeDiffer = 6
	// FinInvalidPieceIndex invalid piece index
	FinInvalidPieceIndex = 7
	// FinInvalidPieceOffset invalid piece offset
	FinInvalidPieceOffset = 8
	// FinFetchPieceFail fetch piece fail
	FinFetchPieceFail = 9
)

// Fin Code  =    SDK => Server
const (
	// FinSDKLowSpeed sdk low speed
	FinSDKLowSpeed = 51
	// FinSDKTimeout sdk timeout
	FinSDKTimeout = 52
	// FinSDKNormalClose sdk normal close
	FinSDKNormalClose = 53
	//FinSDKPieceCheckFail sdk piece chaeck fail
	FinSDKPieceCheckFail = 54
	//FinSDKDecodeFail sdk decode file
	FinSDKDecodeFail = 55
	//FinSDKHandShakeFail sdk hand shake fail
	FinSDKHandShakeFail = 56
)
