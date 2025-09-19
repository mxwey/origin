package blueprint

import "fmt"

type prePortNode struct {
	node         *execNode // 上个结点
	outPortIndex int       // 对应上一个结点的OutPort索引
}

type execNode struct {
	Id       string
	baseExec IBaseExec

	nextNode []*execNode
	nextIdx  int

	preInPort          []*prePortNode // Port的上一个结点
	inPortDefaultValue map[string]any

	variableName string // 如果是变量，则有变量名
}

func (en *execNode) GetInPortDefaultValue(index int) any {
	key := fmt.Sprintf("%d", index)
	v, ok := en.inPortDefaultValue[key]
	if !ok {
		return nil
	}
	return v
}

func (en *execNode) Next() *execNode {
	if en.nextIdx >= len(en.nextNode) {
		return nil
	}

	return en.nextNode[en.nextIdx]
}

func (en *execNode) exec(gr *graph) error {
	e, ok := en.baseExec.(IExec)
	if !ok {
		return fmt.Errorf("exec node %s not exec", en.baseExec.GetName())
	}

	node, ok := en.baseExec.(IBaseExecNode)
	if !ok {
		return fmt.Errorf("exec node %s not exec", en.baseExec.GetName())
	}

	if err := node.initExecNode(gr, en.Id, en.variableName, en.baseExec.GetName()); err != nil {
		return err
	}

	return e.Exec()
}

func (en *execNode) doSetInPort(gr *graph, index int, inPort IPort) error {
	// 找到当前Node的InPort的index的前一个结点
	preNode := en.preInPort[index]
	// 如果前一个结点为空，则填充默认值
	if preNode == nil {
		err := inPort.setAnyVale(en.GetInPortDefaultValue(index))
		if err != nil {
			return err
		}
		return nil
	}

	// 判断上一个结点是否已经执行过
	if _, ok := gr.context[preNode.node.Id]; ok {
		outPort := gr.GetNodeOutPortValue(preNode.node.Id, preNode.outPortIndex)
		if outPort == nil {
			return fmt.Errorf("pre node %s out port index %d not found", preNode.node.Id, preNode.outPortIndex)
		}

		inPort.SetValue(outPort)
		return nil
	}

	// 如果前一个结点没有执行过，则递归执行前一个结点
	return preNode.node.Do(gr)
}

func (en *execNode) Do(gr *graph) error {
	// 重新初始化上下文
	inPorts, outPorts := en.baseExec.CloneInOutPort()
	gr.context[en.Id] = &ExecContext{
		InputPorts:  inPorts,
		OutputPorts: outPorts,
	}

	// 处理InPort结点值
	var err error
	for index := range inPorts {
		if en.baseExec.IsInPortExec(index) {
			continue
		}

		err = en.doSetInPort(gr, index, inPorts[index])
		if err != nil {
			return err
		}
	}

	// 设置执行器相关的上下文信息
	// 如果是变量设置变量名
	// 执行本结点
	if err = en.exec(gr); err != nil {
		return err
	}

	for _, nextNode := range en.nextNode {
		err = nextNode.Do(gr)
		if err != nil {
			return err
		}
	}

	return nil
}
