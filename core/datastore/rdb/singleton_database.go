package rdb

// func (sd *singletonDatabase) QueryRowx(query string, args ...any) *sqlx.Row {
// 	sd.l.Log().Timestamp().Str("query", fmt.Sprintf("%s,%+v", query, args)).Send()
// 	return sd.DB.QueryRowx(query, args...)
// }

// func (sd *singletonDatabase) NamedQuery(query string, arg any) (*sqlx.Rows, error) {
// 	sd.l.Log().Timestamp().Str("query", fmt.Sprintf("%s,%+v", query, arg)).Send()
// 	return sd.DB.NamedQuery(query, arg)
// }

// func (sd *singletonDatabase) PrepareNamedGet(result any, query string, arg any) error {
// 	sd.l.Log().Timestamp().Str("query", fmt.Sprintf("%s,%+v", query, arg)).Send()
// 	stmt, err := sd.DB.PrepareNamed(query)
// 	defer stmt.Close()
// 	if err != nil {
// 		return err
// 	}
// 	return stmt.Get(result, arg)
// }

// func (sd *singletonDatabase) WithTimeoutTxx() (*sqlx.Tx, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
// 	defer cancel()

// 	sd.DB.Connect(context.Background())

// 	tx, err := sd.DB.BeginTxx(ctx, nil)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "database instance WithTimeoutTxx BeginTxx error")
// 	}

// 	return tx, nil
// }

// func(sd *singletonDatabase) Queryx(query string, args ...any) (*sqlx.Rows, error) {

// }

// ####### Test #######
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
