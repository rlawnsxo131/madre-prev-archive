package lib

import (
	"compress/flate"
	"compress/gzip"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/rlawnsxo131/madre-server-v2/constants"
)

func ResponseJsonCompressWriter(rw http.ResponseWriter, r *http.Request, data interface{}) {
	json, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	if len(json) >= 2048 {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			gz, _ := gzip.NewWriterLevel(rw, gzip.DefaultCompression)
			defer gz.Close()
			rw.Header().Set("Content-Encoding", "gzip")
			rw.WriteHeader(http.StatusOK)
			gz.Write(json)
			return
		}
		if strings.Contains(r.Header.Get("Accept-Encoding"), "deflate") {
			df, _ := flate.NewWriter(rw, flate.DefaultCompression)
			defer df.Close()
			rw.Header().Set("Content-Encoding", "deflate")
			rw.WriteHeader(http.StatusOK)
			df.Write(json)
			return
		}
	}
	rw.Write(json)
}

func ResponseErrorWriter(rw http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	message := constants.InternalServerErrorMessage

	if err == sql.ErrNoRows {
		status = http.StatusNotFound
		message = constants.NotFoundErrorMessage
	}

	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(map[string]interface{}{"status": status, "message": message})
	panic(err)
}
