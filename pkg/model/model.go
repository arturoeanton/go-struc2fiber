package model

import "encoding/json"

type InternalUser struct {
	ID                  int64     `json:"id,omitempty"  gorm:"primaryKey"` // Agrega el campo ID
	UserID              string    `json:"user_id,omitempty" gorm:"column:user_id"`
	AvatarURL           string    `json:"avatar_url,omitempty" gorm:"column:avatar_url"`
	BgURL               string    `json:"bg_url,omitempty" gorm:"column:bg_url"`
	LastName            string    `json:"last_name,omitempty" gorm:"column:last_name"`
	FirstName           string    `json:"first_name,omitempty" gorm:"column:first_name"`
	Author              *string   `json:"author,omitempty" gorm:"column:author"`
	LastUpdate          *string   `json:"last_update,omitempty" gorm:"column:last_update"`
	LastUpdateBy        *string   `json:"last_update_by,omitempty" gorm:"column:last_update_by"`
	Username            string    `json:"username" gorm:"column:username"  `
	Password            string    `json:"password" gorm:"column:user_password"`
	Email               string    `json:"email" gorm:"column:email"`
	Phone               string    `json:"phone" gorm:"column:phone"`
	Role                string    `json:"role" gorm:"column:user_role"`
	ApiKey              string    `json:"api_key" gorm:"column:api_key"`
	SecretKey           string    `json:"-" gorm:"column:secret_key"`
	Auth                string    `json:"-" gorm:"column:auth"`
	Theme               string    `json:"theme" gorm:"column:theme"`
	Sound               string    `json:"sound" gorm:"column:sound"`
	Active              string    `json:"active" gorm:"column:active"`
	Forgot              string    `json:"forgot" gorm:"column:forgot"`
	CountBadLogin       int       `json:"count_bad_login" gorm:"column:count_bad_login"`
	Bio                 string    `json:"bio" gorm:"column:bio"`
	Tags                string    `json:"tags" gorm:"column:tags"`
	Skills              []Skill   `json:"skills,omitempty" gorm:"foreignKey:UserID"`                     // GORM: especifica la clave foránea
	Connects            []Connect `json:"connects,omitempty" gorm:"foreignKey:UserID"`                   // GORM: especifica la clave foránea
	DayOffs             []DayOff  `json:"day_offs,omitempty" gorm:"foreignKey:UserID"`                   // GORM: especifica la clave foránea
	CheckinsOfOperator  []Checkin `json:"checkins_of_operator,omitempty" gorm:"foreignKey:OperatorID"`   // GORM: especifica la clave foránea
	CheckinsOfCheckiner []Checkin `json:"checkins_of_checkiner,omitempty" gorm:"foreignKey:CheckinerID"` // GORM: especifica la clave foránea
	CheckinsOfAgent     []Checkin `json:"checkins_of_agent,omitempty" gorm:"foreignKey:AgentID"`         // GORM: especifica la clave foránea
}

func (InternalUser) TableName() string {
	return "internal_user"
}

type Skill struct {
	ID     int64  `json:"id,omitempty"  gorm:"primaryKey"` // Agrega el campo ID
	UserID int64  `json:"user_id,omitempty" gorm:"column:user_id"`
	Name   string `json:"name,omitempty" gorm:"column:name"`
	Value  int    `json:"value,omitempty" gorm:"column:value"`
}

func (Skill) TableName() string {
	return "skill"
}

type Connect struct {
	ID     int64  `json:"id,omitempty"  gorm:"primaryKey"` // Agrega el campo ID
	UserID int64  `json:"user_id,omitempty" gorm:"column:user_id"`
	Name   string `json:"name,omitempty" gorm:"column:name"`
	Title  string `json:"title,omitempty" gorm:"column:title"`
	Value  string `json:"value,omitempty" gorm:"column:value"`
}

func (Connect) TableName() string {
	return "connect"
}

type DayOff struct {
	ID          int64  `json:"id,omitempty"  gorm:"primaryKey"` // Agrega el campo ID
	UserID      int64  `json:"user_id,omitempty" gorm:"column:user_id"`
	Start       string `json:"start,omitempty" gorm:"column:start"`
	End         string `json:"end,omitempty" gorm:"column:end"`
	Description string `json:"description,omitempty" gorm:"column:description"`
}

func (DayOff) TableName() string {
	return "day_off"
}

type Apartment struct {
	ID          int64        `json:"id,omitempty"  gorm:"primaryKey"` // Agrega el campo ID
	AgentID     int64        `json:"agent_id,omitempty" gorm:"column:agent_id"`
	Agent       InternalUser `json:"agent,omitempty" gorm:"foreignKey:AgentID"` // GORM: especifica la clave foránea
	Name        string       `json:"name,omitempty" gorm:"column:name"`
	Description string       `json:"description,omitempty" gorm:"column:description"`
	Rooms       int          `json:"rooms,omitempty" gorm:"column:rooms"`
	Price       int          `json:"price,omitempty" gorm:"column:price"`
	Location    string       `json:"location,omitempty" gorm:"column:location"`
	Images      string       `json:"images,omitempty" gorm:"column:images"`
}

func (Apartment) TableName() string {
	return "apartment"
}

type Checkin struct {
	ID          int64        `json:"id,omitempty"  gorm:"primaryKey"` // Agrega el campo ID
	OperatorID  int64        `json:"operator_id,omitempty" gorm:"column:operator_id"`
	Operator    InternalUser `json:"operator,omitempty" gorm:"foreignKey:OperatorID"` // GORM: especifica la clave foránea
	CheckinerID int64        `json:"checkiner_id,omitempty" gorm:"column:checkiner_id"`
	Checkiner   InternalUser `json:"checkiner,omitempty" gorm:"foreignKey:CheckinerID"` // GORM: especifica la clave foránea
	AgentID     int64        `json:"agent_id,omitempty" gorm:"column:agent_id"`
	Agent       InternalUser `json:"agent,omitempty" gorm:"foreignKey:AgentID"` // GORM: especifica la clave foránea
	StartDate   string       `json:"start_date,omitempty" gorm:"column:start_date"`
	ApartmentID int64        `json:"apartment_id,omitempty" gorm:"column:apartment_id"`
	Apartment   Apartment    `json:"apartment,omitempty" gorm:"foreignKey:ApartmentID"` // GORM: especifica la clave foránea
	EndDate     string       `json:"end_date,omitempty" gorm:"column:end_date"`
	Comment     string       `json:"comment,omitempty" gorm:"column:comment"`
}

func (Checkin) TableName() string {
	return "checkin"
}

type Event struct {
	ID              *int64         `json:"id,omitempty" gorm:"primaryKey"`
	Author          *string        `json:"author,omitempty" gorm:"column:author"`
	LastUpdate      *string        `json:"last_update,omitempty" gorm:"column:last_update"`
	LastUpdateBy    *string        `json:"last_update_by,omitempty" gorm:"column:last_update_by"`
	Title           string         `json:"title" gorm:"column:title" `
	StartDate       string         `json:"start" gorm:"column:start_date"`
	EndDate         string         `json:"end" gorm:"column:end_date"`
	ExtendedPropsID *int64         `json:"-" gorm:"column:extended_props_id"`
	ExtendedProps   *ExtendedProps `json:"extendedProps,omitempty" gorm:"foreignKey:ExtendedPropsID"` // GORM: especifica la clave foránea

}

func (Event) TableName() string {
	return "event"
}

type ExtendedProps struct {
	ID           *int64                   `json:"id,omitempty" gorm:"primaryKey"`
	Author       *string                  `json:"author,omitempty" gorm:"column:author"`
	LastUpdate   *string                  `json:"last_update,omitempty" gorm:"column:last_update"`
	LastUpdateBy *string                  `json:"last_update_by,omitempty" gorm:"column:last_update_by"`
	Description  string                   `json:"description" gorm:"column:description" `
	Email        string                   `json:"email" gorm:"column:email"`
	HeaderString string                   `json:"-" gorm:"column:header"` // Campo para la base de datos
	ItemsString  string                   `json:"-" gorm:"column:items"`  // Campo para la base de datos
	Header       map[string]interface{}   `json:"header" gorm:"-"`        // GORM: ignora este campo
	Items        []map[string]interface{} `json:"items" gorm:"-"`         // GORM: ignora este campo

}

func (ExtendedProps) TableName() string {
	return "extended_props"
}

// MarshalJSON personaliza la serialización de ExtendedProps.
func (e *ExtendedProps) MarshalJSON() ([]byte, error) {
	type Alias ExtendedProps

	// Serializar el campo Items a JSON y almacenarlo en ItemsString
	err := json.Unmarshal([]byte(e.HeaderString), &e.Header)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(e.ItemsString), &e.Items)
	if err != nil {
		return nil, err
	}

	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(e),
	})
}

// UnmarshalJSON personaliza la deserialización de ExtendedProps.
func (e *ExtendedProps) UnmarshalJSON(data []byte) error {
	type Alias ExtendedProps
	aux := &struct {
		Items  []map[string]interface{} `json:"items"`
		Header map[string]interface{}   `json:"header"`
		*Alias
	}{
		Alias: (*Alias)(e),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	dataItems, err := json.Marshal(aux.Items)
	if err != nil {
		return err
	}

	dataHeader, err := json.Marshal(aux.Header)
	if err != nil {
		return err
	}
	e.ItemsString = string(dataItems)
	e.HeaderString = string(dataHeader)
	e.Items = aux.Items

	return nil
}
