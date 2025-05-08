package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	testCases := []struct {
		desc     string
		params   map[string]string
		expected map[string]string
	}{
		{
			desc:   "fully default config",
			params: map[string]string{},
			expected: map[string]string{
				redisCfg: defaultRedisCfg,
				pgCfg:    defaultPGCfg,
				portCfg:  defaultPortCfg,
			},
		},
		{
			desc: "fully custom config",
			params: map[string]string{
				redisCfg: "customRedisURL",
				pgCfg:    "customPGURL",
				portCfg:  "customPort",
			},
			expected: map[string]string{
				redisCfg: "customRedisURL",
				pgCfg:    "customPGURL",
				portCfg:  "customPort",
			},
		},
		{
			desc: "one custom config",
			params: map[string]string{
				redisCfg: "customRedisURL",
			},
			expected: map[string]string{
				redisCfg: "customRedisURL",
				pgCfg:    defaultPGCfg,
				portCfg:  defaultPortCfg,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			os.Clearenv()
			for k, v := range tC.params {
				os.Setenv(k, v)
			}

			cfg := Load()

			for k, v := range tC.expected {
				if k == redisCfg && v != cfg.RedisURL {
					t.Fatalf("expected: %s, received: %s", v, cfg.RedisURL)
				} else if k == pgCfg && v != cfg.PGURL {
					t.Fatalf("expected: %s, received: %s", v, cfg.PGURL)
				} else if k == portCfg && v != cfg.Port {
					t.Fatalf("expected: %s, received: %s", v, cfg.Port)
				}
			}
		})
	}
}
