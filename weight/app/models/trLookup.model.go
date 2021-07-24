package models

type TrLookup struct {
	Code         string `gorm:"column:code" json:"code,omitempty"`
	Value        string `gorm:"column:value" json:"value,omitempty"`
	Description  string `gorm:"column:description" json:"description,omitempty"`
	Notes        string `gorm:"column:notes" json:"notes,omitempty"`
	StatusId     string `gorm:"column:status_id;default:1" json:"status_id,omitempty"`
	CreatedBy    string `gorm:"column:created_by;" json:"created_by,omitempty"`
	CreatedDate  string `gorm:"column:created_date;" json:"created_date,omitempty"`
	CreatedName  string `gorm:"column:created_name;" json:"created_name,omitempty"`
	CreatedFrom  string `gorm:"column:created_from;" json:"created_from,omitempty"`
	ModifiedBy   string `gorm:"column:modified_by;" json:"modified_by,omitempty"`
	ModifiedDate string `gorm:"column:modified_date;" json:"modified_date,omitempty"`
	ModifiedName string `gorm:"column:modified_name;" json:"modified_name,omitempty"`
	ModifiedFrom string `gorm:"column:modified_from;" json:"modified_from,omitempty"`
}

func (TrLookup) TableName() string {
	return "tr_lookup"
}
