package response

type WeightResponse struct {
	WeightRecordId int    `gorm:"primaryKey;column:weight_record_id;" json:"weight_record_id,omitempty"`
	Date           string `gorm:"column:date" json:"date,omitempty"`
	Max            int    `gorm:"column:max" json:"max,omitempty"`
	Min            int    `gorm:"column:min" json:"min,omitempty"`
	StatusId       string `gorm:"column:status_id;default:1" json:"status_id,omitempty"`
	CreatedBy      string `gorm:"column:created_by;" json:"created_by,omitempty"`
	CreatedDate    string `gorm:"column:created_date;" json:"created_date,omitempty"`
	CreatedName    string `gorm:"column:created_name;" json:"created_name,omitempty"`
	CreatedFrom    string `gorm:"column:created_from;" json:"created_from,omitempty"`
	ModifiedBy     string `gorm:"column:modified_by;" json:"modified_by,omitempty"`
	ModifiedDate   string `gorm:"column:modified_date;" json:"modified_date,omitempty"`
	ModifiedName   string `gorm:"column:modified_name;" json:"modified_name,omitempty"`
	ModifiedFrom   string `gorm:"column:modified_from;" json:"modified_from,omitempty"`
}
