package shipengine_api

import (
	"fmt"
	"os"
	"testing"
)

func TestTracking(t *testing.T) {
	shipEngine := New(os.Getenv("API_KEY"))

	tracking, err := shipEngine.TrackPackage(STAMPS, "9405511899223197428490")
	if err != nil {
		t.Errorf("Error tracking package: %s", err.Error())
	}
	fmt.Println(tracking.StatusCode)
}
