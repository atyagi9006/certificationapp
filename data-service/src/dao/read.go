package dao

import (
	"context"
	"crypto/tls"
	"log"
	"strconv"

	driver "github.com/arangodb/go-driver"
	http "github.com/arangodb/go-driver/http"
	"github.com/atyagi9006/certificationapp/data-service/src/config"
	"github.com/atyagi9006/certificationapp/data-service/src/models"
)

type daoRead struct {
	Client driver.Client
}

func NewDBReader() *daoRead {

	url := "http://" + Config.DBConfig.URL + ":" + strconv.Itoa(int(Config.DBConfig.ArangoPort))
	log.Println(url)
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{url},
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	})
	if err != nil {
		log.Fatalf("problem occurred getting connection to db")
	}
	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(Config.DBConfig.UserName, Config.DBConfig.Password),
	})
	if err != nil {
		log.Fatalf("problem occurred getting connection to db")
	}
	return &daoRead{Client: c}
}

func (read *daoRead) GetAdmin(ctx context.Context, adminID string) *models.Admin {
	//query:="FOR d IN Admin FILTER d.Name == @name RETURN d"
	//read.Client.

	return nil
}
func (read *daoRead) GetCandidate(ctx context.Context, candidateID string) *models.Candidate {
	db, err := read.Client.Database(ctx, Config.DBConfig.DatabaseName)

	candidate := models.Candidate{
		CandidateID: candidateID,
	}
	found, err := db.CollectionExists(ctx, config.CandidateCollection)
	if err != nil {
		log.Fatalf("Collection Not Found")
	}
	if found {
		log.Println("Collection Found")
		col, _ := db.Collection(ctx, config.CandidateCollection)
		_, err := col.ReadDocument(ctx, candidateID, &candidate)
		if err != nil {
			log.Fatalln("NO Candidate found error : ", err)
		} else {
			log.Println("Candidate found ")
		}

	}

	return &candidate
}
func (read *daoRead) GetAllCandidates(ctx context.Context) []*models.Candidate { return nil }

func (read *daoRead) GetUser(ctx context.Context, user models.User) *models.User {
	db, err := read.Client.Database(ctx, Config.DBConfig.DatabaseName)

	query := "FOR doc IN user FILTER  doc.email== @email RETURN doc"
	//doc.userName == @usrname
	bindVars := map[string]interface{}{
		//	"usrname": user.UserName,
		"email": user.Email,
	}
	cursor, err := db.Query(ctx, query, bindVars)
	if err != nil {
		// handle error
	}
	resUser := models.User{}
	meta, err := cursor.ReadDocument(ctx, &resUser)
	resUser.UserID = meta.Key
	log.Println("read user : ", resUser.Email, "Key: ", resUser.UserID)
	defer cursor.Close()
	return &resUser
}
func (read *daoRead) GetAllUsers(ctx context.Context, requestedBy models.User) []*models.User {
	var result []*models.User
	db, err := read.Client.Database(ctx, Config.DBConfig.DatabaseName)
	query := "FOR doc IN user FILTER doc._key != @id RETURN doc"
	bindVars := map[string]interface{}{
		"id": requestedBy.UserID,
	}
	cursor, err := db.Query(ctx, query, bindVars)
	if err != nil {
		// handle error
	}
	defer cursor.Close()
	for {
		var user models.User
		meta, err := cursor.ReadDocument(ctx, &user)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			// handle other errors
		}
		user.UserID = meta.Key
		result = append(result, &user)
		log.Printf("Got doc with key '%s' from query\n", meta.Key)
	}

	return result

}
func (read *daoRead) TestLaunch(ctx context.Context, Test models.Test) []*models.Question {

	db, err := read.Client.Database(ctx, Config.DBConfig.DatabaseName)

	//query := "FOR doc IN question FILTER doc.category == @category LIMIT 10 RETURN doc"
	query := "FOR doc IN question FILTER doc.category == @category  SORT RAND() LIMIT 10 RETURN doc"
	bindVars := map[string]interface{}{
		"category": Test.CategoryID,
	}
	cursor, err := db.Query(ctx, query, bindVars)
	if err != nil {
		// handle error
	}
	defer cursor.Close()
	var result []*models.Question
	for {
		var que models.Question
		meta, err := cursor.ReadDocument(ctx, &que)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			// handle other errors
		}
		que.QuestionID = meta.Key
		result = append(result, &que)
		log.Printf("Got doc with key '%s' from query\n", meta.Key)
	}
	return result
}
