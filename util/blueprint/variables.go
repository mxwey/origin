package blueprint

import "fmt"

const GetVariables = "GetVar"
const SetVariables = "SetVar"

type GetVariablesNode struct {
	BaseExecNode
	nodeName string
	varName  string
}

type SetVariablesNode struct {
	BaseExecNode
	nodeName string
	varName  string
}

func (g *GetVariablesNode) GetName() string {
	return g.nodeName
}

func (g *GetVariablesNode) Exec() (int, error) {
	port := g.gr.variables[g.varName]
	if port == nil {
		return -1, fmt.Errorf("variable %s not found,node name %s", g.varName, g.nodeName)
	}

	if !g.SetOutPort(0, port) {
		return -1, fmt.Errorf("set out port failed,node name %s", g.nodeName)
	}

	return 0, nil
}

func (g *SetVariablesNode) GetName() string {
	return g.nodeName
}

func (g *SetVariablesNode) Exec() (int, error) {
	port := g.GetInPort(0)
	if port == nil {
		return -1, fmt.Errorf("get in port failed,node name %s", g.nodeName)
	}

	g.gr.variables[g.varName] = port
	if !g.SetOutPort(0, port) {
		return -1, fmt.Errorf("set out port failed,node name %s", g.nodeName)
	}

	return 0, nil
}
