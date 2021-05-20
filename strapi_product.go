package strapi_rest_client

import (
	"time"
)

type StrapiProductImageFormat struct {
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Size   float64 `json:"size"`
	Name   string  `json:"name"`
	Hash   string  `json:"hash"`
	Ext    string  `json:"ext"`
	Mime   string  `json:"mime"`
	Path   string  `json:"path"`
	URL    string  `json:"url"`
}

func (spif *StrapiProductImageFormat) New(params map[string]interface{}) {

	if val, ok := params["width"]; ok && val != nil {
		spif.Width = int(val.(float64))
	}

	if val, ok := params["height"]; ok && val != nil {
		spif.Height = int(val.(float64))
	}

	if val, ok := params["height"]; ok && val != nil {
		spif.Size = val.(float64)
	}

	if val, ok := params["name"]; ok && val != nil {
		spif.Name = val.(string)
	}

	if val, ok := params["hash"]; ok && val != nil {
		spif.Hash = val.(string)
	}

	if val, ok := params["ext"]; ok && val != nil {
		spif.Ext = val.(string)
	}

	if val, ok := params["mime"]; ok && val != nil {
		spif.Mime = val.(string)
	}

	if val, ok := params["path"]; ok && val != nil {
		spif.Path = val.(string)
	}

	if val, ok := params["url"]; ok && val != nil {
		spif.URL = val.(string)
	}

}

type StrapiProductImageFormats struct {
	Thumbnail *StrapiProductImageFormat `json:"thumbnail"`
	Large     *StrapiProductImageFormat `json:"large"`
	Medium    *StrapiProductImageFormat `json:"medium"`
	Small     *StrapiProductImageFormat `json:"small"`
}

func (spifs *StrapiProductImageFormats) New(params map[string]interface{}) {

	if t, ok := params["thumbnail"]; ok && t != nil {
		spifs.Thumbnail = &StrapiProductImageFormat{}
		spifs.Thumbnail.New(t.(map[string]interface{}))
	}

	if l, ok := params["large"]; ok && l != nil {
		spifs.Large = &StrapiProductImageFormat{}
		spifs.Large.New(l.(map[string]interface{}))
	}

	if m, ok := params["medium"]; ok && m != nil {
		spifs.Medium = &StrapiProductImageFormat{}
		spifs.Medium.New(m.(map[string]interface{}))
	}

	if s, ok := params["small"]; ok && s != nil {
		spifs.Small = &StrapiProductImageFormat{}
		spifs.Small.New(s.(map[string]interface{}))
	}

}

type StrapiProductImage struct {
	ID              int                        `json:"id"`
	Width           int                        `json:"width"`
	Height          int                        `json:"height"`
	Size            float64                    `json:"size"`
	Name            string                     `json:"name"`
	AlternativeText string                     `json:"alternativeText"`
	Caption         string                     `json:"caption"`
	Hash            string                     `json:"hash"`
	Ext             string                     `json:"ext"`
	Mime            string                     `json:"mime"`
	URL             string                     `json:"url"`
	PreviewUrl      string                     `json:"previewUrl"`
	Provider        string                     `json:"provider"`
	CreatedAt       time.Time                  `json:"created_at"`
	UpdatedAt       time.Time                  `json:"updated_at"`
	Formats         *StrapiProductImageFormats `json:"formats"`
}

func (spi *StrapiProductImage) New(params map[string]interface{}) {

	if val, ok := params["id"]; ok && val != nil {
		spi.ID = int(val.(float64))
	}

	if val, ok := params["width"]; ok && val != nil {
		spi.Width = int(val.(float64))
	}

	if val, ok := params["height"]; ok && val != nil {
		spi.Height = int(val.(float64))
	}

	if val, ok := params["height"]; ok && val != nil {
		spi.Size = val.(float64)
	}

	if val, ok := params["name"]; ok && val != nil {
		spi.Name = val.(string)
	}

	if val, ok := params["alternativeText"]; ok && val != nil {
		spi.AlternativeText = val.(string)
	}

	if val, ok := params["caption"]; ok && val != nil {
		spi.Caption = val.(string)
	}

	if val, ok := params["hash"]; ok && val != nil {
		spi.Hash = val.(string)
	}

	if val, ok := params["ext"]; ok && val != nil {
		spi.Ext = val.(string)
	}

	if val, ok := params["mime"]; ok && val != nil {
		spi.Mime = val.(string)
	}

	if val, ok := params["url"]; ok && val != nil {
		spi.URL = val.(string)
	}

	if val, ok := params["previewUrl"]; ok && val != nil {
		spi.PreviewUrl = val.(string)
	}

	if val, ok := params["previewUrl"]; ok && val != nil {
		spi.Provider = val.(string)
	}

	if val, ok := params["created_at"]; ok && val != nil {
		createdAt, t1Err := time.Parse(time.RFC3339, val.(string))
		if t1Err != nil {
			spi.CreatedAt = createdAt
		}
	}

	if val, ok := params["updated_at"]; ok && val != nil {
		updatedAt, t2Err := time.Parse(time.RFC3339, val.(string))
		if t2Err != nil {
			spi.UpdatedAt = updatedAt
		}
	}

	if val, ok := params["formats"]; ok && val != nil {
		spi.Formats = &StrapiProductImageFormats{}
		spi.Formats.New(val.(map[string]interface{}))
	}

}

type StrapiProduct struct {
	ID             int                  `json:"id"`
	Title          string               `json:"title"`
	Description    string               `json:"description"`
	Category       interface{}          `json:"category"`
	Abstract       string               `json:"abstract"`
	Price          float64              `json:"price"`
	AdditionalInfo string               `json:"additional_info"`
	Quantity       int                  `json:"quantity"`
	Permalink      string               `json:"permalink"`
	PublishedAt    time.Time            `json:"published_at"`
	CreatedAt      time.Time            `json:"created_at"`
	UpdatedAt      time.Time            `json:"updated_at"`
	Image          []StrapiProductImage `json:"image"`
}

func (sp *StrapiProduct) New(params map[string]interface{}) {

	sp.ID = int(params["id"].(float64))
	sp.Quantity = int(params["quantity"].(float64))
	sp.Price = params["price"].(float64)

	if val, ok := params["title"]; ok && val != nil {
		sp.Title = val.(string)
	}

	if val, ok := params["description"]; ok && val != nil {
		sp.Description = val.(string)
	}

	if val, ok := params["abstract"]; ok && val != nil {
		sp.Abstract = val.(string)
	}

	if val, ok := params["additional_info"]; ok && val != nil {
		sp.AdditionalInfo = val.(string)
	}

	if val, ok := params["permalink"]; ok && val != nil {
		sp.Permalink = val.(string)
	}

	if val, ok := params["created_at"]; ok && val != nil {
		createdAt, t1Err := time.Parse(time.RFC3339, val.(string))
		if t1Err != nil {
			sp.CreatedAt = createdAt
		}
	}

	if val, ok := params["updated_at"]; ok && val != nil {
		updatedAt, t2Err := time.Parse(time.RFC3339, val.(string))
		if t2Err != nil {
			sp.UpdatedAt = updatedAt
		}
	}

	if val, ok := params["published_at"]; ok && val != nil {
		publishedAt, t2Err := time.Parse(time.RFC3339, val.(string))
		if t2Err != nil {
			sp.PublishedAt = publishedAt
		}
	}

	sp.Image = make([]StrapiProductImage, 0)

	if images, ok := params["image"]; ok && images != nil {
		for _, e := range images.([]interface{}) {
			imageMap := e.(map[string]interface{})
			var im StrapiProductImage
			im.New(imageMap)
			sp.Image = append(sp.Image, im)
		}
	}

}
