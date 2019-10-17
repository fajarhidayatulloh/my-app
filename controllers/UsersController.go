package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/my-app/helpers"
	"github.com/my-app/infrastructures"
	"github.com/my-app/models"
	"github.com/my-app/repositories"
	"github.com/my-app/services"
	"github.com/thedevsaddam/govalidator"
)

// InitUsersController init
func InitUsersController() *UsersController {
	usersRepository := new(repositories.UsersRepository)
	usersRepository.DB = &infrastructures.SQLConnection{}

	usersService := new(services.UsersService)
	usersService.UsersRepository = usersRepository

	usersController := new(UsersController)
	usersController.UsersService = usersService

	return usersController
}

// UsersController behaviour
type UsersController struct {
	UsersService services.IUsersService
}

// func isRequestValid(m *models.UserInput) (bool, error) {
// 	validate := validator.New()
// 	err := validate.Struct(m)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

// StoreUser is create new user
func (r *UsersController) StoreUser(res http.ResponseWriter, req *http.Request) {
	var user models.UserInput

	rules := govalidator.MapData{
		"email":    []string{"required", "email"},
		"name":     []string{"required"},
		"password": []string{"required"},
		"phone":    []string{"required"},
	}

	opts := govalidator.Options{
		Request:         req,
		Rules:           rules,
		RequiredDefault: true,
		Data:            &user,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()

	if len(e) != 0 {
		err := map[string]interface{}{"validationError": e}
		helpers.Response(res, http.StatusUnprocessableEntity, err)
		return
	}
	// body, _ := ioutil.ReadAll(req.Body)
	// err := json.Unmarshal(body, &user)

	// if err != nil {
	// 	helpers.Response(res, http.StatusBadRequest, "Failed Input Data Into Database")
	// 	return
	// }

	result, err := r.UsersService.StoreUser(user)
	if err != nil {
		rs := map[string]interface{}{
			"err":      "error_registration",
			"err_desc": fmt.Sprintf("%s", err),
		}
		helpers.Response(res, http.StatusBadRequest, rs)
		return
	}
	resp := map[string]interface{}{
		"data": result,
	}
	helpers.Response(res, http.StatusCreated, resp)
	return

}

//GetUsers is show users list
func (r *UsersController) GetUsers(res http.ResponseWriter, req *http.Request) {

	users, _ := r.UsersService.GetUsers()
	rs := map[string]interface{}{
		"data": users,
	}
	helpers.Response(res, http.StatusOK, rs)
}

// GetUserByID is user by id
func (r *UsersController) GetUserByID(res http.ResponseWriter, req *http.Request) {
	param := mux.Vars(req)
	id, err := strconv.Atoi(param["id"])

	if err != nil {
		helpers.Response(res, http.StatusBadRequest, err)
		return
	}

	user, err := r.UsersService.GetUserByID(id)
	if err != nil {
		helpers.Response(res, http.StatusBadRequest, err)
		return
	}

	if user.ID == 0 {
		// token := jwt.New(jwt.SigningMethodHS256)
		// claims := token.Claims.(jwt.MapClaims)
		// claims["id"] = 23
		// claims["cif_code"] = "28"

		// // Sign and get the complete encoded token as a string using the secret
		// tokenString, err := token.SignedString([]byte(b64.StdEncoding.EncodeToString([]byte("asf987hdghHudfk76lfJJHF08w478234H"))))
		// fmt.Print(tokenString, err)

		if err != nil {
			helpers.Response(res, http.StatusBadRequest, err)
			return
		}

		helpers.Response(res, http.StatusOK, err)
		return
	}
	rs := map[string]interface{}{
		"data": user,
	}

	helpers.Response(res, http.StatusOK, rs)
	return
}
