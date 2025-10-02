package blueprint

import "fmt"

func init() {
	RegExecNode(&AddInt{})
	RegExecNode(&SubInt{})
	RegExecNode(&MulInt{})
	RegExecNode(&DivInt{})
	RegExecNode(&ModInt{})
}

type AddInt struct {
	BaseExecNode
}

func (em *AddInt) GetName() string {
	return "AddInt"
}

func (em *AddInt) Exec() (int, error) {
	inPortA := em.GetInPort(0)
	if inPortA == nil {
		return -1, fmt.Errorf("AddInt inParam 1 not found")
	}

	inPortB := em.GetInPort(1)
	if inPortB == nil {
		return -1, fmt.Errorf("AddInt inParam 2 not found")
	}

	outPortRet := em.GetOutPort(0)
	if outPortRet == nil {
		return -1, fmt.Errorf("AddInt outParam not found")
	}

	inA, ok := inPortA.GetInt()
	if !ok {
		return -1, fmt.Errorf("AddInt inParam 1 not found")
	}
	inB, ok := inPortB.GetInt()
	if !ok {
		return -1, fmt.Errorf("AddInt inParam 2 not found")
	}
	ret := inA + inB
	outPortRet.SetInt(ret)

	return -1, nil
}

type SubInt struct {
	BaseExecNode
}

func (em *SubInt) GetName() string {
	return "SubInt"
}

func (em *SubInt) Exec() (int, error) {
	inPortA := em.GetInPort(0)
	if inPortA == nil {
		return -1, fmt.Errorf("AddInt inParam 1 not found")
	}

	inPortB := em.GetInPort(1)
	if inPortB == nil {
		return -1, fmt.Errorf("AddInt inParam 2 not found")
	}

	inPortAbs := em.GetInPort(2)
	if inPortAbs == nil {
		return -1, fmt.Errorf("AddInt inParam 2 not found")
	}

	outPortRet := em.GetOutPort(0)
	if outPortRet == nil {
		return -1, fmt.Errorf("AddInt outParam not found")
	}

	inA, ok := inPortA.GetInt()
	if !ok {
		return -1, fmt.Errorf("AddInt inParam 1 not found")
	}
	inB, ok := inPortB.GetInt()
	if !ok {
		return -1, fmt.Errorf("AddInt inParam 2 not found")
	}
	isAbs, ok := inPortAbs.GetBool()
	if !ok {
		return -1, fmt.Errorf("AddInt inParam 2 not found")
	}
	ret := inA - inB
	if isAbs && ret < 0 {
		ret *= -1
	}

	outPortRet.SetInt(ret)

	return -1, nil
}

type MulInt struct {
	BaseExecNode
}

func (em *MulInt) GetName() string {
	return "MulInt"
}

func (em *MulInt) Exec() (int, error) {
	inPortA := em.GetInPort(0)
	if inPortA == nil {
		return -1, fmt.Errorf("AddInt inParam 1 not found")
	}

	inPortB := em.GetInPort(1)
	if inPortB == nil {
		return -1, fmt.Errorf("AddInt inParam 2 not found")
	}

	outPortRet := em.GetOutPort(0)
	if outPortRet == nil {
		return -1, fmt.Errorf("AddInt outParam not found")
	}

	inA, ok := inPortA.GetInt()
	if !ok {
		return -1, fmt.Errorf("AddInt inParam 1 not found")
	}
	inB, ok := inPortB.GetInt()
	if !ok {
		return -1, fmt.Errorf("AddInt inParam 2 not found")
	}

	outPortRet.SetInt(inA * inB)

	return -1, nil
}

type DivInt struct {
	BaseExecNode
}

func (em *DivInt) GetName() string {
	return "DivInt"
}

func (em *DivInt) Exec() (int, error) {
	inPortA := em.GetInPort(0)
	if inPortA == nil {
		return -1, fmt.Errorf("AddInt inParam 1 not found")
	}

	inPortB := em.GetInPort(1)
	if inPortB == nil {
		return -1, fmt.Errorf("AddInt inParam 2 not found")
	}

	inPortRound := em.GetInPort(2)
	if inPortRound == nil {
		return -1, fmt.Errorf("AddInt inParam 2 not found")
	}

	outPortRet := em.GetOutPort(0)
	if outPortRet == nil {
		return -1, fmt.Errorf("AddInt outParam not found")
	}

	inA, ok := inPortA.GetInt()
	if !ok {
		return -1, fmt.Errorf("AddInt inParam 1 not found")
	}
	inB, ok := inPortB.GetInt()
	if !ok {
		return -1, fmt.Errorf("AddInt inParam 2 not found")
	}
	isRound, ok := inPortRound.GetBool()
	if !ok {
		return -1, fmt.Errorf("AddInt inParam 2 not found")
	}

	if inB == 0 {
		return -1, fmt.Errorf("div zero error")
	}

	var ret int64
	if isRound {
		ret = (inA + inB/2) / inB
	} else {
		ret = inA / inB
	}

	outPortRet.SetInt(ret)

	return -1, nil
}

type ModInt struct {
	BaseExecNode
}

func (em *ModInt) GetName() string {
	return "ModInt"
}

func (em *ModInt) Exec() (int, error) {
	inPortA := em.GetInPort(0)
	if inPortA == nil {
		return -1, fmt.Errorf("AddInt inParam 1 not found")
	}

	inPortB := em.GetInPort(1)
	if inPortB == nil {
		return -1, fmt.Errorf("AddInt inParam 2 not found")
	}

	outPortRet := em.GetOutPort(0)
	if outPortRet == nil {
		return -1, fmt.Errorf("AddInt outParam not found")
	}

	inA, ok := inPortA.GetInt()
	if !ok {
		return -1, fmt.Errorf("AddInt inParam 1 not found")
	}
	inB, ok := inPortB.GetInt()
	if !ok {
		return -1, fmt.Errorf("AddInt inParam 2 not found")
	}
	if inB == 0 {
		return -1, fmt.Errorf("mod zero error")
	}

	outPortRet.SetInt(inA % inB)

	return -1, nil
}
