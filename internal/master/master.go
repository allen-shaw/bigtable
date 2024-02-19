package master

import "github.com/allen-shaw/bigtable/internal/common"

type Master struct {
	stateMachine MasterStateMachine

	tabletManager     *TabletManager
	tabletNodeManager *TabletNodeManager
	userManager       *UserManager
	sizeScheduler     Scheduler
	loadScheduler     Scheduler

	// election

	//
	sequenceID common.Counter

	// access entry

}
