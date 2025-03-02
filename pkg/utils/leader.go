package utils

import "hash/fnv"

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func Leader(actives []string) string {
	// first get available nodes
	leaderboard := map[string]uint32{}

	leader := actives[0]

	// Compute who is leader at the moment
	for _, a := range actives {
		leaderboard[a] = hash(a)
		if leaderboard[leader] < leaderboard[a] {
			leader = a
		}
	}
	return leader
}
