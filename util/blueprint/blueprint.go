package blueprint

type Blueprint struct {
	execPool  ExecPool
	graphPool GraphPool
}

func (bm *Blueprint) Init(execDefFilePath string, graphFilePath string) {
	bm.execPool.Load(execDefFilePath)
	bm.graphPool.Load(graphFilePath)
}
