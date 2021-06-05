package config_test

import (
	"github.com/stretchr/testify/assert"
	"go.bnck.me/loki-auth-proxy/internal/config"
	"testing"
)

func TestGet(t *testing.T) {
	actualCfg := config.Get()

	assert.NotNil(t, actualCfg)

	expectedCfg := &config.Config{HTTP: struct {
		Listen string "default:\":8080\""
	}{Listen: ""}, Prometheus: false, Authenticator: struct {
		Name   string
		Config map[string]string
	}{Name: "", Config: map[string]string(nil)}, Log: struct {
		Level string "default:\"info\""
	}{Level: ""}}

	assert.Equal(t, expectedCfg, actualCfg)
}

func TestConfig_Load(t *testing.T) {
	cfg := config.Get()
	assert.NotNil(t, cfg)

	assert.NoError(t, cfg.Load("../../example/loki-auth-proxy.yaml"))

	expectedCfg := &config.Config{HTTP: struct {
		Listen string "default:\":8080\""
	}{Listen: ":8081"}, Prometheus: false, Backend: struct {
		Distributor   string
		QueryFrontend string "yaml:\"query_frontend\""
		Querier       string
	}{Distributor: "distributor:3100", QueryFrontend: "query-frontend:3100", Querier: "querier:3100"}, Authenticator: struct {
		Name   string
		Config map[string]string
	}{Name: "file", Config: map[string]string{"path": "example/loki-auth.yaml"}}, Log: struct {
		Level string "default:\"info\""
	}{Level: "debug"}}

	assert.Equal(t, expectedCfg, cfg)
}
