package http_auth

import (
	"bytes"
	"go.uber.org/zap"
	. "m7s.live/engine/v4"
	"m7s.live/engine/v4/util"
	"net/http"
)

func (p *HttpAuthConfig) changeAuthHook() {
	OnAuthSub = func(promise *util.Promise[ISubscriber]) error {
		go p.checkSubAuthResult(promise)
		return nil
	}
	OnAuthPub = func(promise *util.Promise[IPublisher]) error {
		go p.checkPubAuthResult(promise)
		return nil
	}
}

func (p *HttpAuthConfig) checkAPIOK(addr string, buf []byte) bool {
	plugin.Info("Auth", zap.String("addr", addr))
	if len(addr) == 0 {
		return true
	} else {
		if resp, err := http.Post(addr, "application/json", bytes.NewBuffer(buf)); err == nil {
			if resp.StatusCode == http.StatusOK {
				return true
			}
		}
	}
	return false
}
