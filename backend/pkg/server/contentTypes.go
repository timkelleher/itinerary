package backend

// Add more later
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types
var contentTypesMap = map[string]string{
	"text/plain":       "Plain Text",
	"application/json": "JSON",
}

var contentTypes []ContentType

type ContentType struct {
	ContentType string
	Label       string
}

func ContentTypes() []ContentType {
	if len(contentTypes) == 0 {
		for content, label := range contentTypesMap {
			contentTypes = append(contentTypes, ContentType{ContentType: content, Label: label})
		}
	}
	return contentTypes
}
