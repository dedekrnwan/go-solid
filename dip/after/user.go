package after

type IUser interface {
	GetSender() *User
	GetReceiver() *User
}

type User struct {
	id    int
	name  string
	phone int
	email string
}

func (u *User) GetSender() *User {
	return &User{
		id:    1,
		name:  "bambang",
		phone: 6281,
		email: "bambang@gmail.com",
	}
}

func (u *User) GetReceiver() *User {
	return &User{
		id:    2,
		name:  "Fajrian",
		phone: 6287,
		email: "fajrian@gmail.com",
	}
}
