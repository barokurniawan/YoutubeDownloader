package service

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/barokurniawan/YoutubeDownloader/constant"
	"github.com/barokurniawan/YoutubeDownloader/entity"
)

//YoutubeDL is the core struct
type YoutubeDL struct {
	youtubeURL  string
	YoutubeInfo entity.YoutubeMetadata
}

//SetYoutubeURL set youtube url
func (dl *YoutubeDL) SetYoutubeURL(url string) *YoutubeDL {
	dl.youtubeURL = url
	return dl
}

//GetYoutubeInfo get youtube info data as json
func (dl *YoutubeDL) GetYoutubeInfo() (entity.YoutubeMetadata, error) {
	var youtubeInfo entity.YoutubeMetadata
	output, err := exec.Command("/usr/local/bin/youtube-dl", "--dump-json", dl.youtubeURL).Output()
	if err != nil {
		return youtubeInfo, err
	}

	err = json.Unmarshal(output, &youtubeInfo)
	if err != nil {
		return youtubeInfo, err
	}

	dl.YoutubeInfo = youtubeInfo
	return dl.YoutubeInfo, nil
}

func (dl *YoutubeDL) GetAudioOnly() entity.YoutubeMetadata {

	var output []entity.YoutubeFormats
	for _, element := range dl.YoutubeInfo.Formats {
		if element.Vcodec != constant.EncoderNone {
			continue
		}

		if element.Acodec == constant.EncoderNone {
			continue
		}

		output = append(output, element)
	}

	dl.YoutubeInfo.Formats = output
	return dl.YoutubeInfo
}

func (dl YoutubeDL) getThumbnail() string {
	return fmt.Sprintf("https://i.ytimg.com/vi/%s/maxresdefault.jpg", dl.YoutubeInfo.ID)
}
