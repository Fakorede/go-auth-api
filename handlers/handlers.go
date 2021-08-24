package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"goauthapi/models"
	userrespository "goauthapi/repository/userRespository"
	"goauthapi/utils"
	"log"
	"net/http"
)

type Handler struct {}

func (h Handler) Signup(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}

		json.NewDecoder(r.Body).Decode(&user)

		if user.Email == "" {
			utils.ErrorJSON(w, http.StatusBadRequest, errors.New("email field is required"))
			return
		}

		if user.Password == "" {
			utils.ErrorJSON(w, http.StatusBadRequest, errors.New("password field is required"))
			return
		}

		user.Password = utils.HashPassword(user.Password)

		userRepo := userrespository.UserRepository{}
		user, err := userRepo.Signup(db, user)
		if err != nil {
			utils.ErrorJSON(w, http.StatusInternalServerError, errors.New("something went wrong"))
			return
		}
		
		user.Password = ""

		utils.WriteJSON(w, http.StatusOK, user, "data")

	}
}

func (h Handler) Login(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}

		json.NewDecoder(r.Body).Decode(&user)

		if user.Email == "" {
			utils.ErrorJSON(w, http.StatusBadRequest, errors.New("email field is required"))
			return
		}

		if user.Password == "" {
			utils.ErrorJSON(w, http.StatusBadRequest, errors.New("password field is required"))
			return
		}
		
		password := user.Password

		userRepo := userrespository.UserRepository{}
		user, err := userRepo.Login(db, user)
		if err != nil {
			if err == sql.ErrNoRows {
				utils.ErrorJSON(w, http.StatusBadRequest, errors.New("incorrect credentials"))
				return
			} else {
				log.Fatal(err)
			}
		}

		hashedPassword := user.Password

		err = utils.ComparePasswords(hashedPassword, password)
		if err != nil {
			utils.ErrorJSON(w, http.StatusUnauthorized, errors.New("incorrect credentials"))
			return
		}

		token, err := utils.GenerateToken(user)
		if err != nil {
			log.Fatal(err)
		}

		utils.WriteJSON(w, http.StatusOK, token, "token")

	}
}

func (h Handler) Protected(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("protected endpoint"))
	}
}
