package models

type User struct {
	Email    string `bson:"email"`
	Name     string `bson:"name"`
	Age      int    `bson:"age"`
	Password string `bson:"password"`
}

func (u *User) VerifyEmail() {
	// TODO: Enviar email para o usuario verificando o email dele
}

func (u *User) CheckDuplicateEmail() {
	// TODO: Verificar se o email ja nao existe no banco de dados
}
