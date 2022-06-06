package rest

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"

	"github.com/FeatureOn/api/server/application"

	"github.com/FeatureOn/api/server/adapters/comm/rest/dto"
	"github.com/FeatureOn/api/server/adapters/comm/rest/middleware"
	"github.com/rs/zerolog/log"
)

type ValidatedLogin struct{}

// Claims represents the data for a user login
type Claims struct {
	UserID  string               `json:"userid"`
	Payload jwt.RegisteredClaims `json:"payload"`
}

func (c Claims) Valid() error {
	return c.Payload.Valid()
}

const secretKey = "the_most_secure_secret"
const cookieName = "togglertoken"

var hs = []byte(secretKey)

// Login swagger:route POST PUT /user Login
//
// Handler to log in the user, returns a JWT Token
//
// Responses:
//        200: OK
//		  400: Bad Request
//		  500: Internal Server Error
func (apiContext *APIContext) Login(w http.ResponseWriter, r *http.Request) {
	userLogin := r.Context().Value(ValidatedLogin{}).(dto.LoginRequest)
	userService := application.NewUserService(apiContext.userRepo)
	user, err := userService.CheckUser(userLogin.UserName, userLogin.Password)
	if err != nil {
		respondWithError(w, r, 401, "User not found")
		log.Error().Err(err).Msg("User not found")
		return
	}
	// Create the JWT claims, which includes the username and expiry time
	now := time.Now()
	rclaims := jwt.RegisteredClaims{
		Audience:  jwt.ClaimStrings{"https://toggler.io"},
		ExpiresAt: jwt.NewNumericDate(now.Add(30 * time.Minute)),
		ID:        "featureon",
		IssuedAt:  jwt.NewNumericDate(now),
		Issuer:    "featureon",
		NotBefore: jwt.NewNumericDate(now.Add(30 * time.Minute)),
		Subject:   "featureonlogin",
	}

	pl := Claims{
		Payload: rclaims,
		UserID:  user.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, pl)

	tokenstring, err := token.SignedString(hs)
	if err != nil {
		log.Error().Err(err).Msg("Error creating the token")
		respondWithError(w, r, 500, "Token creation failed")
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:  cookieName,
		Value: tokenstring,
		Path:  "/",
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
	tokenstring := c.Value
	// Initialize a new instance of `Claims`
	claims = &Claims{}

	token, err := jwt.ParseWithClaims(tokenstring, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hs, nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Error validating the token")
		httpStatusCode = http.StatusUnauthorized
		return
	}
	claims = token.Claims.(*Claims)
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
		if claims.Payload.ExpiresAt.Sub(time.Now()) > 30*time.Minute {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Now, create a new token for the current use, with a renewed expiration time
		expirationTime := time.Now().Add(60 * time.Minute)
		claims.Payload.ExpiresAt = jwt.NewNumericDate(expirationTime)
		token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

		tokenstring, err := token.SignedString(hs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Set the new token as the users `token` cookie
		http.SetCookie(w, &http.Cookie{
			Name:    cookieName,
			Value:   tokenstring,
			Expires: expirationTime,
		})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// MiddlewareValidateLoginRequest Checks the integrity of login information in the request and calls next if ok
func (apiContext *APIContext) MiddlewareValidateLoginRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		login, err := middleware.ExtractLoginPayload(r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// validate the login
		errs := apiContext.validation.Validate(login)
		if errs != nil && len(errs) != 0 {
			log.Error().Err(errs[0]).Msg("Error validating the login")

			// return the validation messages as an array
			respondWithJSON(rw, r, http.StatusUnprocessableEntity, errs.Errors())
			return
		}

		// add the rating to the context
		ctx := context.WithValue(r.Context(), ValidatedLogin{}, *login)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
