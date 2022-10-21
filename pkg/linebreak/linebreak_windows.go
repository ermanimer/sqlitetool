//go:build windows

package linebreak

const (
	LineBreak       = "\r\n"
	DoubleLineBreak = LineBreak + LineBreak
)
