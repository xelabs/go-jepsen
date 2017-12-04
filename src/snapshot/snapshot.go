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

// Snapshot tuple.
type Snapshot struct {
	query    xworker.Handler
	update   xworker.Handler
	qworkers []xworker.Worker
	uworkers []xworker.Worker
}

// NewSnapshot creates a new snapshot.
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

// Run used to start the worker.
func (s *Snapshot) Run() {
	s.query.Run()
	s.update.Run()
}

// Stop used to stop the worker.
func (s *Snapshot) Stop() {
	s.query.Stop()
	s.update.Stop()
}

// Query returns the query handler.
func (s *Snapshot) Query() xworker.Handler {
	return s.query
}

// Update returns the update handler.
func (s *Snapshot) Update() xworker.Handler {
	return s.update
}

// Workers returns all the worker handlers.
func (s *Snapshot) Workers() []xworker.Worker {
	return append(s.qworkers, s.uworkers...)
}
