package slashcmd

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const (
	testToken    = "testToken"
	testTeamId   = "testTeamId"
	testDomain   = "testDomain"
	testChanId   = "testChanId"
	testChanName = "testChanName"
	testUserId   = "testUserId"
	testUserName = "testUserName"
	testCmd      = "testCmd"
	testText     = "testText"
	testRespUrl  = "testResponseUrl"
	testTrigId   = "testTrigId"
	testQueryFmt = `%s=%s`
)

var (
	testQueryMap = map[string]string{
		tokenParam:    testToken,
		teamIdParam:   testTeamId,
		teamDomParam:  testDomain,
		chanIdParam:   testChanId,
		chanNameParam: testChanName,
		userIdParam:   testUserId,
		userNameParam: testUserName,
		cmdParam:      testCmd,
		textParam:     testText,
		respUrlParam:  testRespUrl,
		trigIdParam:   testTrigId,
	}
	testQueryStr = generateSlashCmdStr(&testQueryMap)
)

func generateSlashCmdStr(m *map[string]string) string {
	resB := strings.Builder{}
	ctr := 0
	for key, val := range *m {
		if ctr >= 1 {
			fmt.Fprintf(&resB, testQueryFmt, key, val)
			ctr--
		}
		fmt.Fprintf(&resB, "&"+testQueryFmt, key, val)
	}

	return resB.String()
}

func TestParse(t *testing.T) {
	slashCmd, err := Parse(testQueryStr)
	assert.Nil(t, err)
	assert.Equal(t, slashCmd.token, testToken)
	assert.Equal(t, slashCmd.teamId, testTeamId)
	assert.Equal(t, slashCmd.teamDomain, testDomain)
	assert.Equal(t, slashCmd.channelId, testChanId)
	assert.Equal(t, slashCmd.channelName, testChanName)
	assert.Equal(t, slashCmd.userId, testUserId)
	assert.Equal(t, slashCmd.userName, testUserName)
	assert.Equal(t, slashCmd.command, testCmd)
	assert.Equal(t, slashCmd.text, testText)
	assert.Equal(t, slashCmd.responseUrl, testRespUrl)
	assert.Equal(t, slashCmd.triggerId, testTrigId)
}

func TestParseErr(t *testing.T) {
	slashCmd, err := Parse("erred%query%string")
	assert.NotNil(t, err)
	assert.Nil(t, slashCmd)
}
