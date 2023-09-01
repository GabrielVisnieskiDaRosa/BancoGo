package model

type Professor struct {
	Id        int    `gorm:"primaryKey"`
	Nome      string `gorm:"type:varchar(60); not null"`
	Titulacao string `gorm:"type:varchar(60); not null"`
}
