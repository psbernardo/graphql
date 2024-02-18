package database

type KeyValue struct {
	Key  string           `gorm:"type:varchar(25)"`
	Data QueryResultJSONB `gorm:"type:jsonb;default:'[]';not null"`
}

func (k *KeyValue) TableName() string {
	return "Key_Value"
}
