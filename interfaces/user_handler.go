package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"
	"order-service/application"
	domain "order-service/domain"
	"strconv"

	"github.com/gorilla/mux"
	//"github.com/gin-gonic/gin"
)

//Users struct defines the dependencies that will be used
type UserHandler struct {
	us application.UserAppInterface
}

//Users constructor
func NewUserHandler(us application.UserAppInterface) UserHandler {
	return UserHandler{
		us: us,
	}
}

// ErrorResponse is Error response template
type ErrorResponse struct {
	Message string `json:"reason"`
	Error   error  `json:"-"`
}

func (e *ErrorResponse) String() string {
	return fmt.Sprintf("reason: %s, error: %s", e.Message, e.Error.Error())
}

// Respond is response write to ResponseWriter
func Respond(w http.ResponseWriter, code int, src interface{}) {
	var body []byte
	var err error

	switch s := src.(type) {
	case []byte:
		if !json.Valid(s) {
			Error(w, http.StatusInternalServerError, err, "invalid json")
			return
		}
		body = s
	case string:
		body = []byte(s)
	case *ErrorResponse, ErrorResponse:
		// avoid infinite loop
		if body, err = json.Marshal(src); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"reason\":\"failed to parse json\"}"))
			return
		}
	default:
		if body, err = json.Marshal(src); err != nil {
			Error(w, http.StatusInternalServerError, err, "failed to parse json")
			return
		}
	}
	w.WriteHeader(code)
	w.Write(body)
}

// Error is wrapped Respond when error response
func Error(w http.ResponseWriter, code int, err error, msg string) {
	e := &ErrorResponse{
		Message: msg,
		Error:   err,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Respond(w, code, e)
}

// JSON is wrapped Respond when success response
func JsonResponse(w http.ResponseWriter, code int, src interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Respond(w, code, src)
}

func (s *UserHandler) SaveUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		JsonResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	//validate the request:
	validateErr := user.Validate("")
	if len(validateErr) > 0 {
		JsonResponse(w, http.StatusUnprocessableEntity, validateErr)
		return
	}
	newUser, err := s.us.SaveUser(&user)

	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	JsonResponse(w, http.StatusCreated, newUser.PublicUser())
}

func (s *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := domain.Users{} //customize user
	var err error
	users, err = s.us.GetUsers()
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	JsonResponse(w, http.StatusOK, users.PublicUsers())
}

func (s *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	//userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	vars := mux.Vars(r)
	userId, err := strconv.ParseUint(vars["userId"], 10, 64)

	if err != nil {
		JsonResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	user, err := s.us.GetUser(userId)
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	JsonResponse(w, http.StatusOK, user.PublicUser())
}
