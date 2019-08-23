package dbpreader

type DBPediaResult map[string]DBPediaResource

type DBPediaResource map[string][]DBPediaProperty

type DBPediaProperty struct {
	Type     string `json:"type"`
	Value    interface{} `json:"value"`
	Lang     string `json:"lang"`
	DataType string `json:"datatype"`
}
