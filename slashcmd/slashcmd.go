package slashcmd

import (
	"net/url"
)

const (
	tokenParam    = "token"
	teamIdParam   = "team_id"
	teamDomParam  = "team_domain"
	chanIdParam   = "channel_id"
	chanNameParam = "channel_name"
	userIdParam   = "user_id"
	userNameParam = "user_name"
	cmdParam      = "command"
	textParam     = "text"
	respUrlParam  = "response_url"
	trigIdParam   = "trigger_id"
)

type SlashCmd struct {
	token       string
	teamId      string
	teamDomain  string
	channelId   string
	channelName string
	userId      string
	userName    string
	command     string
	text        string
	responseUrl string
	triggerId   string
}

func getValue(param string, vals url.Values) string {
	valList, valOk := vals[param]
	if valOk && len(valList) > 0 {
		return valList[0]
	}
	return ""
}

// Parse takes a raw url query string and parses it into a SlashCmd
// struct. It expects a string like the following:
//
// token=<token>
//      &team_id=<team_id>
//      &team_domain=<team_domain>
//      &channel_id=<channel_id>
//      &channel_name=<channel_name>
//      &user_id=<user_id>
//      &user_name=<user_name>
//      &command=<command>
//      &text=<text>
//      &response_url=<response_url>
//      &trigger_id=<trigger_id>
//
// If the string can't be parsed, an error is returned along with a nil
// SlashCmd struct. Otherwise, the SlashCmd is returned. No validation
// is performed on the values of the struct.
func Parse(rawVals string) (*SlashCmd, error) {
	vals, err := url.ParseQuery(rawVals)
	if err != nil {
		return nil, err
	}

	return &SlashCmd{
		token:       getValue(tokenParam, vals),
		teamId:      getValue(teamIdParam, vals),
		teamDomain:  getValue(teamDomParam, vals),
		channelId:   getValue(chanIdParam, vals),
		channelName: getValue(chanNameParam, vals),
		userId:      getValue(userIdParam, vals),
		userName:    getValue(userNameParam, vals),
		command:     getValue(cmdParam, vals),
		text:        getValue(textParam, vals),
		responseUrl: getValue(respUrlParam, vals),
		triggerId:   getValue(trigIdParam, vals),
	}, nil
}
