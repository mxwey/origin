package blueprint

import (
	"fmt"
	"github.com/duanhf2012/origin/v2/log"
)

// 系统入口ID定义，1000以内
const (
	EntranceID_ArrayParam = 2
	EntranceID_IntParam   = 1
)

func init() {
	RegExecNode(&Entrance_ArrayParam{})
	RegExecNode(&Entrance_IntParam{})
	RegExecNode(&Output{})
	RegExecNode(&Sequence{})
	RegExecNode(&Foreach{})
	RegExecNode(&GetArrayInt{})
	RegExecNode(&GetArrayString{})
	RegExecNode(&GetArrayLen{})
}

type Entrance_ArrayParam struct {
	BaseExecNode
}

func (em *Entrance_ArrayParam) GetName() string {
	return "Entrance_ArrayParam"
}

func (em *Entrance_ArrayParam) Exec() (int, error) {
	return 0, nil
}

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
	val, ok := em.GetInPortInt(1)
	if !ok {
		return 0, fmt.Errorf("Output Exec inParam not found")
	}

	fmt.Printf("Output Exec inParam %d\n", val)
	return 0, nil
}

type Sequence struct {
	BaseExecNode
}

func (em *Sequence) GetName() string {
	return "Sequence"
}

func (em *Sequence) Exec() (int, error) {
	for i := range em.outPort {
		if !em.outPort[i].IsPortExec() {
			break
		}

		err := em.DoNext(i)
		if err != nil {
			return -1, err
		}
	}

	return -1, nil
}

type Foreach struct {
	BaseExecNode
}

func (em *Foreach) GetName() string {
	return "Foreach"
}

func (em *Foreach) Exec() (int, error) {
	startIndex, ok := em.ExecContext.InputPorts[1].GetInt()
	if !ok {
		return 0, fmt.Errorf("foreach Exec inParam not found")
	}
	endIndex, ok := em.ExecContext.InputPorts[2].GetInt()
	if !ok {
		return 0, fmt.Errorf("foreach Exec inParam not found")
	}

	for i := startIndex; i < endIndex; i++ {
		em.ExecContext.OutputPorts[2].SetInt(i)
		err := em.DoNext(0)
		if err != nil {
			return -1, err
		}
	}

	err := em.DoNext(1)
	if err != nil {
		return -1, err
	}

	return -1, nil
}

type GetArrayInt struct {
	BaseExecNode
}

func (em *GetArrayInt) GetName() string {
	return "GetArrayInt"
}

func (em *GetArrayInt) Exec() (int, error) {
	inPort := em.GetInPort(0)
	if inPort == nil {
		return -1, fmt.Errorf("GetArrayInt inParam not found")
	}
	outPort := em.GetOutPort(0)
	if outPort == nil {
		return -1, fmt.Errorf("GetArrayInt outParam not found")
	}

	arrIndexPort := em.GetInPort(1)
	if arrIndexPort == nil {
		return -1, fmt.Errorf("GetArrayInt arrIndexParam not found")
	}
	arrIndex, ok := arrIndexPort.GetInt()
	if !ok {
		return -1, fmt.Errorf("GetArrayInt arrIndexParam not found")
	}

	if arrIndex < 0 || arrIndex >= inPort.GetArrayLen() {
		return -1, fmt.Errorf("GetArrayInt arrIndexParam out of range,index %d", arrIndex)
	}

	val, ok := inPort.GetArrayValInt(int(arrIndex))
	if !ok {
		log.Errorf("GetArrayValInt failed, idx:%d", arrIndex)
		return -1, fmt.Errorf("GetArrayInt inParam not found")
	}

	outPort.SetInt(val)
	return -1, nil
}

type GetArrayString struct {
	BaseExecNode
}

func (em *GetArrayString) GetName() string {
	return "GetArrayString"
}

func (em *GetArrayString) Exec() (int, error) {
	inPort := em.GetInPort(0)
	if inPort == nil {
		return -1, fmt.Errorf("GetArrayInt inParam not found")
	}
	outPort := em.GetOutPort(0)
	if outPort == nil {
		return -1, fmt.Errorf("GetArrayInt outParam not found")
	}

	arrIndexPort := em.GetInPort(1)
	if arrIndexPort == nil {
		return -1, fmt.Errorf("GetArrayInt arrIndexParam not found")
	}
	arrIndex, ok := arrIndexPort.GetInt()
	if !ok {
		return -1, fmt.Errorf("GetArrayInt arrIndexParam not found")
	}

	if arrIndex < 0 || arrIndex >= inPort.GetArrayLen() {
		return -1, fmt.Errorf("GetArrayInt arrIndexParam out of range,index %d", arrIndex)
	}

	val, ok := inPort.GetArrayValStr(int(arrIndex))
	if !ok {
		log.Errorf("GetArrayValStr failed, idx:%d", arrIndex)
		return -1, fmt.Errorf("GetArrayInt inParam not found")
	}

	outPort.SetStr(val)
	return -1, nil
}

type GetArrayLen struct {
	BaseExecNode
}

func (em *GetArrayLen) GetName() string {
	return "GetArrayLen"
}

func (em *GetArrayLen) Exec() (int, error) {
	inPort := em.GetInPort(0)
	if inPort == nil {
		return -1, fmt.Errorf("GetArrayInt inParam not found")
	}
	outPort := em.GetOutPort(0)
	if outPort == nil {
		return -1, fmt.Errorf("GetArrayInt outParam not found")
	}

	outPort.SetInt(inPort.GetArrayLen())
	return -1, nil
}
