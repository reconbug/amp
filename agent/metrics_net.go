package core

import (
	"time"

	"github.com/appcelerator/amp/api/rpc/stats"
	"github.com/docker/docker/api/types"
)

// NetStats net stats
type NetStats struct {
	Time      time.Time
	RxBytes   uint64
	RxDropped uint64
	RxErrors  uint64
	RxPackets uint64
	TxBytes   uint64
	TxDropped uint64
	TxErrors  uint64
	TxPackets uint64
}

// NetStatsDiff diff between two IOStats
type NetStatsDiff struct {
	Duration  int64
	RxBytes   int64
	RxDropped int64
	RxErrors  int64
	RxPackets int64
	TxBytes   int64
	TxDropped int64
	TxErrors  int64
	TxPackets int64
}

// publish one net metrics event
func (a *Agent) setNetMetrics(data *ContainerData, statsData *types.StatsJSON, entry *stats.MetricsEntry) {
	net := a.newNetStats(statsData)
	if data.previousNetStats == nil {
		data.previousNetStats = net
		return
	}
	diff := a.newNetDiff(net, data.previousNetStats)
	if diff == nil {
		return
	}
	data.previousNetStats = net
	entry.Net = &stats.MetricsNetEntry{
		TotalBytes: diff.RxBytes + diff.TxBytes,
		RxBytes:    diff.RxBytes,
		RxDropped:  diff.RxDropped,
		RxErrors:   diff.RxErrors,
		RxPackets:  diff.RxPackets,
		TxBytes:    diff.TxBytes,
		TxDropped:  diff.TxDropped,
		TxErrors:   diff.TxErrors,
		TxPackets:  diff.TxPackets,
	}
}

// create a new net stats
func (a *Agent) newNetStats(stats *types.StatsJSON) *NetStats {
	var net = &NetStats{Time: stats.Read}
	for _, netStats := range stats.Networks {
		net.RxBytes += netStats.RxBytes
		net.RxDropped += netStats.RxDropped
		net.RxErrors += netStats.RxErrors
		net.RxPackets += netStats.RxPackets
		net.TxBytes += netStats.TxBytes
		net.TxDropped += netStats.TxDropped
		net.TxErrors += netStats.TxErrors
		net.TxPackets += netStats.TxPackets
	}
	return net
}

// create a new net diff computing difference between two net stats
func (a *Agent) newNetDiff(newNet *NetStats, previousNet *NetStats) *NetStatsDiff {
	diff := &NetStatsDiff{Duration: int64(newNet.Time.Sub(previousNet.Time).Seconds())}
	if diff.Duration <= 0 {
		return nil
	}
	diff.RxBytes = int64(newNet.RxBytes - previousNet.RxBytes)
	diff.RxDropped = int64(newNet.RxDropped - previousNet.RxDropped)
	diff.RxErrors = int64(newNet.RxErrors - previousNet.RxErrors)
	diff.RxPackets = int64(newNet.RxPackets - previousNet.RxPackets)
	diff.TxBytes = int64(newNet.TxBytes - previousNet.TxBytes)
	diff.TxDropped = int64(newNet.TxDropped - previousNet.TxDropped)
	diff.TxErrors = int64(newNet.TxErrors - previousNet.TxErrors)
	diff.TxPackets = int64(newNet.TxPackets - previousNet.TxPackets)
	return diff
}