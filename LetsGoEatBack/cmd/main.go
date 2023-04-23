package main

import (
	"awesomeProject4/pkg"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	err := pkg.ConnectDB()
	if err != nil {
		fmt.Println("err in main while ConnectDB:", err)
		return
	}

	err = pkg.CreateUserInfoTable()
	if err != nil {
		fmt.Println("err in main while CreateUserInfoTable:", err)
		return
	}

	err = pkg.CreateRequestsTable()
	if err != nil {
		fmt.Println("err in main while CreateRequestsTable:", err)
		return
	}

	r.Post("/api/login", login)
	r.Post("/api/registration", registration)
	r.Post("/api/get/user", getUser)
	r.Post("/api/send/request", sendRequest)
	r.Post("/api/get/request", getRequest)
	r.Post("/api/delete/request", deleteRequest)
	r.Post("/api/get/allusers", allUsers)
	r.Post("/api/change/request/status", changeAccepted)
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("err in main while ListenAndServe:", err)
		return
	}

	defer func() {
		err := pkg.DisconnectDB()
		if err != nil {
			fmt.Println("err in main while DisconnectDB:", err)
			return
		}
	}()
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in login...")
	var user pkg.UserLogin

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println("err in login while Decode:", err)
		return
	}
	fmt.Println("\t", user)

	err = pkg.LogInCheck(user)
	if err != nil {
		fmt.Println("err in login while LogInCheck:", err)

		response := pkg.MakeResp("Login", err.Error())
		err = sendResp(&response, http.StatusConflict, w)
		if err != nil {
			fmt.Println("err in login while sendResp(Conflict)", err)
		}

		return
	}

	response := pkg.MakeResp("Login", "OK")
	err = sendResp(&response, http.StatusOK, w)
	if err != nil {
		fmt.Println("err in login while sendResp(OK)", err)
		return
	}
}

func sendRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in sendRequest...")
	var inv pkg.Invitation

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&inv)
	if err != nil {
		fmt.Println("err in sendRequest while Decode:", err)
		return
	}
	fmt.Println("\t", inv)

	err = pkg.InsertInvitation(inv)
	if err != nil {
		fmt.Println("err in sendRequest while InsertInvitation:", err)

		response := pkg.MakeResp("SendRequest", err.Error())
		err = sendResp(&response, http.StatusConflict, w)
		if err != nil {
			fmt.Println("err in sendRequest while sendResp(Conflict)", err)
		}

		return
	}

	response := pkg.MakeResp("SendRequest", "OK")
	err = sendResp(&response, http.StatusOK, w)
	if err != nil {
		fmt.Println("err in sendRequest while sendResp(OK)", err)
		return
	}
}

func registration(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in registration...")
	var user pkg.UserRegistration

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println("err in registration while Decode:", err)
		return
	}
	fmt.Println("\t", user)

	err = pkg.InsertUser(user)
	if err != nil {
		fmt.Println("err in registration while InsertUser:", err)

		response := pkg.MakeResp("Registration", err.Error())
		err = sendResp(&response, http.StatusConflict, w)
		if err != nil {
			fmt.Println("err in registration while sendResp(Conflict)", err)
		}

		return
	}

	response := pkg.MakeResp("Registration", "OK")
	err = sendResp(&response, http.StatusOK, w)
	if err != nil {
		fmt.Println("err in registration while sendResp(OK)", err)
		return
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in user...")
	var login pkg.Login

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&login)
	if err != nil {
		fmt.Println("err in getUser while Decode:", err)
		return
	}
	fmt.Println("\t", login)

	userColumn, err := pkg.GetUser(login.Login)
	b, err := json.Marshal(userColumn)
	if err != nil {
		fmt.Println(err)
		return
	}

	response := pkg.MakeUserResp(string(b), "GetUser", "OK")
	err = sendUserResp(&response, http.StatusOK, w)
	if err != nil {
		fmt.Println("err in getUser while sendResp(OK)", err)
		return
	}
}

func sendResp(response *pkg.Response, status int, w http.ResponseWriter) error {
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Println("err in sendResp while Marshal", err)
		return err
	}

	_, err = w.Write(b)
	if err != nil {
		fmt.Println("err in sendResp while Write", err)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return nil
}

func sendUserResp(response *pkg.UserResponse, status int, w http.ResponseWriter) error {
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Println("err in sendResp while Marshal", err)
		return err
	}

	_, err = w.Write(b)
	if err != nil {
		fmt.Println("err in sendResp while Write", err)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return nil
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in getRequest...")
	var login pkg.Login

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&login)
	if err != nil {
		fmt.Println("err in getRequest while Decode:", err)
		return
	}
	fmt.Println("\t", login)

	req, err := pkg.GetRequests(login.Login)
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	response := pkg.MakeUserResp(string(b), "GetRequest", "OK")
	err = sendUserResp(&response, http.StatusOK, w)
	if err != nil {
		fmt.Println("err in getRequest while getRequest(OK)", err)
		return
	}
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in allUsers...")

	req, err := pkg.GetAllUsers()
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	response := pkg.MakeUserResp(string(b), "AllUsers", "OK")
	err = sendUserResp(&response, http.StatusOK, w)
	if err != nil {
		fmt.Println("err in allUsers while getRequest(OK)", err)
		return
	}
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in deleteRequest...")
	var id pkg.ID

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&id)
	if err != nil {
		fmt.Println("err in deleteRequest while Decode:", err)
		return
	}
	fmt.Println("\t", id)

	err = pkg.DeleteRequest(id.ID)
	if err != nil {
		fmt.Println("err in deleteRequest while DeleteRequest:", err)

		response := pkg.MakeResp("DeleteRequest", err.Error())
		err = sendResp(&response, http.StatusConflict, w)
		if err != nil {
			fmt.Println("err in deleteRequest while sendResp(Conflict)", err)
		}

		return
	}

	response := pkg.MakeResp("DeleteRequest", "OK")
	err = sendResp(&response, http.StatusOK, w)
	if err != nil {
		fmt.Println("err in deleteRequest while sendResp(OK)", err)
		return
	}
}

func changeAccepted(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in changeAccepted...")
	var id pkg.ID

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&id)
	if err != nil {
		fmt.Println("err in changeAccepted while Decode:", err)
		return
	}
	fmt.Println("\t", id)

	err = pkg.ChangeAccepted(id.ID)
	if err != nil {
		fmt.Println("err in changeAccepted while DeleteRequest:", err)

		response := pkg.MakeResp("ChangeAccepted", err.Error())
		err = sendResp(&response, http.StatusConflict, w)
		if err != nil {
			fmt.Println("err in changeAccepted while sendResp(Conflict)", err)
		}

		return
	}

	response := pkg.MakeResp("ChangeAccepted", "OK")
	err = sendResp(&response, http.StatusOK, w)
	if err != nil {
		fmt.Println("err in changeAccepted while sendResp(OK)", err)
		return
	}
}
