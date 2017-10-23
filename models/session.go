package models

import (
	"github.com/astaxie/beego/session"
)

var (
	globalSessions *session.Manager
)

func Session() *session.Manager {
	if globalSessions == nil {
		globalSessions = NewSession()
	}
	return globalSessions
}

func NewSession() *session.Manager {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "SessionID",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	globalSessions, _ = session.NewManager("memory", sessionConfig)
	go globalSessions.GC()
	return globalSessions
}
