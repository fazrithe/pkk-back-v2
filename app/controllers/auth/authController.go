package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"pkk-back-v2/app/models"
	u "pkk-back-v2/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/bcrypt"
)

type JwtToken struct {
	AccessToken string `json:"token"`
}

var jwt_secret = os.Getenv("jwt_secret")

func Login(w http.ResponseWriter, req *http.Request) {
	user := &models.User{}

	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	defer req.Body.Close()
	email := user.Email
	password := user.Password

	/* Another way to grab the form inputs from the request
	req.ParseForm()
	username := req.FormValue("Username")
	password := req.FormValue("Password")
	*/
	result := models.GetUsername(email)
	if result == nil {
		u.Respond(w, u.Message(false, "Your credentials do not match our records"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))

	if err != nil {
		fmt.Println(err)
		u.Respond(w, u.Message(false, "Your credentials do not match our records"))
		return
	}
	// access token ttl
	ttl := 2 * time.Minute
	accessTokenExpire := os.Getenv("access_token_expire")
	min, err := strconv.Atoi(accessTokenExpire)
	if err != nil {
		log.Println(err)
	}
	if accessTokenExpire != "" {
		ttl = time.Duration(min) * time.Minute
	}
	CreateToken(w, email, password, ttl)
}

func CreateToken(w http.ResponseWriter, email string, password string, ttl time.Duration) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   ttl,
	})

	tokenString, error := token.SignedString([]byte(jwt_secret))
	if error != nil {
		fmt.Println(error)
	}
	resp := u.Message(true, "success")
	resp["token"] = tokenString
	u.Respond(w, resp)
	return
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte(jwt_secret), nil
				})
				if error != nil {
					u.Respond(w, u.Message(false, error.Error()))
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					u.Respond(w, u.Message(false, "Invalid authorization token"))
					return
				}
			}
		} else {
			u.Respond(w, u.Message(false, "An authorization header is required"))
			return
		}
	})
}
