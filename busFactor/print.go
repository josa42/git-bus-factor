package busFactor

import (
	"fmt"
)

// Print :
func Print() {
	// token := github.Token()

	fmt.Printf("🍴  %d forks.\n", 0)
	fmt.Printf("🔭  %d watchers.\n", 0)
	fmt.Printf("🌟  %d stars.\n", 0)
	fmt.Printf("📆  Created about %s ago; last push %s ago.\n", "6 years", "4 days")
	fmt.Printf("🍻  0 PRs: %d closed; %d open; %d%% are closed.\n", 0, 0, 100)
	fmt.Printf("🛠️  Deletions to additions ratio: %f%% (%d/%d).\n", 86.99, -1324, 123412)
	fmt.Printf("📦  0 releases; lates release \"%s\": %s ago.\n", "v0.0.0", "2 months")
	fmt.Printf("🚌  Bus factor: 33.33%% (3 impactful contributors out of 100).\n")
}
