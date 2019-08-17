package main

import (
	"sort"
)

// GameTribeOutput is the JSON structure for the toptribes list
type GameTribeOutput struct {
	TribeID   uint64 `json:"tribeID"`
	TribeName string `json:"tribeName"`
	Index     int    `json:"index"`
}

// TribeCount holds the per tribe number of markers
type TribeCount struct {
	tribeID uint64
	count   uint32
}

func TopNTribes(n int, counts map[uint64]*TribeCount) []uint64 {

	countsSlice := make([]*TribeCount, 0, len(counts))
	for _, v := range counts {
		val := v
		countsSlice = append(countsSlice, val)
	}

	sort.Slice(countsSlice, func (i, j int) bool {
		return countsSlice[i].count > countsSlice[j].count
	})

	var results []uint64
	totalAmt := Min(n, len(countsSlice))
	for _, fc := range countsSlice[:totalAmt] {
		count := fc
		results = append(results, count.tribeID)
	}

	return results
}

