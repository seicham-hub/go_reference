package ducktype

import "fmt"

type (
    // Adインタフェース
    Ad interface {
        GetAdType() int64
    }
    Ads []Ad
    // AdVideo ad video response struct
    AdVideo struct {
        VideoURL string `json:"video_url"`
        AdType   int64  `json:"ad_type"`
    }
    // AdPoster ad poster response struct
    AdPoster struct {
        PosterURL string `json:"poster_url"`
        AdType    int64  `json:"ad_type"`
    }
)

func main() {
    var ads Ads
    ads = append(ads, NewAdVideo())
    ads = append(ads, NewAdPoster())
    for _, ad := range ads {
        fmt.Println(ad)
    }
}

func NewAdVideo() AdVideo {
    return AdVideo{VideoURL: "video url", AdType: int64(1)}
}

// => Adのインタフェース型を持つために、`GetAdType`を実装した。
func (v AdVideo) GetAdType() int64 {
    return v.AdType
}

func NewAdPoster() AdPoster {
    return AdPoster{PosterURL: "poster url", AdType: int64(3)}
}

// => Adのインタフェース型を持つために、`GetAdType`を実装した。
func (p AdPoster) GetAdType() int64 {
    return p.AdType