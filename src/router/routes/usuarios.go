package router

import (
	"book-api/src/controllers"
	"net/http"
)

var userRouter = []Routes{
	{
		URI:          "/usuarios",
		Methods:      http.MethodPost,
		Function:     controllers.CreateUser,
		RequiredAuth: false,
	},
	{
		URI:          "/usuarios",
		Methods:      http.MethodGet,
		Function:     controllers.ListAllUser,
		RequiredAuth: false,
	},
	{
		URI:          "/usuarios/:id",
		Methods:      http.MethodGet,
		Function:     controllers.ListUserById,
		RequiredAuth: false,
	},
	{
		URI:          "/usuarios/:id",
		Methods:      http.MethodDelete,
		Function:     controllers.DeleteUser,
		RequiredAuth: false,
	},
	{
		URI:          "/usuarios/:id",
		Methods:      http.MethodPut,
		Function:     controllers.UpdateUser,
		RequiredAuth: false,
	},
}
