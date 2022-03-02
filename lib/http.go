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
)

type HttpWriter interface {
	ErrorWriter(err error)
	CompressWriter(data interface{})
}

type httpWriter struct {
	rw http.ResponseWriter
	r  *http.Request
}

func NewHttpWriter(rw http.ResponseWriter, r *http.Request) HttpWriter {
	return &httpWriter{
		rw: rw,
		r:  r,
	}
}

func (writer *httpWriter) ErrorWriter(err error) {
	status := http.StatusInternalServerError
	message := constants.ErrInternalServerMessage

	if err == sql.ErrNoRows {
		status = http.StatusNotFound
		message = constants.ErrNotFoundMessage
	}

	writer.rw.WriteHeader(status)
	json.NewEncoder(writer.rw).Encode(map[string]interface{}{"status": status, "message": message})
	log.Printf("%+v", err)
}

func (writer *httpWriter) CompressWriter(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		err = errors.Wrap(err, "ResponseJsonCompressWriter: json parse error")
		writer.ErrorWriter(err)
	}

	if len(jsonData) >= 2048 {
		if strings.Contains(writer.r.Header.Get("Accept-Encoding"), "gzip") {
			gz, err := gzip.NewWriterLevel(writer.rw, gzip.DefaultCompression)
			if err != nil {
				err = errors.Wrap(err, "ResponseJsonCompressWriter: gzip compress error")
				writer.ErrorWriter(err)
				return
			}
			defer gz.Close()
			writer.rw.Header().Set("Content-Encoding", "gzip")
			writer.rw.WriteHeader(http.StatusOK)
			gz.Write(jsonData)
			return
		}
		if strings.Contains(writer.r.Header.Get("Accept-Encoding"), "deflate") {
			df, err := flate.NewWriter(writer.rw, flate.DefaultCompression)
			if err != nil {
				err = errors.Wrap(err, "ResponseJsonCompressWriter: dfalte compress error")
				writer.ErrorWriter(err)
				return
			}
			defer df.Close()
			writer.rw.Header().Set("Content-Encoding", "deflate")
			writer.rw.WriteHeader(http.StatusOK)
			df.Write(jsonData)
		}
	}

	writer.rw.Write(jsonData)
}
