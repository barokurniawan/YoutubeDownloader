package entity

import "encoding/json"

type YoutubeMetadata struct {
	ID         string           `json:"id"`
	UploadDate string           `json:"upload_date"`
	Title      string           `json:"title"`
	Creator    string           `json:"creator"`
	Formats    []YoutubeFormats `json:"formats"`
}

func (meta YoutubeMetadata) ToJson() string {
	json, err := json.Marshal(meta)
	if err != nil {
		return ""
	}

	return string(json)
}
