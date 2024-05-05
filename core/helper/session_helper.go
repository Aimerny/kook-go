package helper

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

// StoreSession save session id as local file
func StoreSession(sessionId string, maxSN int) {
	file, err := os.OpenFile("session_id", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Errorf("save session id failed, %e", err)
		return
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%s %d", sessionId, maxSN))
}

func ReloadSession() (string, int) {
	sessionIdAndSN, err := os.ReadFile("session_id")
	if err != nil {
		log.Errorf("read session id failed, %e", err)
		return "", 0
	}
	if len(sessionIdAndSN) == 0 {
		return "", 0
	}

	res := strings.Split(string(sessionIdAndSN), " ")
	maxSN, err := strconv.Atoi(res[1])
	if err != nil {
		log.Errorf("stored maxSN is invalid, set 0")
	}
	log.Infof("reload session id: %s, maxSN: %d", res[0], maxSN)
	return res[0], maxSN
}
