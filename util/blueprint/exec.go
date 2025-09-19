package blueprint

import "fmt"

type IBaseExecNode interface {
	initExecNode(gr *graph, nodeId string, variableName string, nodeName string) error
}

type IBaseExec interface {
	GetName() string
	SetExec(exec IExec)
	IsInPortExec(index int) bool
	IsOutPortExec(index int) bool
	GetInPortCount() int
	CloneInOutPort() ([]IPort, []IPort)

	GetInPort(index int) IPort
	GetOutPort(index int) IPort
}

type IExec interface {
	GetName() string
	Exec() error
}

type IExecData interface {
}

type BaseExec struct {
	Name        string
	Title       string
	Package     string
	Description string

	InPort  []IPort
	OutPort []IPort
	IExec
}

type InputConfig struct {
	Name      string `json:"name"`
	PortType  string `json:"type"`
	DataType  string `json:"data_type"`
	HasInput  bool   `json:"has_input"`
	PinWidget string `json:"pin_widget"`
}

type OutInputConfig struct {
	Name     string `json:"name"`
	PortType string `json:"type"`
	DataType string `json:"data_type"`
	HasInput bool   `json:"has_input"`
}

type BaseExecConfig struct {
	Name        string           `json:"name"`
	Title       string           `json:"title"`
	Package     string           `json:"package"`
	Description string           `json:"description"`
	IsPure      bool             `json:"is_pure"`
	Inputs      []InputConfig    `json:"inputs"`
	Outputs     []OutInputConfig `json:"outputs"`
}

func (em *BaseExec) AppendInPort(port ...IPort) {
	em.InPort = append(em.InPort, port...)
}

func (em *BaseExec) AppendOutPort(port ...IPort) {
	em.OutPort = append(em.OutPort, port...)
}

func (em *BaseExec) GetName() string {
	return em.Name
}

func (em *BaseExec) SetExec(exec IExec) {
	em.IExec = exec
}

func (em *BaseExec) CloneInOutPort() ([]IPort, []IPort) {
	inPorts := make([]IPort, 0, 2)
	for _, port := range em.InPort {
		if port.IsPortExec() {
			continue
		}

		inPorts = append(inPorts, port.Clone())
	}
	outPorts := make([]IPort, 0, 2)

	for _, port := range em.OutPort {
		if port.IsPortExec() {
			continue
		}
		outPorts = append(outPorts, port.Clone())
	}

	return inPorts, outPorts
}

func (em *BaseExec) IsInPortExec(index int) bool {
	if index >= len(em.InPort) || index < 0 {
		return false
	}

	return em.InPort[index].IsPortExec()
}
func (em *BaseExec) IsOutPortExec(index int) bool {
	if index >= len(em.OutPort) || index < 0 {
		return false
	}

	return em.OutPort[index].IsPortExec()
}

func (em *BaseExec) GetInPortCount() int {
	return len(em.InPort)
}

func (em *BaseExec) GetInPort(index int) IPort {
	if index >= len(em.InPort) || index < 0 {
		return nil
	}
	return em.InPort[index]
}

func (em *BaseExec) GetOutPort(index int) IPort {
	if index >= len(em.OutPort) || index < 0 {
		return nil
	}
	return em.OutPort[index]
}

type BaseExecNode struct {
	*ExecContext
	gr           *graph
	variableName string
	nodeName     string
}

func (en *BaseExecNode) initExecNode(gr *graph, nodeId string, variableName string, nodeName string) error {
	ctx, ok := gr.context[nodeId]
	if !ok {
		return fmt.Errorf("node %s not found", nodeId)
	}
	en.ExecContext = ctx
	en.gr = gr
	en.variableName = variableName
	en.nodeName = nodeName
	return nil
}

func (en *BaseExecNode) GetInPort(index int) IPort {
	if index >= len(en.InputPorts) || index < 0 {
		return nil
	}
	return en.InputPorts[index]
}

func (en *BaseExecNode) GetOutPort(index int) IPort {
	if index >= len(en.OutputPorts) || index < 0 {
		return nil
	}
	return en.OutputPorts[index]
}

func (en *BaseExecNode) SetOutPort(index int, val IPort) bool {
	if index >= len(en.OutputPorts) || index < 0 {
		return false
	}
	en.OutputPorts[index] = val
	return true
}

func (en *BaseExecNode) GetInPortInt(index int) (Port_Int, bool) {
	port := en.GetInPort(index)
	if port == nil {
		return 0, false
	}
	return port.GetInt()
}

func (en *BaseExecNode) GetInPortFloat(index int) (Port_Float, bool) {
	port := en.GetInPort(index)
	if port == nil {
		return 0, false
	}
	return port.GetFloat()
}

func (en *BaseExecNode) GetInPortStr(index int) (Port_Str, bool) {
	port := en.GetInPort(index)
	if port == nil {
		return "", false
	}
	return port.GetStr()
}

func (en *BaseExecNode) GetInPortArrayValInt(index int, idx int) (Port_Int, bool) {
	port := en.GetInPort(index)
	if port == nil {
		return 0, false
	}
	return port.GetArrayValInt(idx)
}

func (en *BaseExecNode) GetInPortArrayValStr(idx int) (Port_Str, bool) {
	port := en.GetInPort(idx)
	if port == nil {
		return "", false
	}
	return port.GetArrayValStr(idx)
}

func (en *BaseExecNode) GetInPortBool(index int) (Port_Bool, bool) {
	port := en.GetInPort(index)
	if port == nil {
		return false, false
	}
	return port.GetBool()
}

func (en *BaseExecNode) GetOutPortInt(index int) (Port_Int, bool) {
	port := en.GetOutPort(index)
	if port == nil {
		return 0, false
	}
	return port.GetInt()
}

func (en *BaseExecNode) GetOutPortFloat(index int) (Port_Float, bool) {
	port := en.GetOutPort(index)
	if port == nil {
		return 0, false
	}
	return port.GetFloat()
}

func (en *BaseExecNode) GetOutPortStr(index int) (Port_Str, bool) {
	port := en.GetOutPort(index)
	if port == nil {
		return "", false
	}
	return port.GetStr()
}

func (en *BaseExecNode) GetOutPortArrayValInt(index int, idx int) (Port_Int, bool) {
	port := en.GetOutPort(index)
	if port == nil {
		return 0, false
	}
	return port.GetArrayValInt(idx)
}

func (en *BaseExecNode) GetOutPortArrayValStr(index int, idx int) (Port_Str, bool) {
	port := en.GetOutPort(index)
	if port == nil {
		return "", false
	}
	return port.GetArrayValStr(idx)
}

func (en *BaseExecNode) GetOutPortBool(index int) (Port_Bool, bool) {
	port := en.GetInPort(index)
	if port == nil {
		return false, false
	}
	return port.GetBool()
}

func (en *BaseExecNode) SetInPortInt(index int, val Port_Int) bool {
	port := en.GetInPort(index)
	if port == nil {
		return false
	}
	return port.SetInt(val)
}

func (en *BaseExecNode) SetInPortFloat(index int, val Port_Float) bool {
	port := en.GetInPort(index)
	if port == nil {
		return false
	}
	return port.SetFloat(val)
}

func (en *BaseExecNode) SetInPortStr(index int, val Port_Str) bool {
	port := en.GetInPort(index)
	if port == nil {
		return false
	}
	return port.SetStr(val)
}

func (en *BaseExecNode) SetInBool(index int, val Port_Bool) bool {
	port := en.GetInPort(index)
	if port == nil {
		return false
	}
	return port.SetBool(val)
}

func (en *BaseExecNode) SetInPortArrayValInt(index int, idx int, val Port_Int) bool {
	port := en.GetInPort(index)
	if port == nil {
		return false
	}
	return port.SetArrayValInt(idx, val)
}

func (en *BaseExecNode) SetInPortArrayValStr(index int, idx int, val Port_Str) bool {
	port := en.GetInPort(index)
	if port == nil {
		return false
	}
	return port.SetArrayValStr(idx, val)
}

func (en *BaseExecNode) AppendInPortArrayValInt(index int, val Port_Int) bool {
	port := en.GetInPort(index)
	if port == nil {
		return false
	}
	return port.AppendArrayValInt(val)
}

func (en *BaseExecNode) AppendInPortArrayValStr(index int, val Port_Str) bool {
	port := en.GetInPort(index)
	if port == nil {
		return false
	}
	return port.AppendArrayValStr(val)
}

func (en *BaseExecNode) GetInPortArrayLen(index int) int {
	port := en.GetInPort(index)
	if port == nil {
		return 0
	}
	return port.GetArrayLen()
}

func (en *BaseExecNode) SetOutPortInt(index int, val Port_Int) bool {
	port := en.GetOutPort(index)
	if port == nil {
		return false
	}
	return port.SetInt(val)
}

func (en *BaseExecNode) SetOutPortFloat(index int, val Port_Float) bool {
	port := en.GetOutPort(index)
	if port == nil {
		return false
	}
	return port.SetFloat(val)
}

func (en *BaseExecNode) SetOutPortStr(index int, val Port_Str) bool {
	port := en.GetOutPort(index)
	if port == nil {
		return false
	}
	return port.SetStr(val)
}

func (en *BaseExecNode) SetOutPortBool(index int, val Port_Bool) bool {
	port := en.GetOutPort(index)
	if port == nil {
		return false
	}
	return port.SetBool(val)
}

func (en *BaseExecNode) SetOutPortArrayValInt(index int, idx int, val Port_Int) bool {
	port := en.GetOutPort(index)
	if port == nil {
		return false
	}
	return port.SetArrayValInt(idx, val)
}

func (en *BaseExecNode) SetOutPortArrayValStr(index int, idx int, val Port_Str) bool {
	port := en.GetOutPort(index)
	if port == nil {
		return false
	}
	return port.SetArrayValStr(idx, val)
}

func (en *BaseExecNode) AppendOutPortArrayValInt(index int, val Port_Int) bool {
	port := en.GetOutPort(index)
	if port == nil {
		return false
	}
	return port.AppendArrayValInt(val)
}

func (en *BaseExecNode) AppendOutPortArrayValStr(index int, val Port_Str) bool {
	port := en.GetOutPort(index)
	if port == nil {
		return false
	}
	return port.AppendArrayValStr(val)
}

func (en *BaseExecNode) GetOutPortArrayLen(index int) int {
	port := en.GetOutPort(index)
	if port == nil {
		return 0
	}
	return port.GetArrayLen()
}
