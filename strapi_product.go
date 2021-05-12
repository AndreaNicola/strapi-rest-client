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
	spif.Width = int(params["width"].(float64))
	spif.Height = int(params["height"].(float64))
	spif.Size = params["height"].(float64)
	spif.Name = params["name"].(string)
	spif.Hash = params["hash"].(string)
	spif.Ext = params["ext"].(string)
	spif.Mime = params["mime"].(string)
	spif.Path = params["path"].(string)
	spif.URL = params["url"].(string)
}

type StrapiProductImageFormats struct {
	Thumbnail StrapiProductImageFormat `json:"thumbnail"`
	Large     StrapiProductImageFormat `json:"large"`
	Medium    StrapiProductImageFormat `json:"medium"`
	Small     StrapiProductImageFormat `json:"small"`
}

func (spifs *StrapiProductImageFormats) New(params map[string]interface{}) {

	var thumbnail StrapiProductImageFormat
	var large StrapiProductImageFormat
	var medium StrapiProductImageFormat
	var small StrapiProductImageFormat

	thumbnail.New(params["thumbnail"].(map[string]interface{}))
	large.New(params["large"].(map[string]interface{}))
	medium.New(params["medium"].(map[string]interface{}))
	small.New(params["small"].(map[string]interface{}))

	spifs.Thumbnail = thumbnail
	spifs.Large = large
	spifs.Medium = medium
	spifs.Small = small

}

type StrapiProductImage struct {
	ID              int                       `json:"id"`
	Width           int                       `json:"width"`
	Height          int                       `json:"height"`
	Size            float64                   `json:"size"`
	Name            string                    `json:"name"`
	AlternativeText string                    `json:"alternativeText"`
	Caption         string                    `json:"caption"`
	Hash            string                    `json:"hash"`
	Ext             string                    `json:"ext"`
	Mime            string                    `json:"mime"`
	URL             string                    `json:"url"`
	PreviewUrl      string                    `json:"previewUrl"`
	Provider        string                    `json:"provider"`
	CreatedAt       time.Time                 `json:"created_at"`
	UpdatedAt       time.Time                 `json:"updated_at"`
	Formats         StrapiProductImageFormats `json:"formats"`
}

func (spi *StrapiProductImage) New(params map[string]interface{}) {
	spi.ID = int(params["id"].(float64))
	spi.Width = int(params["width"].(float64))
	spi.Height = int(params["height"].(float64))
	spi.Size = params["height"].(float64)
	spi.Name = params["name"].(string)
	spi.AlternativeText = params["alternativeText"].(string)
	spi.Caption = params["caption"].(string)
	spi.Hash = params["hash"].(string)
	spi.Ext = params["ext"].(string)
	spi.Mime = params["mime"].(string)
	spi.URL = params["uRL"].(string)
	spi.PreviewUrl = params["previewUrl"].(string)
	spi.Provider = params["provider"].(string)
	createdAt, t1Err := time.Parse(time.RFC3339, params["created_at"].(string))
	updatedAt, t2Err := time.Parse(time.RFC3339, params["updated_at"].(string))

	if t1Err != nil {
		spi.CreatedAt = createdAt
	}

	if t2Err != nil {
		spi.UpdatedAt = updatedAt
	}

	var formats StrapiProductImageFormats
	formats.New(params["formats"].(map[string]interface{}))
	spi.Formats = formats

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
	PublishedAt    time.Time            `json:"published_at"`
	CreatedAt      time.Time            `json:"created_at"`
	UpdatedAt      time.Time            `json:"updated_at"`
	Image          []StrapiProductImage `json:"image"`
}

func (sp *StrapiProduct) New(params map[string]interface{}) {

	sp.ID = int(params["id"].(float64))
	sp.Quantity = int(params["quantity"].(float64))
	sp.Price = params["price"].(float64)

	sp.Title = params["title"].(string)
	sp.Description = params["description"].(string)
	sp.Abstract = params["abstract"].(string)
	sp.AdditionalInfo = params["additional_info"].(string)

	createdAt, t1Err := time.Parse(time.RFC3339, params["created_at"].(string))
	updatedAt, t2Err := time.Parse(time.RFC3339, params["updated_at"].(string))
	publishedAt, t3Err := time.Parse(time.RFC3339, params["published_at"].(string))

	if t1Err != nil {
		sp.CreatedAt = createdAt
	}

	if t2Err != nil {
		sp.UpdatedAt = updatedAt
	}

	if t3Err != nil {
		sp.PublishedAt = publishedAt
	}

	var image = make([]StrapiProductImage, 0)

	for _, e := range params["image"].([]interface{}) {

		imageMap := e.(map[string]interface{})
		var im StrapiProductImage
		im.New(imageMap)

		image = append(image, im)

	}

	sp.Image = image

}

