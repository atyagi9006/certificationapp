package router

import (
	"net/http"

	"github.com/atyagi9006/certificationapp/core-service/src/handler"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func getRoutes() Routes {
	routes := Routes{
		Route{
			"index",
			"GET",
			"/",
			handler.Index,
		},
		Route{
			"SignUp",
			"POST",
			"/signup",
			handler.SignUP,
		},
		Route{
			"SignIn",
			"POST",
			"/signin",
			handler.SignIN,
		},
		Route{
			"getAll users",
			"POST",
			"/users",
			handler.GetAllUser,
		},
		Route{
			"getALLCatagory",
			"GET",
			"/category",
			handler.SignIN,
		},
		/* Route{
			"getCatagory",
			"GET",
			"/category/{id}",
			handler.SignIN,
		},
		Route{
			"getCatagory",
			"POST",
			"/category",
			handler.SignIN,
		},
		Route{
			"getCatagory",
			"PUT",
			"/category",
			handler.SignIN,
		},
		Route{
			"getCatagory",
			"DELETE",
			"/category",
			handler.SignIN,
		}, */
		Route{
			"getQuestions",
			"POST",
			"/testlaunch", //get Questions
			handler.TestLaunch,
		},
		Route{
			"getTest",
			"POST",
			"/answers",
			handler.GetAnswers,
		},
		Route{
			"update Candidate",
			"POST",
			"/updatecandidate",
			handler.UpdateCandidate,
		},
		Route{
			"get Candidate",
			"POST",
			"/candidate",
			handler.GetCandidate,
		},
		Route{
			"createTest",
			"POST",
			"/checkemail",
			handler.CheckEmail,
		},
		/*	Route{
				"updateTest",
				"PUT",
				"/Test/{id}",
				handler.SignIN,
			},
			Route{
				"DeleteTest",
				"DELETE",
				"/Test/{id}",
				handler.SignIN,
			},
			//candidate
			Route{
				"currentteststats",
				"POST",
				"/currentstat/{id}",
				handler.SignIN,
			},
			Route{
				"evaluate",
				"POST",
				"/evaluate/{id}",
				handler.SignIN,
			},
			Route{
				"getResult",
				"POST",
				"/result/{candidateID}",
				handler.SignIN,
			}, */
		Route{
			"CORS",
			"OPTIONS",
			"/{anyurl:[a-z]+}",
			handler.Index,
		},
	}
	return routes
}
