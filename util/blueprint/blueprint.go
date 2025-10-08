package blueprint

import (
	"fmt"
)

type Blueprint struct {
	execPool  ExecPool
	graphPool GraphPool

	blueprintModule IBlueprintModule
}

func (bm *Blueprint) Init(execDefFilePath string, graphFilePath string, blueprintModule IBlueprintModule) error {
	err := bm.execPool.Load(execDefFilePath)
	if err != nil {
		return err
	}

	for _, e := range execNodes {
		if !bm.execPool.Register(e) {
			return fmt.Errorf("register exec failed,exec:%s", e.GetName())
		}
	}

	err = bm.graphPool.Load(&bm.execPool, graphFilePath, blueprintModule)
	if err != nil {
		return err
	}

	bm.blueprintModule = blueprintModule
	return nil
}

func (bm *Blueprint) Create(graphName string, graphID int64) IGraph {
	return bm.graphPool.Create(graphName, graphID)
}
