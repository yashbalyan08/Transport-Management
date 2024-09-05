package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"

	models "github.com/yashbalyan08/system/models"
)

var store = sessions.NewCookieStore([]byte("secure-key"))

// CheckPasswordHash compares hashed password with the entered one
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// LoginHandler processes login requests
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// fmt.Println("User: %v", user)

	// Query database for the user's hashed password
	var hashedPassword string
	err = db.QueryRow("SELECT password FROM users WHERE id = ?", user.Id).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Verify password
	if !CheckPasswordHash(user.Password, hashedPassword) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// If user successfully logs in, create a session
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = true
	session.Values["username"] = user.Username
	session.Save(r, w)

	// Successful login
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful!"})
}

// LogoutHandler logs out the user
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out!"})
}
