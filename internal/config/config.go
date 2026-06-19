package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Kusa_www_Username             string
	Kusa_www_Timezone             string
	Kusa_www_MinimumContributions int
	Kusa_www_DryRun               bool
}

func Load() (Config, error) {
	username := os.Getenv("KUSA_WWW_USERNAME")
	if username == "" {
		return Config{}, fmt.Errorf("KUSA_WWW_USERNAME is required")
	}
	timezone := os.Getenv("Kusa_www_TIMEZONE")
	if timezone == "" {
		return Config{}, fmt.Errorf("KUSA_WWW_TIMEZONE is required")
	}
	minimumContributionsStr := os.Getenv("KUSA_WWW_MINIMUM_CONTRIBUTIONS")
	if minimumContributionsStr == "" {
		return Config{}, fmt.Errorf("KUSA_WWW_MINIMUM_CONTRIBUTIONS is required")
	}
	minimumContributions, err := strconv.Atoi(minimumContributionsStr)
	if err != nil {
		return Config{}, fmt.Errorf("KUSA_WWW_MINIMUM_CONTRIBUTIONS must be an integer: %w", err)
	}
	dryRunStr := os.Getenv("KUSA_WWWD_RY_RUN")
	if dryRunStr == "" {
		return Config{}, fmt.Errorf("KUSA_WWW_DRY_RUN is required")
	}
	dryRun, err := strconv.ParseBool(dryRunStr)
	if err != nil {
		return Config{}, fmt.Errorf("KUSA_WWW_DRY_RUN must be true or false: %w", err)
	}
	return Config{
		Kusa_www_Username:             username,
		Kusa_www_Timezone:             timezone,
		Kusa_www_MinimumContributions: minimumContributions,
		Kusa_www_DryRun:               dryRun,
	}, nil

}
