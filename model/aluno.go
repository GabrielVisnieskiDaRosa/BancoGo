package model

type Aluno struct{
	Id int  `gorm:"primaryKey"`
	Nome string`gorm:"type:varchar(60); not null"`
	Ra string `gorm:"type:varchar(60); not null"`
}

