package dao

import (
	"context"
	"database/sql"
	"homework/internal/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var ErrRecordNotFound = errors.New("Not Found")


var Provider = wire.NewSet(NewDB, NewDao)

type Dao interface {
	GetUser(ctx context.Context, id int) (*model.User, error)
}

type dao struct {
	db *sql.DB
}

func (d *dao) GetUser(ctx context.Context, id int) (*model.User, error) {
	user := &model.User{}
	row := d.db.QueryRowContext(ctx, "select id,name from users where id=?", id)
	err := row.Scan(&user.Id, &user.Name)
	if err == sql.ErrNoRows {
		return nil, errors.Wrap(ErrRecordNotFound, "No exist the user")
	}
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get user")
	}
	return user, nil
}

func NewDao(db *sql.DB) Dao {
	return &dao{db: db}
}

func NewDB() (db *sql.DB, cleanup func(), err error) {
	db, err = sql.Open("mysql", viper.GetString("mysql.dsn"))
	cleanup = func() {
		if err == nil {
			db.Close()
		}
	}
	return
}

