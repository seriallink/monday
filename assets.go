package monday

import "time"

type Asset struct {
	CreatedAt        time.Time `json:"created_at"        graphql:"created_at"`
	FileExtension    string    `json:"file_extension"    graphql:"file_extension"`
	FileSize         int       `json:"file_size"         graphql:"file_size"`
	Id               ID        `json:"id"                graphql:"id"`
	Name             string    `json:"name"              graphql:"name"`
	OriginalGeometry string    `json:"original_geometry" graphql:"original_geometry"`
	PublicUrl        string    `json:"public_url"        graphql:"public_url"`
	UploadedBy       *User     `json:"uploaded_by"       graphql:"uploaded_by"`
	Url              string    `json:"url"               graphql:"url"`
	UrlThumbnail     string    `json:"url_thumbnail"     graphql:"url_thumbnail"`
}
