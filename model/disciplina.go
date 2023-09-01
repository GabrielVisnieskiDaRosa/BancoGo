package model



type Disciplina struct {
	Id          int    `gorm:"primaryKey"`
	Nome        string `gorm:"type:varchar(60); not null"`
	Atualizado  bool   `gorm:"-"`
//	IdProfessor int    //Chave estrangeira para o professor
//	Professor   Professor `gorm:"foreingKey:IdProfessor` //Constroi a relação de chave estrangeira

	ProfessorId int
	Professor Professor

	Alunos [] Aluno `gorm:"many2many:disciplina_aluno"`

}
