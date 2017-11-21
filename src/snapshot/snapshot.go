/*
 * go-jepsen
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package snapshot

import (
	"xcommon"
	"xworker"
)

type Snapshot struct {
	query    xworker.Handler
	update   xworker.Handler
	qworkers []xworker.Worker
	uworkers []xworker.Worker
}

func NewSnapshot(conf *xcommon.Conf) *Snapshot {
	// updates thread.
	conf.UThreads = 1
	// Query thread.
	conf.QThreads = 16

	qworkers := xworker.CreateWorkers(conf, conf.QThreads)
	uworkers := xworker.CreateWorkers(conf, conf.UThreads)
	return &Snapshot{
		qworkers: qworkers,
		uworkers: uworkers,
		query:    NewQuery(conf, qworkers),
		update:   NewUpdate(conf, uworkers),
	}
}

func (s *Snapshot) Run() {
	s.query.Run()
	s.update.Run()
}

func (s *Snapshot) Stop() {
	s.query.Stop()
	s.update.Stop()
}

func (s *Snapshot) Query() xworker.Handler {
	return s.query
}

func (s *Snapshot) Update() xworker.Handler {
	return s.update
}

func (s *Snapshot) Workers() []xworker.Worker {
	return append(s.qworkers, s.uworkers...)
}
