package npncontroller

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"emperror.dev/errors"
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/libnpn/npnweb"
	"golang.org/x/text/language"
)

// Used by RespondJSON
type JSONResponse struct {
	Status   string    `json:"status"`
	Message  string    `json:"message"`
	Path     string    `json:"path"`
	Occurred time.Time `json:"occurred"`
}

// Constructs a new JSONResponse
func NewJSONResponse(status string, msg string, path string) *JSONResponse {
	return &JSONResponse{Status: status, Message: msg, Path: path, Occurred: time.Now()}
}

type errorResult struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Helper for transforming the result of a template
func T(_ int, err error) (string, error) {
	return "", err
}

// Respond with an error and optional messages
func EResp(err error, msgs ...string) (string, error) {
	msg := strings.Join(msgs, "\n")
	if len(msg) == 0 {
		return "", err
	}
	return "", errors.Wrap(err, msg)
}

// Respond with an error based on the provided message
func ENew(msg string) (string, error) {
	return "", errors.New(msg)
}

// Serialize body to JSON, and respond with correct MIME type
func RespondJSON(w http.ResponseWriter, filename string, body interface{}, logger *logrus.Logger) (string, error) {
	return RespondMIME(filename, "application/json", "json", npncore.ToJSONBytes(body, logger, true), w)
}

// Respond with provided byte array and MIME type
func RespondMIME(filename string, mime string, ext string, ba []byte, w http.ResponseWriter) (string, error) {
	w.Header().Set("Content-Type", mime+"; charset=UTF-8")
	if len(filename) > 0 {
		if !strings.HasSuffix(filename, "."+ext) {
			filename = filename + "." + ext
		}
		w.Header().Set("Content-Disposition", `attachment; filename="`+filename+`"`)
	}
	WriteCORS(w)
	if len(ba) == 0 {
		return "", errors.New("no bytes available to write")
	}
	_, err := w.Write(ba)
	return "", errors.Wrap(err, "cannot write to response")
}

// Writes CORS headers to a provided http.ResponseWriter
func WriteCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Method", "GET,POST,DELETE,PUT,PATCH,OPTIONS,HEAD")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func logComplete(startNanos int64, ctx *npnweb.RequestContext, status int, r *http.Request) {
	delta := npncore.TimerEnd(startNanos)
	ms := npncore.MicrosToMillis(language.AmericanEnglish, delta)
	args := map[string]interface{}{"elapsed": delta, npncore.KeyStatus: status}
	msg := fmt.Sprintf("%v %v returned [%v] in [%v]", r.Method, r.URL.Path, status, ms)
	ctx.Logger.Debug(msg, args)
}
