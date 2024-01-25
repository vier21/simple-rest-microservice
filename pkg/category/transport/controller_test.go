package transport

import (
	"bytes"
	"encoding/json"
	"fmt"

	"gorest/internal/utils"
	"gorest/pkg/category/repository"
	"gorest/pkg/category/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

type WebRequest struct {
	Id     string    `json:"id"`
	Name   string    `json:"name"`
	Tstamp time.Time `json:"tstamp"`
}

func newDBTest() (*sqlx.DB, error) {
	dsn := "root:root123@tcp(127.0.0.1:2109)/category_test?parseTime=true"
	DB, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return nil, err
	}

	DB.SetMaxOpenConns(50)
	DB.SetConnMaxIdleTime(50 * time.Second)

	if err := DB.Ping(); err != nil {
		fmt.Printf("error ping database: %s", err)
		return nil, err
	}

	return DB, nil
}

func newRouter(db *sqlx.DB) http.Handler {
	db, err := newDBTest()
	repo := repository.NewDataStore(db)
	svc := service.NewService(repo, validator.New())

	mux := mux.NewRouter()

	InitHttpHandler(mux, svc)

	if err != nil {
		return nil
	}

	return mux
}

func insertCtg(t *testing.T, db *sqlx.DB) map[string]interface{} {
	router := newRouter(db)

	ctg := WebRequest{
		Name: utils.RandomString(5),
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(ctg)

	request, err := http.NewRequest(http.MethodPost, "https://localhost:3030/category/", &buf)
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fail()
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	resp := recorder.Result()
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)
	return responseBody
}

func truncateTable(db *sqlx.DB) {
	db.Exec("TRUNCATE category")
}
func TestSaveCtgHandlerSuccess(t *testing.T) {
	db, err := newDBTest()
	router := newRouter(db)

	if err != nil {
		t.Fail()
	}
	ctg := WebRequest{
		Name: utils.RandomString(5),
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(ctg)

	request, err := http.NewRequest(http.MethodPost, "https://localhost:3030/category/", &buf)
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fail()
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	resp := recorder.Result()
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	fmt.Println(responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, ctg.Name, responseBody["data"].(map[string]interface{})["name"])

}

func TestSaveCtgHandlerFail(t *testing.T) {
	db, err := newDBTest()

	router := newRouter(db)
	if err != nil {
		t.Fail()
	}

	buf := strings.NewReader("{}")

	request, err := http.NewRequest(http.MethodPost, "https://localhost:3030/category/", buf)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Content-Type", "application/json")

	if err != nil {
		t.Fail()
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	resp := recorder.Result()
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestUpdateCtgHandler(t *testing.T) {
	db, err := newDBTest()
	truncateTable(db)
	router := newRouter(db)

	data := insertCtg(t, db)
	id := data["data"].(map[string]interface{})["id"].(string)

	if err != nil {
		t.Fail()
	}
	ctg := WebRequest{
		Name: utils.RandomString(5),
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(ctg)

	request, err := http.NewRequest(http.MethodPut, "https://localhost:3030/category/"+id, &buf)
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fail()
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	resp := recorder.Result()
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, ctg.Name, responseBody["data"].(map[string]interface{})["name"])

}

func TestUpdateCtgHandlerFail(t *testing.T) {
	db, err := newDBTest()
	truncateTable(db)
	router := newRouter(db)

	data := insertCtg(t, db)
	id := data["data"].(map[string]interface{})["id"].(string)

	if err != nil {
		t.Fail()
	}

	buf := strings.NewReader("{}")
	request, err := http.NewRequest(http.MethodPut, "https://localhost:3030/category/"+id, buf)
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fail()
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	resp := recorder.Result()
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))

}
func TestDeleteCtgHandlerSuccess(t *testing.T) {
	db, err := newDBTest()
	router := newRouter(db)
	data := insertCtg(t, db)

	id := data["data"].(map[string]interface{})["id"].(string)
	if err != nil {
		t.Fail()
	}

	request, err := http.NewRequest(http.MethodDelete, "https://localhost:3030/category/"+id, nil)
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fail()
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	resp := recorder.Result()
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestFindAllCtgHandlerSuccess(t *testing.T) {
	db, err := newDBTest()
	router := newRouter(db)
	_ = insertCtg(t, db)
	if err != nil {
		t.Fail()
	}

	request, err := http.NewRequest(http.MethodGet, "https://localhost:3030/category/", nil)
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fail()
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	resp := recorder.Result()
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody["data"].([]interface{})[0])

	assert.Equal(t, 200, resp.StatusCode)
}

func TestFindByIdaveCtgHandlerSuccess(t *testing.T) {
	db, err := newDBTest()
	router := newRouter(db)
	data := insertCtg(t, db)

	id := data["data"].(map[string]interface{})["id"].(string)
	if err != nil {
		t.Fail()
	}

	request, err := http.NewRequest(http.MethodGet, "https://localhost:3030/category/"+id, nil)
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fail()
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	resp := recorder.Result()
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody["data"])

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, id, responseBody["data"].(map[string]interface{})["id"])
}
func TestFindByIdaveCtgHandlerFail(t *testing.T) {
	db, err := newDBTest()
	router := newRouter(db)
	_ = insertCtg(t, db)

	if err != nil {
		t.Fail()
	}

	request, err := http.NewRequest(http.MethodGet, "https://localhost:3030/category/"+"1", nil)
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fail()
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	resp := recorder.Result()
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody["data"])

	assert.Equal(t, 400, resp.StatusCode)
}
