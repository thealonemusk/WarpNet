package services

import (
	"context"
	"time"

	"github.com/thealonemusk/WarpNet/pkg/node"
	"github.com/thealonemusk/WarpNet/pkg/protocol"
	"github.com/thealonemusk/WarpNet/pkg/utils"

	"github.com/thealonemusk/WarpNet/pkg/blockchain"
)

func AliveNetworkService(announcetime, scrubTime, maxtime time.Duration) node.NetworkService {
	return func(ctx context.Context, c node.Config, n *node.Node, b *blockchain.Ledger) error {
		t := time.Now()
		// By announcing periodically our service to the blockchain
		b.Announce(
			ctx,
			announcetime,
			func() {
				// Keep-alive
				b.Add(protocol.HealthCheckKey, map[string]interface{}{
					n.Host().ID().String(): time.Now().UTC().Format(time.RFC3339),
				})

				// Keep-alive scrub
				nodes := AvailableNodes(b, maxtime)
				if len(nodes) == 0 {
					return
				}
				lead := utils.Leader(nodes)
				if !t.Add(scrubTime).After(time.Now()) {
					// Update timer so not-leader do not attempt to delete bucket afterwards
					// prevent cycles
					t = time.Now()

					if lead == n.Host().ID().String() {
						// Automatically scrub after some time passed
						b.DeleteBucket(protocol.HealthCheckKey)
					}
				}
			},
		)
		return nil
	}
}

// Alive announce the node every announce time, with a periodic scrub time for healthchecks
// the maxtime is the time used to determine when a node is unreachable (after maxtime, its unreachable)
func Alive(announcetime, scrubTime, maxtime time.Duration) []node.Option {
	return []node.Option{
		node.WithNetworkService(AliveNetworkService(announcetime, scrubTime, maxtime)),
	}
}

// AvailableNodes returns the available nodes which sent a healthcheck in the last maxTime
func AvailableNodes(b *blockchain.Ledger, maxTime time.Duration) (active []string) {
	for u, t := range b.LastBlock().Storage[protocol.HealthCheckKey] {
		var s string
		t.Unmarshal(&s)
		parsed, _ := time.Parse(time.RFC3339, s)
		if parsed.Add(maxTime).After(time.Now().UTC()) {
			active = append(active, u)
		}
	}

	return active
}
