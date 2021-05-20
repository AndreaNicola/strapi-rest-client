package strapi_rest_client

import (
	"time"
)

type StrapiUserRole struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

func (sur *StrapiUserRole) New(params map[string]interface{}) {

	if val, ok := params["id"]; ok && val != nil {
		sur.ID = int(val.(float64))
	}

	if val, ok := params["name"]; ok && val != nil {
		sur.Name = val.(string)
	}

	if val, ok := params["description"]; ok && val != nil {
		sur.Description = val.(string)
	}

	if val, ok := params["type"]; ok && val != nil {
		sur.Type = val.(string)
	}

}

type StrapiUser struct {
	ID        int             `json:"id"`
	Username  string          `json:"username"`
	Email     string          `json:"email"`
	Provider  string          `json:"provider"`
	Confirmed bool            `json:"confirmed"`
	Blocked   bool            `json:"blocked"`
	Role      *StrapiUserRole `json:"role"`
	Disabled  bool            `json:"disabled"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func (su *StrapiUser) New(params map[string]interface{}) {

	if val, ok := params["id"]; ok && val != nil {
		su.ID = int(val.(float64))
	}

	if val, ok := params["username"]; ok && val != nil {
		su.Username = val.(string)
	}

	if val, ok := params["email"]; ok && val != nil {
		su.Email = val.(string)
	}

	if val, ok := params["provider"]; ok && val != nil {
		su.Provider = val.(string)
	}

	if val, ok := params["disabled"]; ok && val != nil {
		su.Disabled = val.(bool)
	}

	if val, ok := params["confirmed"]; ok && val != nil {
		su.Confirmed = val.(bool)
	}

	if val, ok := params["created_at"]; ok && val != nil {
		createdAt, t1Err := time.Parse(time.RFC3339, val.(string))
		if t1Err == nil {
			su.CreatedAt = createdAt
		}
	}

	if val, ok := params["updated_at"]; ok && val != nil {
		updatedAt, t2Err := time.Parse(time.RFC3339, val.(string))
		if t2Err == nil {
			su.UpdatedAt = updatedAt
		}
	}

	if blocked := params["blocked"]; blocked != nil {
		su.Blocked = blocked.(bool)
	}

	if val, ok := params["role"]; ok && val != nil {
		su.Role = &StrapiUserRole{}
		su.Role.New(val.(map[string]interface{}))
	}

}
