package entity

type YoutubeFormats struct {
	Format            string `json:"format"`
	URL               string `json:"url"`
	Vcodec            string `json:"vcodec"`
	EXT               string `json:"ext"`
	Acodec            string `json:"acodec"`
	Filesize          int64  `json:"filesize"`
	DownloaderOptions Chunk  `json:"downloader_options"`
}

type Chunk struct {
	HTTPChunkSize int64 `json:"http_chunk_size"`
}
