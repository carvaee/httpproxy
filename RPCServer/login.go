package RPCServer
type User struct {
	UserName string
	NickName string
	Image string
}

func Login(userName, password string) (user *User,err error){
	user =&User{UserName: "myName",
		NickName: "little",
		Image: "./HTTPServer/template/cut.jpg",
	}
	return  user, err
}
