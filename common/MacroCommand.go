package common

type MacroCommand interface {
	ICommand
	Appendable[ICommand]
}
