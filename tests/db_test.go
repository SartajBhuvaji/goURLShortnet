package tests

import (
	"testing"

	"github.com/SartajBhuvaji/utils"
)

func TestSetupRedis(t *testing.T) {
	redisClient, err := utils.SetupRedis()
	if err != nil {
		t.Fatalf("Failed to get Reddis Client %v", err)
	}
	defer redisClient.Close()
}
