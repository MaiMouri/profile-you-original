package service

type SignupService interface {
	SignupUser(email string, password string) bool
	Register_proc(email string, password string) string
}
type signupInformation struct {
	email    string
	password string
}

// func (info *loginInformation) Register_proc(email string, password string) bool {
// 	return info.email == email && info.password == password
// }

// func Register_proc(email string, password string) string {
// 	fmt.Fprintf(os.Stderr, "*** Register_proc ***\n")

// 	now := time.Now()

// 	str_out := JWTService.Generate_token_proc(email, now)
// 	db := sqlite.New()
// 	// db, _ := gorm.Open("mysql", "scott:tiger123@/test_db")
// 	connect, err := db.DB()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer connect.Close()

// 	sql_str := "insert into users (email,password,created_at) values (?,?,?)"
// 	db.Exec(sql_str, email, password, now)

// 	return str_out
// }
