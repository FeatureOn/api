package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/application"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/adapters/comm/rest/dto"
	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Claims represents the data for a user login
type Claims struct {
	UserID   string `json:"userid"`
	TenantID string `json:"tenantid"`
	jwt.StandardClaims
}

const secretKey = "the_most_secure_secret"
const cookieName = "togglertoken"

// Login swagger:route POST PUT /user Login
//
// Handler to login the user, returns a JWT Token
//
// Responses:
//        200: OK
//		  400: Bad Request
//		  500: Internal Server Error
func (apiContext *APIContext) Login(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Error().Err(err)
		return
	}
	if len(payload) == 0 {
		respondWithError(w, r, 400, viper.GetString("PayloadMissingMsg"))
		log.Error()
		return
	}
	var userLogin dto.LoginRequest
	err = json.Unmarshal(payload, &userLogin)
	if err != nil {
		respondWithError(w, r, 400, viper.GetString("CannotReadPayloadMsg"))
		log.Error().Err(err).Msg(viper.GetString("CannotReadPayloadMsg"))
		return
	}
	userService := application.NewUserService(apiContext.userRepo)
	user, err := userService.CheckUser(userLogin.UserName, userLogin.Password)
	if err != nil {
		respondWithError(w, r, 400, "User not found")
		log.Error().Err(err).Msg("User not found")
		return
	}
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(6000 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	var jwtKey = []byte(secretKey)
	// Create the JWT key used to create the signature
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:  cookieName,
		Value: tokenString,
		//Expires: expirationTime,
		Path: "/",
		//Domain:  "cookieseal.com",
	})
}

func checkLogin(r *http.Request) (status bool, httpStatusCode int, claims *Claims) {
	// We can obtain the session token from the requests cookies, which come with every request
	// Initialize a new instance of `Claims`
	c, err := r.Cookie(cookieName)
	status = false
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			httpStatusCode = http.StatusUnauthorized
			return
		}
		// For any other type of error, return a bad request status
		httpStatusCode = http.StatusBadRequest
		return
	}

	// Get the JWT string from the cookie
	tknStr := c.Value
	var jwtKey = []byte(secretKey)
	// Initialize a new instance of `Claims`
	claims = &Claims{}
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {

		return jwtKey, nil
	})
	if !tkn.Valid {
		httpStatusCode = http.StatusUnauthorized
		return
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			httpStatusCode = http.StatusUnauthorized
			return
		}
		httpStatusCode = http.StatusBadRequest
		return
	}
	status = true
	return
}

// Refresh swagger:route POST PUT /login Refresh
//
// Handler to refresh a JWT Token
//
// Responses:
//        200: OK
//		  400: Bad Request
//		  500: Internal Server Error
func (apiContext *APIContext) Refresh(w http.ResponseWriter, r *http.Request) {
	status, _, claims := checkLogin(r)
	if status {
		// We ensure that a new token is not issued until enough time has elapsed
		// In this case, a new token will only be issued if the old token is within
		// 30 seconds of expiry. Otherwise, return a bad request status
		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Now, create a new token for the current use, with a renewed expiration time
		expirationTime := time.Now().Add(60 * time.Minute)
		claims.ExpiresAt = expirationTime.Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		var jwtKey = []byte(secretKey)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Set the new token as the users `token` cookie
		http.SetCookie(w, &http.Cookie{
			Name:    cookieName,
			Value:   tokenString,
			Expires: expirationTime,
		})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
