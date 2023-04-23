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

func MakeUserResp(users string, title string, msg string) UserResponse {
	return UserResponse{Id: GetResponseID(), Title: title, Date: time.Now().Unix(), Msg: msg, Users: users}
}

type Response struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Date  int64  `json:"date"`
	Msg   string `json:"msg"`
}

type UserResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Date  int64  `json:"date"`
	Msg   string `json:"msg"`
	Users string `json:"users"`
}

type GetUserResponse struct {
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

type Invitation struct {
	SenderUser   string `json:"sender_user"`
	ReceiverUser string `json:"receiver_user"`
	Place        string `json:"place"`
	Cuisine      string `json:"cuisine"`
	Invitation   string `json:"invitation"`
	Accepted     bool   `json:"accepted"`
}

type UserLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Login struct {
	Login string `json:"login"`
}

type ID struct {
	ID string `json:"id"`
}

type UserColumn struct {
	ID          string `db:"user_id"`
	Name        string `db:"name"`
	Surname     string `db:"surname"`
	Email       string `db:"email"`
	Login       string `db:"login"`
	Password    string `db:"password"`
	Preferences string `db:"preferences"`
}

type InvColumn struct {
	ID         string `db:"req_id"`
	Sender     string `db:"senderuser"`
	Receiver   string `db:"receiveruser"`
	Place      string `db:"place"`
	Cuisine    string `db:"cuisine"`
	Invitation string `db:"invitation"`
	Accepted   bool   `db:"accepted"`
}

type User struct {
	Name     string
	Surname  string
	Email    string
	Login    string
	Password string
}
