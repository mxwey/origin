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

func (g *GetVariablesNode) Exec() error {
	port := g.gr.variables[g.varName]
	if port == nil {
		return fmt.Errorf("variable %s not found,node name %s", g.varName, g.nodeName)
	}

	if !g.SetOutPort(0, port) {
		return fmt.Errorf("set out port failed,node name %s", g.nodeName)
	}

	return nil
}

func (g *SetVariablesNode) GetName() string {
	return g.nodeName
}

func (g *SetVariablesNode) Exec() error {
	port := g.GetInPort(0)
	if port == nil {
		return fmt.Errorf("get in port failed,node name %s", g.nodeName)
	}

	g.gr.variables[g.varName] = port
	if !g.SetOutPort(0, port) {
		return fmt.Errorf("set out port failed,node name %s", g.nodeName)
	}

	return nil
}
