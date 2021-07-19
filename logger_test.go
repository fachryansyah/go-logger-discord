package gologgerdiscord

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogError(t *testing.T) {

	l := Init(AuthConfig{
		Token: "your token",
		Type:  "Bot",
	}, "your channel id")

	err := l.Error(
		"500 Internal server error (http://localhost:1337/test/error)",
		"Error structs.go on line 32",
		"cannot use jsonValue (type []byte) as type io.Reader in argument to http.Post: []byte does not implement io.Reader (missing Read method)",
		"https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go/24455606#24455606",
	)
	assert.Equal(t, err, nil)

	err = l.Warn(
		"Test Warning log",
		"Im a warning log",
		"cannot use jsonValue (type []byte) as type io.Reader in argument to http.Post: []byte does not implement io.Reader (missing Read method)",
		"https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go/24455606#24455606",
	)
	assert.Equal(t, err, nil)

	err = l.Info(
		"Test Info log",
		"Im a info log",
		"cannot use jsonValue (type []byte) as type io.Reader in argument to http.Post: []byte does not implement io.Reader (missing Read method)",
		"https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go/24455606#24455606",
	)
	assert.Equal(t, err, nil)

	err = l.Success(
		"Test Succes log",
		"Im a success log",
		"Build succesfully, took 32s",
		"https://localhost:1337",
	)
	assert.Equal(t, err, nil)

}
