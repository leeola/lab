package component

import "github.com/leeola/lab/tui"

type Text string

func (c Text) Init(t tui.Tui) {}

func (c Text) Render(r tui.Region) error {
	return r.HStr(0, 0, string(c))
}
