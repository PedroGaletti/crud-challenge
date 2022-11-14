package accounts

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm/logger"
)

var (
	id              = 1
	document_number = "12345678900"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository IAccountRepository
	controller IAccountController
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	// s.DB.LogMode(true)

	s.repository = NewAccountRepository(s.DB)
	s.controller = NewAccountController(s.repository)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestNewAccountController() {
	controller := NewAccountController(s.repository)
	require.NotNil(s.T(), controller)
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *Suite) TestStore() {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	gin.CreateTestContext(w)
	r := gin.Default()
	r.POST("/accounts", s.controller.Store)

	/* Empty request body bad request case */
	req, _ := http.NewRequest(http.MethodPost, "/accounts", nil)
	r.ServeHTTP(w, req)

	/* Internal server error case */
	internalRequestBody := `{"document_number": "1234567890"}`
	internalBody := strings.NewReader(string(internalRequestBody))
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT").WithArgs("1234567890").WillReturnError(errors.New("error"))
	s.mock.ExpectRollback()
	req, _ = http.NewRequest(http.MethodPost, "/accounts", internalBody)
	r.ServeHTTP(w, req)

	/* Success case */
	successRequestBody := `{"document_number": "12345678900"}`
	successBody := strings.NewReader(string(successRequestBody))
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT").WithArgs(document_number).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	req, _ = http.NewRequest(http.MethodPost, "/accounts", successBody)
	r.ServeHTTP(w, req)

	/* Empty document number request body bad request case */
	emptyDocumentNumber := `{"document_number": ""}`
	emptyBody := strings.NewReader(string(emptyDocumentNumber))
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT").WithArgs("").WillReturnError(errors.New("error"))
	s.mock.ExpectRollback()
	req, _ = http.NewRequest(http.MethodPost, "/accounts", emptyBody)
	r.ServeHTTP(w, req)
}

func (s *Suite) TestShow() {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	gin.CreateTestContext(w)
	r := gin.Default()
	r.GET("/accounts/:id", s.controller.Show)

	/* Empty query param bad request case */
	req, _ := http.NewRequest(http.MethodGet, "/accounts/a", nil)
	r.ServeHTTP(w, req)

	/* No content response case */
	req, _ = http.NewRequest(http.MethodGet, "/accounts/10", nil)
	s.mock.ExpectQuery("SELECT").WillReturnError(logger.ErrRecordNotFound)
	r.ServeHTTP(w, req)

	/* Internal server error case */
	req, _ = http.NewRequest(http.MethodGet, "/accounts/10", nil)
	s.mock.ExpectQuery("SELECT").WillReturnError(errors.New("error"))
	r.ServeHTTP(w, req)

	/* Success case */
	req, _ = http.NewRequest(http.MethodGet, "/accounts/1", nil)
	s.mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "document_number"}).AddRow(id, document_number))
	r.ServeHTTP(w, req)
}

func (s *Suite) TestInjectDependency() {
	InjectDependency(gin.Default().Group(""), s.DB)
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *Suite) TestRouter() {
	controller := NewAccountController(s.repository)
	require.NotNil(s.T(), controller)

	Router(gin.Default().Group(""), controller)
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}
