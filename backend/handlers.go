package main;

import(
	"encoding/json"
	"fmt"
	"net/http"

)

type User struct{
	Email    string `json:"email"`
	Password string `json:"password"`
}

// in-memory storage for users (email -> password)
var users = make(map[string]string)

//Sign Up Handling

func signupHandler(w http.ResponseWriter, r *http.Request){
	if r.Method	 == "POST" {
		var newUser User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil{
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			return

		}
		//storing  user information
		users[newUser.Email]= newUser.Password
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w,"User Signed up Successfully")
	}else{
		http.Error(w, "Invalid Request method", http.StatusMethodNotAllowed)

	}
}

// login handling

func loginHandler( w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		var loginUser User
		err:= json.NewDecoder(r.Body).Decode(&loginUser)
		if err !=nil{
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			return

		}

		//checking wether user exist or not

		storedPassword , exists := users[loginUser.Email]
		if !exists || storedPassword != loginUser.Password{
			http.Error(w, "Invalid  email or password", http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Login Successfully")

	}else{
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
