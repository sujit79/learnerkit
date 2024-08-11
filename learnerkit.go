package learnerkit

import (
	. "gorgonia.org/gorgonia"
)

type LearnerKitModel struct {
	model_kit []LearnerKit
}

type LearnerKit struct {
	node    LearnerKitNode
	graph   LearnerKitGraph
	machine LearnerKitMachine
	op      LearnerKitOps
}

type LearnerKitNode struct {
	node *[]*Node
}

type LearnerKitOps struct {
	ops Op
}

type LearnerKitGraph struct {
	g *ExprGraph
}

type LearnerKitMachine struct {
	machine VM
}

func ApplyOps_LearnerKit(op Op, children ...*Node) (*Node, error) {
	return ApplyOp(op, children...)
}

func ApplyOps_With_Name_LearnerKit(op Op, name string, children ...*Node) (*Node, error) {
	return ApplyOpWithName(op, name, children...)
}

func ApplyOps_LearnerKitModel(model LearnerKitModel) ([][]Node, error) {
	results := make([][]Node, len(model.model_kit))
	var e error
	for i := 0; i < len(model.model_kit); i++ {
		result := make([]Node, len(*model.model_kit[i].node.node))
		for j := 0; j < len(*model.model_kit[i].node.node); j++ {
			value, e := ApplyOps_LearnerKit(model.model_kit[i].op.ops, *model.model_kit[j].node.node...)
			if e != nil {
				break
			}
			result = append(result, *value)
		}
		results[i] = result
	}
	return results, e
}

func ApplyOps_WithName_LearnerKitModel_WithName(model LearnerKitModel, name string) ([][]Node, error) {
	results := make([][]Node, len(model.model_kit))
	var e error
	for i := 0; i < len(model.model_kit); i++ {
		result := make([]Node, len(*model.model_kit[i].node.node))
		for j := 0; j < len(*model.model_kit[i].node.node); j++ {
			value, e := ApplyOps_With_Name_LearnerKit(model.model_kit[i].op.ops, name, *model.model_kit[j].node.node...)
			if e != nil {
				break
			}
			result = append(result, *value)
		}
		results[i] = result
	}
	return results, e
}

func Initialize_LearnerKit(graph *ExprGraph, machine VM, op Op, node_v ...*Node) LearnerKit {
	learnerkitnode := LearnerKitNode{
		node: &node_v,
	}
	learnerkitgraph := LearnerKitGraph{
		g: graph,
	}
	learnerkitmachine := LearnerKitMachine{
		machine: machine,
	}
	learnerkitops := LearnerKitOps{
		ops: op,
	}
	learnerkit := LearnerKit{
		node:    learnerkitnode,
		graph:   learnerkitgraph,
		machine: learnerkitmachine,
		op:      learnerkitops,
	}
	return learnerkit
}

func Initialize_LearnerKitModel(kit []LearnerKit) LearnerKitModel {
	learner_model := LearnerKitModel{
		model_kit: kit,
	}
	return learner_model
}
