package transactions

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
)

var (
	account_id           = 1
	operation_id         = 1
	payment_operation_id = 4
	amount               = 5
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository ITransactionRepository
	controller ITransactionController
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

	s.repository = NewTransactionRepository(s.DB)
	s.controller = NewTransactionController(s.repository)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestNewTransactionController() {
	controller := NewTransactionController(s.repository)
	require.NotNil(s.T(), controller)
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *Suite) TestStore() {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	gin.CreateTestContext(w)
	r := gin.Default()
	r.POST("/transactions", s.controller.Store)

	/* Empty request body bad request case */
	req, _ := http.NewRequest(http.MethodPost, "/transactions", nil)
	r.ServeHTTP(w, req)

	/* Internal server error case */
	internalRequestBody := `{"account_id": 1, "operation_id": 1, "amount": 5}`
	internalBody := strings.NewReader(string(internalRequestBody))
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT").WithArgs(account_id, operation_id, -(amount), sqlmock.AnyArg()).WillReturnError(errors.New("error"))
	s.mock.ExpectRollback()
	req, _ = http.NewRequest(http.MethodPost, "/transactions", internalBody)
	r.ServeHTTP(w, req)

	/* Success case */
	successRequestBody := `{"account_id": 1, "operation_id": 4, "amount": 5}`
	successBody := strings.NewReader(string(successRequestBody))
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT").WithArgs(account_id, payment_operation_id, amount, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	req, _ = http.NewRequest(http.MethodPost, "/transactions", successBody)
	r.ServeHTTP(w, req)

	/* zero amount request body bad request case */
	zeroAmount := `{"account_id": 1, "operation_id": 1, "amount": 0}`
	zBody := strings.NewReader(string(zeroAmount))
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT").WithArgs(account_id, operation_id, amount).WillReturnError(errors.New("error"))
	s.mock.ExpectRollback()
	req, _ = http.NewRequest(http.MethodPost, "/transactions", zBody)
	r.ServeHTTP(w, req)
}

func (s *Suite) TestInjectDependency() {
	InjectDependency(gin.Default().Group(""), s.DB)
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *Suite) TestRouter() {
	controller := NewTransactionController(s.repository)
	require.NotNil(s.T(), controller)

	Router(gin.Default().Group(""), controller)
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}
