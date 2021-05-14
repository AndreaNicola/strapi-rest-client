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
	sur.ID = int(params["id"].(float64))
	sur.Name = params["name"].(string)
	sur.Description = params["description"].(string)
	sur.Type = params["type"].(string)
}

type StrapiUser struct {
	ID        int            `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Provider  string         `json:"provider"`
	Confirmed bool           `json:"confirmed"`
	Blocked   bool          `json:"blocked"`
	Role      StrapiUserRole `json:"role"`
	Disabled  bool           `json:"disabled"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func (su *StrapiUser) New(params map[string]interface{}) {

	su.ID = int(params["id"].(float64))
	su.Username = params["username"].(string)
	su.Email = params["email"].(string)
	su.Provider = params["provider"].(string)
	su.Disabled = params["disabled"].(bool)
	su.Confirmed = params["confirmed"].(bool)

	createdAt, t1Err := time.Parse(time.RFC3339, params["created_at"].(string))
	updatedAt, t2Err := time.Parse(time.RFC3339, params["updated_at"].(string))

	if t1Err == nil {
		su.CreatedAt = createdAt
	}

	if t2Err == nil {
		su.UpdatedAt = updatedAt
	}

	if blocked := params["blocked"]; blocked != nil {
		su.Blocked = blocked.(bool)
	}

	var role StrapiUserRole
	role.New(params["role"].(map[string]interface{}))

}


