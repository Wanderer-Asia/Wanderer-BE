package filters

type Sort struct {
	Column    string `query:"sort"`
	Direction bool   `query:"dir"`
}
