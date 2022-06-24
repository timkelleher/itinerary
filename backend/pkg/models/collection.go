package models

type Collection struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	PathPrefix  string `db:"path_prefix"`
	ContentType string `db:"content_type"`
	Endpoints   []Endpoint
}

type CollectionRequest struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	PathPrefix  string `json:"path_prefix"`
	ContentType string `json:"content_type"`
}
