package app

var DefaultConfig = &Config{
	ChainCodeID: "itracker",
	ChannelID:   "inventory",
	ConfigFile:  "config.yaml",
	OrgName:     "IADFactory",
	UserName:    "User1",
}

type Config struct {
	ConfigFile  string
	ChannelID   string
	ChainCodeID string
	OrgName     string
	UserName    string
}
