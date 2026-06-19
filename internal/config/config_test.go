package config

import (
	"testing"
)

func TestLoadSuccess(t *testing.T) {
	t.Setenv("KUSA_WWW_USERNAME", "test")
	t.Setenv("KUSA_WWW_TIMEZONE", "Asia/Tokyo")
	t.Setenv("KUSA_WWW_MINIMUM_CONTRIBUTIONS", "10")
	t.Setenv("KUSA_WWW_DRY_RUN", "true")
	cfg, err := Load()
	if err != nil {
		t.Fatal("load is failed", err)
	}
	if cfg.Kusa_www_Username != "test" {
		t.Fatal("KUSA_WWW_Username is wrong")
	}
	if cfg.Kusa_www_Timezone != "Asia/Tokyo" {
		t.Fatal("KUSA_WWW_Timezone is wrong")
	}
	if cfg.Kusa_www_MinimumContributions != 10 {
		t.Fatal("MINIMUM_CONTRIBUTIONS is Wrong")
	}
	if !cfg.Kusa_www_DryRun {
		t.Fatalf("got %v, want %v", cfg.Kusa_www_DryRun, true)
	}
}
func TestUsernameMissing(t *testing.T) {
	t.Setenv("KUSA_WWW_USERNAME", "")
	t.Setenv("KUSA_WWW_TIMEZONE", "Asia/Tokyo")
	t.Setenv("KUSA_WWW_MINIMUM_CONTRIBUTIONS", "10")
	t.Setenv("KUSA_WWW_DRY_RUN", "true")
	_, err := Load()
	if err == nil {
		t.Fatal("Load() should return an error when KUSA_WWW_USERNAME is missing")
	}
}
func TestTimezoneMissing(t *testing.T) {
	t.Setenv("KUSA_WWW_USERNAME", "test")
	t.Setenv("KUSA_WWW_TIMEZONE", "")
	t.Setenv("KUSA_WWW_MINIMUM_CONTRIBUTIONS", "10")
	t.Setenv("KUSA_WWW_DRY_RUN", "true")
	_, err := Load()
	if err == nil {
		t.Fatal("Load() should return an error when KUSA_WWW_TIMEZONE is missing")
	}
}
func TestMinimum_contributionsInvalid(t *testing.T) {
	t.Setenv("KUSA_WWW_USERNAME", "test")
	t.Setenv("KUSA_WWW_TIMEZONE", "Asia/Tokyo")
	t.Setenv("KUSA_WWW_MINIMUM_CONTRIBUTIONS", "hoge")
	t.Setenv("KUSA_WWW_DRY_RUN", "true")
	_, err := Load()
	if err == nil {
		t.Fatal("Load() should return an error when KUSA_WWW_MINIMUM_CONTRIBUTIONS is invalid")
	}
}
func TestDry_runInvalid(t *testing.T) {
	t.Setenv("KUSA_WWW_USERNAME", "test")
	t.Setenv("KUSA_WWW_TIMEZONE", "Asia/Tokyo")
	t.Setenv("KUSA_WWW_MINIMUM_CONTRIBUTIONS", "10")
	t.Setenv("KUSA_WWW_DRY_RUN", "hoge")
	_, err := Load()
	if err == nil {
		t.Fatal("Load() should return an error when KUSA_WWW_DRY_RUN is invalid")
	}
}
