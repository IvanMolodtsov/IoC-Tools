package command

import "github.com/IvanMolodtsov/IoC-Tools/common"

type RevertCommand struct {
	up   common.ICommand
	down common.ICommand
}

func (tc *RevertCommand) Invoke() error {
	return tc.up.Invoke()
}
