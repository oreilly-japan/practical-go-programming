package dtest

import (
	"context"
	"database/sql"
	"io/ioutil"
	"reflect"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// test-user-start
func TestFetchUser(t *testing.T) {
	connStr := "host=localhost port=5432 user=testuser dbname=testdb password=pass sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	sqlBytes, err := ioutil.ReadFile("./schema.sql")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := db.ExecContext(context.TODO(), string(sqlBytes)); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name         string
		userID       string
		inputTestSQL string
		want         *User
		hasErr       bool
	}{
		{
			name:         "1件取得",
			userID:       "0001",
			inputTestSQL: "./testdata/input_user_1.sql",
			want:         &User{UserID: "0001", UserName: "gopher1"},
			hasErr:       false,
		},
		{
			name:         "0件取得",
			userID:       "9999",
			inputTestSQL: "./testdata/input_user_2.sql",
			want:         nil,
			hasErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlBytes, err := ioutil.ReadFile(tt.inputTestSQL)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := db.ExecContext(context.TODO(), string(sqlBytes)); err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() {
				if _, err := db.ExecContext(context.TODO(), `TRUNCATE users;`); err != nil {
					t.Fatal(err)
				}
			})

			got, err := FetchUser(context.TODO(), tt.userID)
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
