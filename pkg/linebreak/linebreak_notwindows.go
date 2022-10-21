//go:build !windows

package linebreak

const (
	LineBreak       = "\n"
	DoubleLineBreak = LineBreak + LineBreak
)
