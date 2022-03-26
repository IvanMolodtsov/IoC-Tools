package command

import "github.com/IvanMolodtsov/IoC-Tools/common"

type BasicMacroCommand struct {
	cmds []common.ICommand
}

func (mc *BasicMacroCommand) Invoke() error {
	for _, cmd := range mc.cmds {
		err := cmd.Invoke()
		if err != nil {
			return err
		}
	}
	return nil
}

func (mc *BasicMacroCommand) Append(cmd common.ICommand) {
	mc.cmds = append(mc.cmds, cmd)
}
