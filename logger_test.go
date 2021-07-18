package gologgerdiscord

import "testing"

func TestLogError(t *testing.T) {
	var log Logger
	log.Error(
		"500 Internal server error (http://localhost:1337/test/error)",
		"Error structs.go on line 32",
		"cannot use jsonValue (type []byte) as type io.Reader in argument to http.Post: []byte does not implement io.Reader (missing Read method)",
		"https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go/24455606#24455606",
		"Local",
	)
}
