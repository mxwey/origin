package blueprint

import (
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

func OnRegister(bm *ExecPool) error {
	bm.Register(&Entrance_IntParam{})
	return nil
}

func TestExecMgr(t *testing.T) {
	//
	var bp Blueprint
	err := bp.Init("./json/", "./vgf/", OnRegister)
	if err != nil {
		t.Fatalf("init failed,err:%v", err)
	}
}
