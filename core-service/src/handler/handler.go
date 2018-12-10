package handler

import (
	"log"

	"github.com/atyagi9006/certificationapp/core-service/src/config"

	//"github.com/atyagi9006/certificationapp/grpcproto"
	"encoding/json"
	"io"
	"io/ioutil"

	"net/http"
	"strings"

	"github.com/atyagi9006/certificationapp/core-service/src/helper"
	"github.com/atyagi9006/certificationapp/grpcproto"

	"github.com/atyagi9006/certificationapp/core-service/src/util"

	"github.com/atyagi9006/certificationapp/core-service/src/models"
)

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	setupResponse(&w, r)
	w.Write([]byte("{\"result\":\"OK\"}"))
}
func SignUP(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	var inputUser models.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &inputUser); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	if strings.Trim(inputUser.Type, " ") == "" {
		inputUser.Type = config.UserTYPE
	}
	conn := util.GetGRPCConn()

	client := grpcproto.NewUserServiceClient(conn)
	ures, err := helper.CreateUser(client, &inputUser)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	if ures.Status {
		candClient := grpcproto.NewCandidateServiceClient(conn)
		newCandidate := &models.Candidtate{
			CandidtateID: "candidate_" + ures.UserID,
		}
		cres, err := helper.CreateCandidate(candClient, newCandidate)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"message\":\"user creation failed \"}"))
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(cres); err != nil {
				panic(err)
			}
		}
	}
}

func SignIN(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	var inputUser models.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &inputUser); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	conn := util.GetGRPCConn()
	client := grpcproto.NewUserServiceClient(conn)
	ures, err := helper.GetUser(client, &inputUser)

	if strings.EqualFold(util.GetMD5Hash(inputUser.Password), ures.Password) {
		responsePayload := models.User{
			Name:     ures.Name,
			Email:    ures.Email,
			UserID:   ures.UserID,
			Type:     ures.Type,
			UserName: ures.UserName,
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(responsePayload); err != nil {
			panic(err)
		}

	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("{\"Error\":\"Wrong E-mail or Password\"}"))

	}
}
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	var inputUser models.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &inputUser); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	conn := util.GetGRPCConn()

	client := grpcproto.NewUserServiceClient(conn)
	ures, err := helper.GetAllUsers(client, &inputUser)
	var responsePayload []models.User
	for _, rusr := range ures.Users {
		responsePayload = append(responsePayload, models.User{
			Name:     rusr.Name,
			Email:    rusr.Email,
			UserID:   rusr.UserID,
			Type:     rusr.Type,
			UserName: rusr.UserName,
		})

	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(responsePayload); err != nil {
		panic(err)
	}
}

func TestLaunch(w http.ResponseWriter, r *http.Request) {
	funcName := "TestLaunch"
	log.Println("Initiating Test - " + funcName)
	setupResponse(&w, r)
	var inputTest models.Test
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &inputTest); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	conn := util.GetGRPCConn()

	client := grpcproto.NewTestServiceClient(conn)
	responsePayload, err := helper.GetQuestionsAndProcess(client, &inputTest)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(responsePayload); err != nil {
		panic(err)
	}
}
func GetAnswers(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	var inputQues []string
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &inputQues); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	responsePayload, err := helper.GetAnswers(inputQues)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(responsePayload); err != nil {
		panic(err)
	}
}

func UpdateCandidate(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	var inputCandidate models.Candidtate
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &inputCandidate); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	conn := util.GetGRPCConn()
	candClient := grpcproto.NewCandidateServiceClient(conn)
	cres, err := helper.UpdateCandidate(candClient, &inputCandidate)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(cres); err != nil {
			panic(err)
		}
	}

}
func GetCandidate(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	var candidate models.Candidtate
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &candidate); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	conn := util.GetGRPCConn()
	client := grpcproto.NewCandidateServiceClient(conn)
	responsePayload, err := helper.GetCandidate(client, &candidate)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(responsePayload); err != nil {
		panic(err)
	}
}
func CheckEmail(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	var inputUser models.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &inputUser); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	if strings.Trim(inputUser.Type, " ") == "" {
		inputUser.Type = config.UserTYPE
	}
	conn := util.GetGRPCConn()

	client := grpcproto.NewUserServiceClient(conn)
	ures, err := helper.GetUser(client, &inputUser)
	if err != nil {
		log.Fatal(err)
	}

	responsePayload := make(map[string]string)
	if ures.Email != "" {
		responsePayload["valid"] = "false"
	} else {
		responsePayload["valid"] = "true"
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(responsePayload); err != nil {
		log.Fatal(err)
	}
}
