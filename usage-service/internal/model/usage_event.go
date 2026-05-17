package model

import "github.com/seakee/cpa-manager/usage-service/internal/usage"

type UsageEvent = usage.Event

type InsertResult struct {
	Inserted int `json:"inserted"`
	Skipped  int `json:"skipped"`
}
