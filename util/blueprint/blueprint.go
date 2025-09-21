package blueprint

type Blueprint struct {
	execPool  ExecPool
	graphPool GraphPool
}

func (bm *Blueprint) Init(execDefFilePath string, graphFilePath string, onRegister func(execPool *ExecPool) error) error {
	err := bm.execPool.Load(execDefFilePath)
	if err != nil {
		return err
	}

	err = onRegister(&bm.execPool)
	if err != nil {
		return err
	}

	err = bm.graphPool.Load(&bm.execPool, graphFilePath)
	if err != nil {
		return err
	}
	return nil
}
