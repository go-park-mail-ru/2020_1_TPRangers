package tools

type GoPoolInterface interface {
	Schedule(task func())
}
