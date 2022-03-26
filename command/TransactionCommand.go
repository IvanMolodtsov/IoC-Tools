package command

import (
	"fmt"

	"github.com/IvanMolodtsov/IoC-Tools/common"
)

type TransactionCommand struct {
	revert common.MacroCommand
	cmds   []common.ICommand
}

func (mt *TransactionCommand) Invoke() error {
	for _, cmd := range mt.cmds {
		err := cmd.Invoke()
		if err != nil {
			revertError := mt.revert.Invoke()
			return TransactionError{cause: err, revertError: revertError}
		} else {
			if cmd, ok := cmd.(*RevertCommand); ok {
				mt.revert.Append(cmd.down)
			}
		}
	}
	return nil
}

func (mt *TransactionCommand) Append(cmd common.ICommand) {
	mt.cmds = append(mt.cmds, cmd)
}

type TransactionError struct {
	cause       error
	revertError error
}

func (e TransactionError) Error() string {
	var msg string
	if e.revertError != nil {
		msg = fmt.Sprintf("transaction failed with %e. revert failed with %e", e.cause, e.revertError)
	} else {
		msg = fmt.Sprintf("transaction failed with %e.", e.cause)
	}
	return msg
}
