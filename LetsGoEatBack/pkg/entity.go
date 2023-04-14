package pkg

import (
	"github.com/jmoiron/sqlx"
	"time"
)

var ResponseID = 0
var UserID = 0
var DB *sqlx.DB

var BDConfig = map[string]string{
	"user":     "postgres",
	"password": "1308249756",
	"host":     "localhost",
	"port":     "15432",
	"dbname":   "myProjDB",
}

func GetResponseID() int {
	ResponseID += 1
	return ResponseID
}

func MakeResp(title string, msg string) Response {
	return Response{Id: GetResponseID(), Title: title, Date: time.Now().Unix(), Msg: msg}
}

type Response struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Date  int64  `json:"date"`
	Msg   string `json:"msg"`
}

type UserRegistration struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserColumn struct {
	ID       string `db:"user_id"`
	Name     string `db:"name"`
	Surname  string `db:"surname"`
	Email    string `db:"email"`
	Login    string `db:"login"`
	Password string `db:"password"`
}

type User struct {
	Name     string
	Surname  string
	Email    string
	Login    string
	Password string
}
