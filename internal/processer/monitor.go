package processer

import (
	"time"

	"github.com/ThreatCode/Proxy-Check/internal"
)

func Monitor() {
	for {
		// Sleep for interval seconds
		time.Sleep(time.Duration(internal.Options.MonitorInterval) * time.Second)
		// Take a snapshot
		internal.State.Snapshot()
	}
}
