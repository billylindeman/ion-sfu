package sfu

import (
	"fmt"
	"strings"

	log "github.com/pion/ion-log"
)

// SignalConfig params for http listener / grpc / websocket server
type SignalConfig struct {
	FQDN     string
	Key      string
	Cert     string
	HTTPAddr string
	GRPCAddr string
}

// CoordinatorConfig is used to coordinate sessions / relay across multiple nodes
type CoordinatorConfig struct {
	Local *struct {
		Enabled bool
	}
	Etcd *struct {
		Enabled bool
		Hosts   []string
	}
}

// ICEServerConfig defines parameters for ice servers
type ICEServerConfig struct {
	URLs       []string `mapstructure:"urls"`
	Username   string   `mapstructure:"username"`
	Credential string   `mapstructure:"credential"`
}

type Candidates struct {
	IceLite    bool     `mapstructure:"icelite"`
	NAT1To1IPs []string `mapstructure:"nat1to1"`
}

// WebRTCConfig defines parameters for ice
type WebRTCConfig struct {
	ICEPortRange []uint16          `mapstructure:"portrange"`
	ICEServers   []ICEServerConfig `mapstructure:"iceserver"`
	Candidates   Candidates        `mapstructure:"candidates"`
	SDPSemantics string            `mapstructure:"sdpsemantics"`
}

// Config for base SFU
type Config struct {
	SFU struct {
		Ballast int64 `mapstructure:"ballast"`
	} `mapstructure:"sfu"`
	WebRTC      WebRTCConfig      `mapstructure:"webrtc"`
	Log         log.Config        `mapstructure:"log"`
	Router      RouterConfig      `mapstructure:"router"`
	Signal      SignalConfig      `mapstructure:"signal"`
	Coordinator CoordinatorConfig `mapstructure:"coordinator"`
}

// NodeURL public endpoint to hit
func (c *Config) NodeURL() string {
	port := strings.Split(c.Signal.HTTPAddr, ":")[1]

	if c.Signal.Key != "" && c.Signal.Cert != "" {
		return fmt.Sprintf("wss://%v:%v", c.Signal.FQDN, port)
	}
	return fmt.Sprintf("ws://%v:%v", c.Signal.FQDN, port)
}
