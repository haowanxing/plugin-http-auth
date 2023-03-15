package http_auth

import (
	"encoding/json"
	"errors"
	. "m7s.live/engine/v4"
	"m7s.live/engine/v4/util"
)

func (p *HttpAuthConfig) checkSubAuthResult(promise *util.Promise[ISubscriber]) {
	var auth bool
	if len(p.OnSubAddr) == 0 {
		auth = true
	} else {
		var subscriber = promise.Value.GetSubscriber()
		stream := subscriber.Stream
		requestData := &req{
			Action:   "subscribe",
			App:      stream.AppName,
			Stream:   stream.StreamName,
			Param:    subscriber.Args.Encode(),
			ClientID: subscriber.ID,
		}
		if buf, err := json.Marshal(requestData); err == nil {
			auth = p.checkAPIOK(p.OnSubAddr, buf)
		}
	}
	if !auth {
		promise.Reject(errors.New("subscribe auth failed"))
	} else {
		promise.Resolve()
	}
}
