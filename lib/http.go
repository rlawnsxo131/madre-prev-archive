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
)

const (
	ErrNotFoundMessage       = "NOT_FOUND"             // 404
	ErrInternalServerMessage = "INTERNAL_SERVER_ERROR" // 500
)

type HttpWriter interface {
	WriteError(err error)
	WriteCompress(data interface{})
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

func (writer *httpWriter) WriteError(err error) {
	status := http.StatusInternalServerError
	message := ErrInternalServerMessage

	if err == sql.ErrNoRows {
		status = http.StatusNotFound
		message = ErrNotFoundMessage
	}

	writer.rw.WriteHeader(status)
	json.NewEncoder(writer.rw).Encode(map[string]interface{}{"status": status, "message": message})
	log.Printf("%+v", err)
}

func (writer *httpWriter) WriteCompress(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		err = errors.Wrap(err, "ResponseJsonWriteCompress: json parse error")
		writer.WriteError(err)
	}

	// When an error occurs in the compress process, should I change it to return uncompressed json?
	if len(jsonData) >= 2048 {
		if strings.Contains(writer.r.Header.Get("Accept-Encoding"), "gzip") {
			gz, err := gzip.NewWriterLevel(writer.rw, gzip.DefaultCompression)
			if err != nil {
				err = errors.Wrap(err, "ResponseJsonWriteCompress: gzip compress error")
				writer.WriteError(err)
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
				err = errors.Wrap(err, "ResponseJsonWriteCompress: dfalte compress error")
				writer.WriteError(err)
				return
			}
			defer df.Close()
			writer.rw.Header().Set("Content-Encoding", "deflate")
			writer.rw.WriteHeader(http.StatusOK)
			df.Write(jsonData)
			return
		}
	}

	writer.rw.Write(jsonData)
}
