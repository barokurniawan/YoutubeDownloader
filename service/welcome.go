package service

import (
	"fmt"
	"net/http"

	"github.com/barokurniawan/YoutubeDownloader/constant"
)

type Welcome struct {
	*RouteServiceProvide
}

func (welcome *Welcome) Route() {
	welcome.RouteServiceProvide.RegisterRoute("/", http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		var q = req.URL.Query()
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")

		if q.Get("u") == "" {
			fmt.Fprint(writer, "{\"info\" : false, \"message\" : \"usage: ?u=your+youtube+url&filter=audio(filter is optional)\"}")
			return
		}

		dl := YoutubeDL{}
		dl.SetYoutubeURL(q.Get("u"))

		output, err := dl.GetYoutubeInfo()
		if err != nil {
			fmt.Fprint(writer, fmt.Sprintf("{\"info\" : false, \"message\" : \"%s\"}", err.Error()))
			return
		}

		if q.Get("filter") == constant.FilterOnlyAudio {
			output = dl.GetAudioOnly()
		}

		fmt.Fprint(writer, output.ToJson())
	}))
}
