package dtest

import (
	"context"
	"database/sql"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// test-user-start
func TestMockFetchUser(t *testing.T) {
	type mock struct {
		db      *sql.DB
		sqlmock sqlmock.Sqlmock
	}
	tests := []struct {
		name   string
		userID string
		mock   mock
		want   *User
		hasErr bool
	}{
		{
			name:   "1件取得",
			userID: "0001",
			// モックの実装
			mock: func() mock {
				db, m, err := sqlmock.New()
				if err != nil {
					t.Fatal(err)
				}
				m.ExpectQuery(regexp.QuoteMeta(`SELECT user_id, user_name FROM users WHERE user_id = $1;`)).
					WithArgs("0001").
					WillReturnRows(sqlmock.NewRows([]string{"user_id", "user_name"}).
						AddRow("0001", "gopher1"),
					)
				return mock{db, m}
			}(),
			want:   &User{UserID: "0001", UserName: "gopher1"},
			hasErr: false,
		},
		{
			name:   "0件取得",
			userID: "9999",
			mock: func() mock {
				db, m, err := sqlmock.New()
				if err != nil {
					t.Fatal(err)
				}
				m.ExpectQuery(regexp.QuoteMeta(`SELECT user_id, user_name FROM users WHERE user_id = $1;`)).
					WithArgs("9999").
					WillReturnError(sql.ErrNoRows)
				return mock{db, m}
			}(),
			want:   nil,
			hasErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db = tt.mock.db
			got, err := FetchUser(context.Background(), tt.userID)
			if (err != nil) != tt.hasErr {
				t.Fatalf("FetchUser() error = %v, hasErr %v", err, tt.hasErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// test-user-end
