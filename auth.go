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
		if len(promise.Value.GetSubscriber().ID) == 0 {
			promise.Resolve()
			return nil
		}
		go p.checkSubAuthResult(promise)
		return nil
	}
	OnAuthPub = func(promise *util.Promise[IPublisher]) error {
		if len(promise.Value.GetPublisher().ID) == 0 {
			promise.Resolve()
			return nil
		}
		go p.checkPubAuthResult(promise)
		return nil
	}
}

func (p *HttpAuthConfig) checkAPIOK(addr string, buf []byte) bool {
	if len(addr) == 0 {
		return true
	} else {
		if resp, err := http.Post(addr, "application/json", bytes.NewBuffer(buf)); err == nil {
			plugin.Info("auth resp", zap.String("addr", addr), zap.Int("http_code", resp.StatusCode))
			if resp.StatusCode == http.StatusOK {
				return true
			}
		} else {
			plugin.Error("auth req fail", zap.Error(err))
		}
	}
	return false
}
