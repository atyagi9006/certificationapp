package dao

import (
	"context"
	"crypto/tls"
	"log"
	"strconv"

	"github.com/atyagi9006/certificationapp/data-service/src/config"

	driver "github.com/arangodb/go-driver"
	http "github.com/arangodb/go-driver/http"
	"github.com/atyagi9006/certificationapp/data-service/src/models"
)

type daoWrite struct {
	Client driver.Client
}

func NewDBWriter() *daoWrite {

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
	return &daoWrite{Client: c}
}

func (d *daoWrite) CreateCandidate(ctx context.Context, candidate *models.Candidate) bool {
	candidate.Key = candidate.CandidateID
	key := CreateCollection(d, ctx, *candidate, Config.DBConfig.DatabaseName, config.CandidateCollection)
	if key == "" {
		return false
	}
	return true
}
func (d *daoWrite) UpdateCandidate(ctx context.Context, candidateID string, candidate *models.Candidate) bool {
	db, err := d.Client.Database(ctx, Config.DBConfig.DatabaseName)
	if err != nil {
		log.Fatalf("Database Not Found")
	}
	found, err := db.CollectionExists(ctx, config.CandidateCollection)
	if err != nil {
		log.Fatalf("Collection Not Found")
	}
	if found {
		log.Println("Collection Found")
		doc := config.CandidateCollection + "/" + candidateID
		query := `LET doc = DOCUMENT("` + doc + `") UPDATE doc WITH { examAttemptList: PUSH(doc.examAttemptList, @input) } IN  ` + config.CandidateCollection + `  RETURN {id:doc._id,status:true}`
		bindVars := map[string]interface{}{
			"input": candidate.ExamAttemptList[0],
		}
		cursor, err := db.Query(ctx, query, bindVars)
		if err != nil {
			log.Fatalf("document not created %v", err)
		} else {
			log.Println("Document Created")
		}
		res := models.CandidateUpdateRes{}
		_, err = cursor.ReadDocument(ctx, &res)
		log.Println("Operation Updated id: ", res.ID, " sucessful: ", res.Status)
		defer cursor.Close()
		return res.Status
	}
	return false
}

func (d *daoWrite) DeleteCandidate(ctx context.Context, candidateID string) bool { return false }
func (d *daoWrite) CreateUser(ctx context.Context, user *models.User) bool {
	key := CreateCollection(d, ctx, *user, Config.DBConfig.DatabaseName, config.UserCollection)
	if key == "" {
		return false
	}
	user.UserID = key
	return true

}
func (d *daoWrite) UpdateUser(ctx context.Context, user *models.User) bool { return false }
func (d *daoWrite) DeleteUser(ctx context.Context, userName string) bool   { return false }

func CreateCollection(d *daoWrite, ctx context.Context, document interface{}, dbName, collectionName string) string {
	var result string
	db, err := d.Client.Database(ctx, dbName)
	if err != nil {
		log.Fatalf("Database Not Found")
	}

	found, err := db.CollectionExists(ctx, collectionName)
	if err != nil {
		log.Fatalf("Collection Not Found")
	}
	if found {
		log.Println("Collection Found")
		col, _ := db.Collection(ctx, collectionName)
		meta, err := col.CreateDocument(ctx, document)
		if err != nil {
			log.Fatalf("document not created")
		} else {
			log.Println("Document Created")
		}
		result = meta.Key
	}
	return result
}
