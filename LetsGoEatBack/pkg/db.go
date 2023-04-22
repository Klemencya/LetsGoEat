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

func CreateFriendsTable() error {
	fmt.Println("in CreateFriendsTable...")
	err := checkConnection()
	if err != nil {
		fmt.Println("err in CreateFriendsTable while checkConnection:", err)
		return err
	}

	var query = `
	CREATE TABLE IF NOT EXISTS friends (
		login VARCHAR ( 50 ) UNIQUE NOT NULL,
		friendLogin VARCHAR ( 50 ) UNIQUE NOT NULL
	);`

	_, err = DB.Exec(query)
	if err != nil {
		fmt.Println("err in CreateFriendsTable while Exec(query):", err)
		return err
	}

	fmt.Println("\tfriends table is ready")
	return nil
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

	//schema := fmt.Sprintf(`INSERT INTO user_info (name, surname, email, login, password) VALUES ($1, $2, $3, $4, $5);`,
	//	user.Name, user.Surname, user.Email, user.Login, user.Password)

	fmt.Println("\tuser " + user.Login + " has been inserted into user_info")
	return nil
}
