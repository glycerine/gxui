package basic

import (
	"github.com/glycerine/gxui"
	"github.com/glycerine/gxui/mixins"
)

func CreateTableLayout(theme *Theme) gxui.TableLayout {
	l := &mixins.TableLayout{}
	l.Init(l, theme)
	return l
}
