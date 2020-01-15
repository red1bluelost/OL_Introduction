package pageholder

import (
	"github.com/red1bluelost/OL_Introduction/page"
)

type PageHolder struct {
	Pages map[string]page.Pager
}

func (p *PageHolder) LoadPage(title string) *page.Pager {

}