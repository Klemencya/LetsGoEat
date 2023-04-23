package pkg

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func ConnectDB() error {
	fmt.Println("in connectDB...")
	fmt.Println("\tconnecting DB...")
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		BDConfig["user"],
		BDConfig["password"],
		BDConfig["host"],
		BDConfig["port"],
		BDConfig["dbname"])

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println("DB connection failed in ConnectDB:", err)
		return err
	}
	fmt.Println("\tDB is connected")
	DB = db
	return nil
}

func checkConnection() error {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		BDConfig["user"],
		BDConfig["password"],
		BDConfig["host"],
		BDConfig["port"],
		BDConfig["dbname"])

	_, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println("DB connection failed in ConnectDB:", err)
		return err
	}
	fmt.Println("\tDB connection checked")
	return nil
}

func DisconnectDB() error {
	fmt.Println("in disconnectDB...")
	sql := DB.DB

	defer func() {
		err := sql.Close()
		if err != nil {
			fmt.Println("err in DisconnectDB while Close:", err)
			return
		}
	}()
	fmt.Println("\tDB connection is closed")
	return nil
}

func CreateUserInfoTable() error {
	fmt.Println("in CreateUserInfoTable...")
	err := checkConnection()
	if err != nil {
		fmt.Println("err in CreateUserInfoTable while checkConnection:", err)
		return err
	}

	var query = `
	CREATE TABLE IF NOT EXISTS user_info (
		user_id serial PRIMARY KEY,
		name VARCHAR ( 50 ) NOT NULL,
		surname VARCHAR ( 50 ) NOT NULL,
		email VARCHAR ( 50 ) NOT NULL,
		login VARCHAR ( 50 ) UNIQUE NOT NULL,
		password VARCHAR ( 50 ) NOT NULL,
		preferences TEXT
	);`

	_, err = DB.Exec(query)
	if err != nil {
		fmt.Println("err in CreateUserInfoTable while Exec(query):", err)
		return err
	}

	fmt.Println("\tuser_info table is ready")
	return nil
}

func CreateRequestsTable() error {
	fmt.Println("in CreateRequestsTable...")
	err := checkConnection()
	if err != nil {
		fmt.Println("err in CreateRequestsTable while checkConnection:", err)
		return err
	}

	var query = `
	CREATE TABLE IF NOT EXISTS requests (
		req_id serial PRIMARY KEY,
		senderUser VARCHAR ( 50 ) NOT NULL,
		receiverUser VARCHAR ( 50 ) NOT NULL,
		place VARCHAR ( 50 ) NOT NULL,
		cuisine VARCHAR ( 50 ) NOT NULL,
		invitation VARCHAR ( 50 ) NOT NULL,
		accepted BOOLEAN NOT NULL
	);`

	_, err = DB.Exec(query)
	if err != nil {
		fmt.Println("err in CreateRequestsTable while Exec(query):", err)
		return err
	}

	fmt.Println("\trequests table is ready")
	return nil
}

func GetRequests(login string) ([]InvColumn, error) {
	fmt.Println("in GetRequest...")
	var res []InvColumn
	err := checkConnection()
	if err != nil {
		fmt.Println("err in GetRequest while checkConnection:", err)
		return res, err
	}

	query := `SELECT * FROM requests WHERE receiveruser = $1`
	var invColumn []InvColumn
	err = DB.Select(&invColumn, query, login)
	if err != nil {
		fmt.Println("err in GetRequest while Exec(query):", err)
		return res, err
	}

	return invColumn, nil
}

func GetUser(login string) (UserColumn, error) {
	fmt.Println("in GetUser...")
	var res UserColumn
	err := checkConnection()
	if err != nil {
		fmt.Println("err in GetUser while checkConnection:", err)
		return res, err
	}

	query := `SELECT * FROM user_info WHERE login = $1`
	var userColumn []UserColumn
	err = DB.Select(&userColumn, query, login)
	if err != nil {
		fmt.Println("err in LogInCheck while Exec(query):", err)
		return res, err
	}

	return userColumn[0], nil
}

func LogInCheck(user UserLogin) error {
	fmt.Println("in LogInCheck...")
	err := checkConnection()
	if err != nil {
		fmt.Println("err in LogInCheck while checkConnection:", err)
		return err
	}

	query := `SELECT * FROM user_info WHERE login = $1`
	var userColumn []UserColumn
	err = DB.Select(&userColumn, query, user.Login)
	if err != nil {
		fmt.Println("err in LogInCheck while Exec(query):", err)
		return err
	}

	if len(userColumn) == 0 {
		print("err in LogInCheck: user does not exist")
		return errors.New("err in LogInCheck: user does not exist")
	} else if userColumn[0].Password != user.Password {
		print("err in LogInCheck: wrong password")
		return errors.New("err in LogInCheck: wrong password")
	}

	return nil
}

func InsertUser(user UserRegistration) error {
	fmt.Println("in InsertUser...")
	err := checkConnection()
	if err != nil {
		fmt.Println("err in InsertUser while checkConnection:", err)
		return err
	}

	query := `INSERT INTO user_info (name, surname, email, login, password, preferences) VALUES ($1, $2, $3, $4, $5, '');`
	_, err = DB.Exec(query, user.Name, user.Surname, user.Email, user.Login, user.Password)
	if err != nil {
		fmt.Println("err in InsertUser while Exec(query):", err)
		return err
	}

	fmt.Println("\tuser " + user.Login + " has been inserted into user_info")
	return nil
}

func InsertInvitation(inv Invitation) error {
	fmt.Println("in InsertInvitation...")
	err := checkConnection()
	if err != nil {
		fmt.Println("err in InsertInvitation while checkConnection:", err)
		return err
	}

	query := `INSERT INTO requests (senderuser, receiveruser, place, cuisine, invitation, accepted) VALUES ($1, $2, $3, $4, $5, $6);`
	_, err = DB.Exec(query, inv.SenderUser, inv.ReceiverUser, inv.Place, inv.Cuisine, inv.Invitation, false)
	if err != nil {
		fmt.Println("err in InsertInvitation while Exec(query):", err)
		return err
	}

	fmt.Println("\tinvitation from " + inv.SenderUser + " has been inserted into requests")
	return nil
}

func GetAllUsers() ([]UserColumn, error) {
	fmt.Println("in GetAllUsers...")
	var res []UserColumn
	err := checkConnection()
	if err != nil {
		fmt.Println("err in GetAllUsers while checkConnection:", err)
		return res, err
	}

	query := `SELECT * FROM user_info`
	var uc []UserColumn
	err = DB.Select(&uc, query)
	if err != nil {
		fmt.Println("err in GetAllUsers while Exec(query):", err)
		return res, err
	}

	return uc, nil
}

func DeleteRequest(id string) error {
	fmt.Println("in GetRequest...")
	err := checkConnection()
	if err != nil {
		fmt.Println("err in DeleteRequest while checkConnection:", err)
		return err
	}

	fmt.Println(id, len(id))
	query := `DELETE FROM requests WHERE req_id = $1`
	_, err = DB.Exec(query, id)
	if err != nil {
		fmt.Println("err in DeleteRequest while Exec(query):", err)
		return err
	}

	return nil
}

func ChangeAccepted(id string) error {
	fmt.Println("in ChangeAccepted...")
	err := checkConnection()
	if err != nil {
		fmt.Println("err in ChangeAccepted while checkConnection:", err)
		return err
	}

	query := `SELECT * FROM requests WHERE req_id = $1`
	var ic []InvColumn
	err = DB.Select(&ic, query, id)
	if err != nil {
		fmt.Println("err in ChangeAccepted while Exec(query):", err)
		return err
	}

	if len(ic) > 0 {
		query = `UPDATE requests SET accepted = $1 WHERE req_id = $2`
		acc := ic[0].Accepted
		_, err = DB.Exec(query, !acc, id)
		if err != nil {
			fmt.Println("err in ChangeAccepted while Exec(query):", err)
			return err
		}
	}
	return nil
}
