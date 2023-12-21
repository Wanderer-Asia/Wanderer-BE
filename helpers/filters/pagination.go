package filters

type Pagination struct {
	Start int `query:"start"`
	Limit int `query:"limit"`
}
