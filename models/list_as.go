package models

type ListAs struct {
	ID        int       `json:"id"`
	Role      string    `json:"role" grom:"type: varchar(255)"`

}

type ListAsResponse struct {
	ID   int    `json:"id"`
	Role string `json:"role" grom:"type: varchar(255)"`
}

func (ListAsResponse) TableName() string {
	return "List_as"
}
