package johndb

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/lib/env"
	"github.com/rlawnsxo131/madre-server-v3/lib/logger"
	"github.com/rs/zerolog"
)

var (
	instanceDatabase *singletonDatabase
	onceDatabase     sync.Once
)

type Database interface {
	Queryx(query string, args ...any) (*sqlx.Rows, error)
	QueryRowx(query string, args ...any) *sqlx.Row
	NamedQuery(query string, arg any) (*sqlx.Rows, error)
}

type singletonDatabase struct {
	DB *sqlx.DB
	l  *zerolog.Logger
}

func DatabaseInstance() (*singletonDatabase, error) {
	var connectError error

	onceDatabase.Do(func() {
		psqlInfo := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			env.DatabaseHost(),
			env.DatabasePort(),
			env.DatabaseUser(),
			env.DatabasePassword(),
			env.DatabaseDBName(),
			env.DatabaseSSLMode(),
		)
		logger.DefaultLogger().Info().
			Timestamp().Str("database connection info", psqlInfo).Send()

		db, err := sqlx.Connect("postgres", psqlInfo)
		if err != nil {
			connectError = errors.Wrap(err, "sqlx connect fail")
		}

		l := zerolog.New(os.Stderr).With().Logger()
		instanceDatabase = &singletonDatabase{db, &l}
	})

	return instanceDatabase, connectError
}

func (sd *singletonDatabase) WithTimeoutTxx() (*sqlx.Tx, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	tx, err := sd.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "database instance WithTimeoutTxx BeginTxx error")
	}

	return tx, nil
}

// func(sd *singletonDatabase) Queryx(query string, args ...any) (*sqlx.Rows, error) {

// }

// ####### Test ######
// ctx := context.Background()
// tx, err := db.DB.BeginTxx(ctx, nil)
// if err != nil {
// 	log.Println("tx error")
// }
// defer tx.Rollback()
// user := account.User{
// 	Email:      "asdf",
// 	OriginName: utils.NewNullString("asdf"),
// 	Username:   "asdf",
// 	PhotoUrl:   utils.NewNullString("asdf"),
// }
// u, err := tx.NamedExec("INSERT INTO public.user(email, origin_name, username, photo_url) VALUES(:email, :origin_name, :username, :photo_url) RETURNING id", &user)
// if err != nil {
// 	log.Println("user error", err)
// }
// log.Println(u)

// socialAccount := account.SocialAccount{
// 	SocialID: "asdf",
// 	Provider: "GOOGLE",
// }
// sa, err := tx.NamedExec("INSERT INTO social_account(user_id, provider, social_id) VALUES(:user_id, :provider, :social_id) RETURNING id", &socialAccount)
// if err != nil {
// 	log.Println("socialaccount error", err)
// }
// log.Println(sa)
// err = tx.Commit()

// if err != nil {
// 	log.Println("commit error", err)
// }
