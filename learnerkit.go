package learnerkit

import (
	. "gorgonia.org/gorgonia"
)

type LearnerKitModel struct {
	model_kit *[]*LearnerKit
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

func Initialize_LearnerKitModel(kit ...*LearnerKit) LearnerKitModel {
	learner_model := LearnerKitModel{
		model_kit: &kit,
	}
	return learner_model
}
