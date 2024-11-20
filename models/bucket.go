package models

import (
	"time"
)

type Bucket struct {
	ID        uint     `gorm:"primaryKey;autoIncrement:true,index"`
	Name      string   `gorm:"unique;not null"`
	Objects   []Object `gorm:"foreignKey:BucketID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BucketMetadata struct {
	ID string `msgpack:"id"`
}
