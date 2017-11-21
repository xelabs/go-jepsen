/*
 * benchyou
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xcmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
	"xcommon"
	"xworker"
)

type Monitor struct {
	conf    *xcommon.Conf
	workers []xworker.Worker
	ticker  *time.Ticker
	seconds uint64
}

func NewMonitor(conf *xcommon.Conf, workers []xworker.Worker) *Monitor {
	return &Monitor{
		conf:    conf,
		workers: workers,
		ticker:  time.NewTicker(time.Second),
	}
}

func (m *Monitor) Start() {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	go func() {
		oldm := &xworker.Metric{}
		for range m.ticker.C {
			m.seconds++

			newm := xworker.AllWorkersMetric(m.workers)
			wops := float64(newm.WNums - oldm.WNums)
			rops := float64(newm.QNums - oldm.QNums)
			fmt.Fprintln(w, "time   \t\t   thds\t \tw-ops\t \tr-ops\t \terror(s)\t \ttotal-ops")
			line := fmt.Sprintf("[%ds]\t\t[r:%d,u:%d]\t \t%d\t \t%d\t \t%v\t \t%v\n",
				m.seconds,
				m.conf.QThreads,
				m.conf.UThreads,
				int(wops),
				int(rops),
				newm.QErrs,
				newm.WNums+newm.QNums,
			)
			fmt.Fprintln(w, line)

			w.Flush()
			*oldm = *newm
		}
	}()
}

func (m *Monitor) Stop() {
	m.ticker.Stop()
	xworker.StopWorkers(m.workers)
}
