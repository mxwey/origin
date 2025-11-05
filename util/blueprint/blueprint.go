package blueprint

import (
	"fmt"
	"sync/atomic"
)

type Blueprint struct {
	execNodes []IExecNode // 注册的定义执行结点

	execPool  ExecPool
	graphPool GraphPool

	blueprintModule IBlueprintModule
	mapGraph       map[int64]IGraph
	seedID int64
	cancelTimer func(*uint64)bool
}

func (bm *Blueprint) RegExecNode(execNode IExecNode) {
	bm.execNodes = append(bm.execNodes, execNode)
}

func (bm *Blueprint) regSysNode(){
	bm.RegExecNode(&AddInt{})
	bm.RegExecNode(&SubInt{})
	bm.RegExecNode(&MulInt{})
	bm.RegExecNode(&DivInt{})
	bm.RegExecNode(&ModInt{})
	bm.RegExecNode(&RandNumber{})

	bm.RegExecNode(&Entrance_ArrayParam{})
	bm.RegExecNode(&Entrance_IntParam{})
	bm.RegExecNode(&Entrance_Timer{})
	bm.RegExecNode(&DebugOutput{})
	bm.RegExecNode(&Sequence{})
	bm.RegExecNode(&Foreach{})
	bm.RegExecNode(&ForeachIntArray{})

	bm.RegExecNode(&GetArrayInt{})
	bm.RegExecNode(&GetArrayString{})
	bm.RegExecNode(&GetArrayLen{})
	bm.RegExecNode(&CreateIntArray{})
	bm.RegExecNode(&CreateStringArray{})
	bm.RegExecNode(&AppendIntegerToArray{})
	bm.RegExecNode(&AppendStringToArray{})

	bm.RegExecNode(&BoolIf{})
	bm.RegExecNode(&GreaterThanInteger{})
	bm.RegExecNode(&LessThanInteger{})
	bm.RegExecNode(&EqualInteger{})
	bm.RegExecNode(&RangeCompare{})
	bm.RegExecNode(&Probability{})
	bm.RegExecNode(&CreateTimer{})
}

func (bm *Blueprint) Init(execDefFilePath string, graphFilePath string, blueprintModule IBlueprintModule,cancelTimer func(*uint64)bool) error {
	bm.regSysNode()
	err := bm.execPool.Load(execDefFilePath)
	if err != nil {
		return err
	}

	for _, e := range bm.execNodes {
		if !bm.execPool.Register(e) {
			return fmt.Errorf("register exec failed,exec:%s", e.GetName())
		}
	}

	err = bm.graphPool.Load(&bm.execPool, graphFilePath, blueprintModule)
	if err != nil {
		return err
	}

	bm.cancelTimer = cancelTimer
	bm.blueprintModule = blueprintModule
	bm.mapGraph = make(map[int64]IGraph,128)
	return nil
}

func (bm *Blueprint) Create(graphName string) int64 {
	if graphName == "" {
		return 0
	}

	graphID := atomic.AddInt64(&bm.seedID, 1)
	gr := bm.graphPool.Create(graphName, graphID)
	if gr == nil {
		return 0
	}

	bm.mapGraph[graphID] = gr
	return graphID
}

func (bm *Blueprint) TriggerEvent(graphID int64, eventID int64, args ...any) error{
	graph := bm.mapGraph[graphID]
	if graph == nil {
		return fmt.Errorf("can not find graph:%d", graphID)
	}

	_,err:= graph.Do(eventID, args...)
	return err
}

func (bm *Blueprint) Do(graphID int64, entranceID int64, args ...any) (Port_Array,error){
	graph := bm.mapGraph[graphID]
	if graph == nil {
		return nil,fmt.Errorf("can not find graph:%d", graphID)
	}

	return graph.Do(entranceID, args...)
}

func (bm *Blueprint) ReleaseGraph(graphID int64) {
	defer delete(bm.mapGraph, graphID)
	graph := bm.mapGraph[graphID]
	if graph == nil {
		return
	}

	graph.Release()
}

func (bm *Blueprint) CancelTimerId(graphID int64, timerId *uint64) bool{
	tId := *timerId
	bm.cancelTimer(timerId)

	graph := bm.mapGraph[graphID]
	if graph == nil {
		return false
	}

	gr,ok := graph.(*Graph)
	if !ok {
		return false
	}

	delete(gr.mapTimerID, tId)
	return true
}