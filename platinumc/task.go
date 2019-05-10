package platinumc

// Task can record Client task type
type Task struct {
	TransportProtocol string
	ServerAddress     string
	FileIndex         string
	BlockIndex        uint
	SavePath          string
	StartPieceIndex   uint
	FetchPieceCount   uint
	CheckPiece        bool
	VerboseLog        bool
	Identifier        string
}
