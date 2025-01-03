package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AshkanShakiba/UserGate/pkg/server"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := server.NewServer(db)

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "John Doe")
	mock.ExpectQuery("select \\* from users where id=1").WillReturnRows(rows)

	req := httptest.NewRequest(http.MethodGet, "/user", bytes.NewReader([]byte(`{"ID": 1}`)))
	w := httptest.NewRecorder()

	s.GetUser(w, req)

	res := w.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)

	var user server.User
	err = json.NewDecoder(res.Body).Decode(&user)
	require.NoError(t, err)
	require.Equal(t, int64(1), user.ID)
	require.Equal(t, "John Doe", user.Name)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := server.NewServer(db)

	mock.ExpectExec(`INSERT INTO users \(name\) VALUES \(\?\)`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader([]byte(`{"Name": "Jane Doe"}`)))
	w := httptest.NewRecorder()

	s.CreateUser(w, req)

	res := w.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)

	require.NoError(t, mock.ExpectationsWereMet())
}
