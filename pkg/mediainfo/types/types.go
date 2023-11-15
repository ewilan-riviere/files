package types

func FilterItems(key string, items []MediainfoOutputItem) string {
	for _, item := range items {
		if item.Key == key {
			return item.Value
		}
	}

	return ""
}

type MediainfoOutput struct {
	Category string
	Items    []MediainfoOutputItem
}

type MediainfoOutputItem struct {
	Key   string
	Value string
}

type MediainfoItem struct {
	UniqueID           string
	CompleteName       string
	Format             string
	FormatVersion      string
	FileSize           string
	Duration           string
	OverallBitRate     string
	Album              string
	AlbumPerformer     string
	PartPosition       string
	TrackName          string
	TrackNamePosition  string
	Performer          string
	Genre              string
	RecordedDate       string
	FrameRate          string
	MovieName          string
	EncodedDate        string
	WritingApplication string
	WritingLibrary     string
	Cover              string
	CoverType          string
	CoverMime          string
	Comment            string
	Attachments        string
	Videos             []MediainfoItemVideo
	Audios             []MediainfoItemAudio
	Texts              []MediainfoItemText
	Images             []MediainfoItemImage
	Menu               map[string]string
}

type MediainfoItemVideo struct {
	ID                            string
	Format                        string
	FormatInfo                    string
	FormatProfile                 string
	FormatSettings                string
	FormatSettingsCABAC           string
	FormatSettingsReferenceFrames string
	CodecID                       string
	Duration                      string
	BitRate                       string
	Width                         string
	Height                        string
	DisplayAspectRatio            string
	FrameRateMode                 string
	FrameRate                     string
	ColorSpace                    string
	ChromaSubsampling             string
	BitDepth                      string
	ScanType                      string
	BitsPixelFrame                string
	StreamSize                    string
	Title                         string
	WritingLibrary                string
	EncodingSettings              string
	Default                       string
	Forced                        string
	ColorRange                    string
	ColorPrimaries                string
	TransferCharacteristics       string
	MatrixCoefficients            string
}

type MediainfoItemAudio struct {
	ID               string
	Format           string
	FormatInfo       string
	FormatProfile    string
	FormatVersion    string
	CommercialName   string
	CodecID          string
	Duration         string
	BitRateMode      string
	BitRate          string
	Channel          string
	ChannelLayout    string
	SamplingRate     string
	FrameRate        string
	Compression      string
	StreamSize       string
	WritingLibrary   string
	EncodingSettings string
	Title            string
	Language         string
	ServiceKind      string
	Default          string
	Forced           string
}

type MediainfoItemText struct {
	ID             string
	Format         string
	CodecID        string
	CodecIDInfo    string
	Duration       string
	BitRate        string
	FrameRate      string
	CountOfElement string
	StreamSize     string
	Title          string
	Language       string
	Default        string
	Forced         string
}

type MediainfoItemImage struct {
	ID                string
	Format            string
	Width             string
	Height            string
	ColorSpace        string
	ChromaSubsampling string
	BitDepth          string
	CompressionMode   string
	StreamSize        string
}
