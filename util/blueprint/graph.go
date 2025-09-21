package blueprint

import "fmt"

type IGraph interface {
	Do(entranceID int64) error
}

type graph struct {
	context         map[string]*ExecContext // 上下文
	entrance        map[int64]*execNode     // 入口
	variables       map[string]IPort        // 变量
	globalVariables map[string]IPort        // 全局变量
}

type nodeConfig struct {
	Id     string `json:"id"`
	Class  string `json:"class"`
	Module string `json:"module"`
	//Pos         []float64              `json:"pos"`
	PortDefault map[string]interface{} `json:"port_defaultv"`
}

type edgeConfig struct {
	EdgeID       string `json:"edge_id"`
	SourceNodeID string `json:"source_node_id"`
	DesNodeId    string `json:"des_node_id"`

	SourcePortIndex int `json:"source_port_index"`
	DesPortIndex    int `json:"des_port_index"`
}

type variablesConfig struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type graphConfig struct {
	GraphName string `json:"graph_name"`
	Time      string `json:"time"`

	Nodes     []nodeConfig      `json:"nodes"`
	Edges     []edgeConfig      `json:"edges"`
	Variables []variablesConfig `json:"variables"`
}

func (gc *graphConfig) GetVariablesByName(varName string) *variablesConfig {
	for _, varCfg := range gc.Variables {
		if varCfg.Name == varName {
			return &varCfg
		}
	}

	return nil
}

func (gc *graphConfig) GetNodeByID(nodeID string) *nodeConfig {
	for _, node := range gc.Nodes {
		if node.Id == nodeID {
			return &node
		}
	}

	return nil
}

func (gr *graph) Do(entranceID int64) error {
	entranceNode := gr.entrance[entranceID]
	if entranceNode == nil {
		return fmt.Errorf("entranceID:%d not found", entranceID)
	}

	gr.variables = map[string]IPort{}
	if gr.globalVariables == nil {
		gr.globalVariables = map[string]IPort{}
	}

	return entranceNode.Do(gr)
}

func (gr *graph) GetNodeInPortValue(nodeID string, inPortIndex int) IPort {
	if ctx, ok := gr.context[nodeID]; ok {
		if inPortIndex >= len(ctx.InputPorts) || inPortIndex < 0 {
			return nil
		}

		return ctx.InputPorts[inPortIndex]
	}
	return nil
}

func (gr *graph) GetNodeOutPortValue(nodeID string, outPortIndex int) IPort {
	if ctx, ok := gr.context[nodeID]; ok {
		if outPortIndex >= len(ctx.OutputPorts) || outPortIndex < 0 {
			return nil
		}
		return ctx.OutputPorts[outPortIndex]
	}
	return nil
}
