package models

import "strings"

const GLOBAL_PATH_PREFIX = "/api"

type Endpoint struct {
	ID                 int64 `db:"id"`
	CollectionID       int64 `db:"collection_id"`
	Collection         *Collection
	Name               string `db:"name"`
	Method             string `db:"method"`
	ResponseStatusCode string `db:"response_status_code"`
	PathPrefix         string `db:"path_prefix"`
	Path               string `db:"path"`
	ComputedPath       string `db:"computed_path"`
	Output             string `db:"output"`
}

type EndpointRequest struct {
	ID                 int64  `json:"id"`
	CollectionID       int64  `json:"collection_id"`
	Name               string `json:"name"`
	Method             string `json:"method"`
	ResponseStatusCode string `json:"response_status_code"`
	Path               string `json:"path"`
	Output             string `json:"output"`
}

func (e Endpoint) ComputePath() string {
	fullPath := []string{GLOBAL_PATH_PREFIX}

	prefixTrimmed := strings.Trim(e.Collection.PathPrefix, " /")
	pathTrimmed := strings.Trim(e.Path, " /")
	prefix := strings.Split(prefixTrimmed, "/")
	path := strings.Split(pathTrimmed, "/")

	fullPath = append(fullPath, prefix...)
	fullPath = append(fullPath, path...)
	return strings.Join(fullPath, "/")
}
