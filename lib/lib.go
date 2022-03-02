package lib

import (
	"compress/flate"
	"compress/gzip"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/constants"
	uuid "github.com/satori/go.uuid"
)

/**
 * http
 */
// func httpLogger(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		len := r.ContentLength
// 		body := make([]byte, len)
// 		r.Body.Read(body)
// 		log.Printf("query: %v", r.URL.RawQuery)
// 		log.Printf("body: %v", string(body))
// 		next.ServeHTTP(w, r)
// 	})
// }

func ResponseErrorWriter(rw http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	message := constants.InternalServerErrorMessage

	if err == sql.ErrNoRows {
		status = http.StatusNotFound
		message = constants.NotFoundErrorMessage
	}

	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(map[string]interface{}{"status": status, "message": message})
	log.Printf("%+v", err)
}

func ResponseJsonCompressWriter(rw http.ResponseWriter, r *http.Request, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		err = errors.Wrap(err, "ResponseJsonCompressWriter: json parse error")
		ResponseErrorWriter(rw, err)
		return
	}

	// When an error occurs in the compress process, should I change it to return uncompressed json?
	if len(jsonData) >= 2048 {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			gz, err := gzip.NewWriterLevel(rw, gzip.DefaultCompression)
			if err != nil {
				err = errors.Wrap(err, "ResponseJsonCompressWriter: gzip compress error")
				ResponseErrorWriter(rw, err)
				return
			}
			defer gz.Close()
			rw.Header().Set("Content-Encoding", "gzip")
			rw.WriteHeader(http.StatusOK)
			gz.Write(jsonData)
			return
		}
		if strings.Contains(r.Header.Get("Accept-Encoding"), "deflate") {
			df, err := flate.NewWriter(rw, flate.DefaultCompression)
			if err != nil {
				err = errors.Wrap(err, "ResponseJsonCompressWriter: dfalte compress error")
				ResponseErrorWriter(rw, err)
				return
			}
			defer df.Close()
			rw.Header().Set("Content-Encoding", "deflate")
			rw.WriteHeader(http.StatusOK)
			df.Write(jsonData)
			return
		}
	}

	rw.Write(jsonData)
}

/**
 * uuid
 * github.com/satori/go.uuid
 */
func GenerateUUID() string {
	uuid := uuid.NewV4()
	return uuid.String()
}

/**
 * default value
 */
func IfIsNotExistGetDefaultIntValue(value int, defaultValue int) int {
	if value == 0 {
		value = defaultValue
	}
	return value
}
