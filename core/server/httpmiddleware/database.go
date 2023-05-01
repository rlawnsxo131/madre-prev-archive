package httpmiddleware

import (
	"net/http"

	"github.com/rlawnsxo131/madre-server-v3/core/datastore/rdb"
)

// TODO connection 을 ctx 에 전달하는게 아니면 의미가 크게 없을것 같기도 해서,
// 어차피 singleton 이니 route 세팅할때 넣어줄지 고민해보기
func Database(db rdb.SingletonDatabase) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dbCtx := rdb.SetDBCtx(
				r.Context(),
				db,
			)

			next.ServeHTTP(
				w,
				r.WithContext(dbCtx),
			)
		})
	}
}
