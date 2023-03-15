package http_auth

import (
	. "m7s.live/engine/v4"
	"m7s.live/engine/v4/config"
)

type req struct {
	Action   string `json:"action"`
	App      string `json:"app"`
	Stream   string `json:"stream"`
	Param    string `json:"param"`
	ClientID string `json:"client_id"`
}

type HttpAuthConfig struct {
	OnPubAddr string
	OnSubAddr string
}

func (p *HttpAuthConfig) OnEvent(event any) {
	switch event.(type) {
	case FirstConfig: //插件初始化逻辑
		p.changeAuthHook()
	case config.Config: //插件热更新逻辑
	}
}

var plugin = InstallPlugin(new(HttpAuthConfig))
