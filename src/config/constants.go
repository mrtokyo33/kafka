package config

const (
	AppName = "Kafka"
	Version = "1.0.0"

	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
)

const (
	MemeColor       = 0xffbc47
	MemeMaxRetries  = 10
	MemeEmojiUpvote = "⬆️"

	MemeWeightHigh = 2
	MemeWeightLow  = 1
)

var DefaultMemeSubreddits = map[string]int{
	"ShitpostBR":     MemeWeightHigh,
	"botecodoreddit": MemeWeightHigh,
	"DiretoDoZapZap": MemeWeightHigh,
	"MemesBR":        MemeWeightLow,
	"eu_nvr":         MemeWeightLow,
	"O_PACOTE":       MemeWeightLow,
}
