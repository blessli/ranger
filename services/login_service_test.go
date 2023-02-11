package services

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_getJwtToken(t *testing.T) {
	accessExpire := int64(86400)
	now := time.Now().Unix()
	jwtToken, err := getJwtToken("ad879037-c7a4-4063-9236-6bfc35d54b7d", now, accessExpire, 1)
	assert.Nil(t, err)
	t.Log(jwtToken)
}

func Test_loginService_Login(t *testing.T) {
	s:=NewLoginService()
	rsp,err:=s.Login("admin","123456")
	assert.Nil(t,err)
	assert.NotNil(t,rsp)
	t.Logf("%+v",rsp)
}
