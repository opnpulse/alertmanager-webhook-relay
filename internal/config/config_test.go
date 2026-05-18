package config

import "testing"

func TestLoadFromEnvParsesDedupeWindowSeconds(t *testing.T) {
	t.Setenv("WEBHOOK_URL", "https://example.com/webhook")
	t.Setenv("DEDUPE_WINDOW_SECONDS", "45")

	cfg, err := LoadFromEnv()
	if err != nil {
		t.Fatalf("expected config to load, got %v", err)
	}
	if got := int(cfg.DedupeWindow.Seconds()); got != 45 {
		t.Fatalf("expected dedupe window of 45 seconds, got %d", got)
	}
}
