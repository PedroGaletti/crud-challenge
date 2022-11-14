package accounts

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository IAccountRepository
	controller IAccountController
	Account    *Account
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

	s.DB.LogMode(true)

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

	requestBody := `{"document_number": 12345678900}`
	body := strings.NewReader(string(requestBody))

	r.POST("/accounts", s.controller.Store)

	// Test with empty body
	req, err := http.NewRequest(http.MethodPost, "/accounts", nil)
	if err != nil {
		s.T().Fatal(fmt.Printf("Couldn't create request: %v\n", err))
	}

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		require.Error(s.T(), errors.New("Request with error"))
	}

	// Test with a valid array
	req, err = http.NewRequest(http.MethodPost, "/accounts", body)
	if err != nil {
		s.T().Fatal(fmt.Printf("Couldn't create request: %v\n", err))
	}

	r.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		require.NoError(s.T(), err)
	}
}
