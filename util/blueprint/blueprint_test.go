package blueprint

import (
	"fmt"
	"testing"
)

type Entrance_IntParam struct {
	BaseExecNode
}

func (em *Entrance_IntParam) GetName() string {
	return "Entrance_IntParam"
}

func (em *Entrance_IntParam) Exec() (int, error) {
	return 0, nil
}

type Output struct {
	BaseExecNode
}

func (em *Output) GetName() string {
	return "Output"
}
func (em *Output) Exec() (int, error) {
	val, ok := em.GetInPortInt(0)
	if !ok {
		return 0, fmt.Errorf("Output Exec inParam not found")
	}

	fmt.Printf("Output Exec inParam %d", val)
	return 0, nil
}

func OnRegister(bm *ExecPool) error {
	bm.Register(&Entrance_IntParam{})
	bm.Register(&Output{})
	return nil
}

const (
	EntranceID_IntParam = 3
)

func TestExecMgr(t *testing.T) {
	//
	var bp Blueprint
	err := bp.Init("./json/", "./vgf/", OnRegister)
	if err != nil {
		t.Fatalf("init failed,err:%v", err)
	}

	graph := bp.Create("test1")

	err = graph.Do(EntranceID_IntParam, 1, 2, 3)
	if err != nil {
		t.Fatalf("do failed,err:%v", err)
	}

	graph.Release()
}
