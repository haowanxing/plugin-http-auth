package http_auth

import (
	"encoding/json"
	"errors"
	. "m7s.live/engine/v4"
	"m7s.live/engine/v4/util"
)

func (p *HttpAuthConfig) checkPubAuthResult(promise *util.Promise[IPublisher]) {
	var auth bool
	if len(p.OnPubAddr) == 0 {
		auth = true
	} else {
		var publisher = promise.Value.GetPublisher()
		stream := publisher.Stream
		requestData := &req{
			Action:   "publish",
			App:      stream.AppName,
			Stream:   stream.StreamName,
			Param:    publisher.Args.Encode(),
			ClientID: publisher.ID,
		}
		if buf, err := json.Marshal(requestData); err == nil {
			auth = p.checkAPIOK(p.OnPubAddr, buf)
		}
	}
	if !auth {
		promise.Reject(errors.New("publish auth failed"))
	} else {
		promise.Resolve()
	}
}
