package main

import (
	"context"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/timestamp"
	"github.com/thanos-io/thanos/pkg/block/metadata"
	"github.com/thanos-io/thanos/pkg/testutil/e2eutil"
	"log"
	"strconv"
	"time"
)

func main() {
	ctx := context.Background()
	numSeries := 100000
	numPods := 1000
	numClusters := 100
	numSamples := 1 * 60 * 30

	podID := 0
	clusterID := 0
	series := make([]labels.Labels, numSeries)
	for i := 0; i < numSeries; i++ {
		series[i] = labels.FromStrings(
			"job", "nginx",
			"__name__", "http_requests_total",
			"cluster", "k8s-"+strconv.Itoa(clusterID),
			"pod", "nginx-"+strconv.Itoa(podID),
			"series_id", strconv.Itoa(i),
		)
		podID = (podID + 1) % numPods
		clusterID = (clusterID + 1) % numClusters
	}

	now := time.Now()
	from := timestamp.FromTime(now)
	to := timestamp.FromTime(now.Add(2 * time.Hour))
	_, err := e2eutil.CreateBlock(
		ctx,
		"./",
		series,
		numSamples,
		from,
		to,
		nil,
		0,
		metadata.NoneFunc,
	)
	if err != nil {
		log.Fatal(err)
	}
}
