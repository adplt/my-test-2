package structs

import (
	"github.com/google/uuid"
)

type TxWeightRecord struct {
	WeightRecordId uuid.UUID `json:"weight_record_id,omitempty" swaggertype:"string" example:"ddccc2ad-4ebe-4b59-92b1-8417d5e5a0cf"`
	Date           string    `json:"date,omitempty" swaggertype:"string" example:"2021-07-24"`
	Max            int64     `json:"max,omitempty" swaggertype:"string" example:"55"`
	Min            string    `json:"min,omitempty" swaggertype:"string" example:"50"`
	Differences    int       `gorm:"column:differences" json:"differences,omitempty"`
	StatusId       string    `json:"status_id,omitempty" swaggertype:"string" example:"1"`
	CreatedBy      string    `json:"created_by,omitempty" swaggertype:"string" example:"DBeaver"`
	CreatedDate    string    `json:"created_date,omitempty" swaggertype:"string" example:"2021-07-20 09:34:12"`
	CreatedName    string    `json:"created_name,omitempty" swaggertype:"string" example:"DBeaver"`
	CreatedFrom    string    `json:"created_from,omitempty" swaggertype:"string" example:"DBeaver"`
	ModifiedBy     string    `json:"modified_by,omitempty" swaggertype:"string" example:"DBeaver"`
	ModifiedDate   string    `json:"modified_date,omitempty" swaggertype:"string" example:"2021-07-20 09:34:12"`
	ModifiedName   string    `json:"modified_name,omitempty" swaggertype:"string" example:"DBeaver"`
	ModifiedFrom   string    `json:"modified_from,omitempty" swaggertype:"string" example:"DBeaver"`
}

func (TxWeightRecord) TableName() string {
	return "tx_weight_record"
}
