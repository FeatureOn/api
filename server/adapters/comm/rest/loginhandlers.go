package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/FeatureOn/api/application"

	"github.com/FeatureOn/api/adapters/comm/rest/dto"
	middleware "github.com/FeatureOn/api/adapters/comm/rest/middleware"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/rs/zerolog/log"
)

type ValidatedLogin struct{}

// Claims represents the data for a user login
type Claims struct {
	UserID  string      `json:"userid"`
	Payload jwt.Payload `json:payload`
}

const secretKey = "the_most_secure_secret"
const cookieName = "togglertoken"

var hs = jwt.NewHS256([]byte(secretKey))

// Login swagger:route POST PUT /user Login
//
// Handler to login the user, returns a JWT Token
//
// Responses:
//        200: OK
//		  400: Bad Request
//		  500: Internal Server Error
func (ctx *APIContext) Login(w http.ResponseWriter, r *http.Request) {
	userLogin := r.Context().Value(ValidatedLogin{}).(dto.LoginRequest)
	userService := application.NewUserService(ctx.userRepo)
	user, err := userService.CheckUser(userLogin.UserName, userLogin.Password)
	if err != nil {
		respondWithError(w, r, 401, "User not found")
		log.Error().Err(err).Msg("User not found")
		return
	}
	// Create the JWT claims, which includes the username and expiry time
	now := time.Now()
	pl := Claims{
		Payload: jwt.Payload{
			Issuer:         "toggler",
			Subject:        "togglerlogin",
			Audience:       jwt.Audience{"https://toggler.io", "https://jwt.io"},
			ExpirationTime: jwt.NumericDate(now.Add(30 * time.Minute)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          "toggler",
		},
		UserID: user.ID,
	}

	token, err := jwt.Sign(pl, hs)
	if err != nil {
		log.Error().Err(err).Msg("Error creating the token")
		respondWithError(w, r, 500, "Token creation failed")
		return
	}

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tokenstring := string(token[:])
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:  cookieName,
		Value: tokenstring,
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
	tokenstring := c.Value
	// Initialize a new instance of `Claims`
	claims = &Claims{}
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	_, err = jwt.Verify([]byte(tokenstring), hs, claims)

	if err != nil {
		httpStatusCode = http.StatusUnauthorized
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
func (ctx *APIContext) Refresh(w http.ResponseWriter, r *http.Request) {
	status, _, claims := checkLogin(r)
	if status {
		// We ensure that a new token is not issued until enough time has elapsed
		// In this case, a new token will only be issued if the old token is within
		// 30 seconds of expiry. Otherwise, return a bad request status
		if claims.Payload.ExpirationTime.Sub(time.Now()) > 30*time.Minute {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Now, create a new token for the current use, with a renewed expiration time
		expirationTime := time.Now().Add(60 * time.Minute)
		claims.Payload.ExpirationTime = jwt.NumericDate(expirationTime)
		token, err := jwt.Sign(claims, hs)
		if err != nil {
			log.Error().Err(err).Msg("Error creating the token")
			respondWithError(w, r, 500, "Token creation failed")
			return
		}

		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tokenstring := string(token[:])
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
func (ctx *APIContext) MiddlewareValidateLoginRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		login, err := middleware.ExtractLoginPayload(r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// validate the login
		errs := ctx.validation.Validate(login)
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
