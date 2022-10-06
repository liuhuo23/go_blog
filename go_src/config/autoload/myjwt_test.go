package autoload

import (
	uuid "github.com/satori/go.uuid"
	"testing"
	"time"
)

func Test_Jwt(t *testing.T) {
	userid, _ := uuid.FromString("0000-0000-0000-0000")
	token, err := GenerateToken(userid, time.Now().Add(time.Second))
	if err != nil {
		return
	}
	t.Log(token)
}
