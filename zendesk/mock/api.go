package mock

import (
	"github.com/abebars/go-zendesk/zendesk"
)

var _ zendesk.API = (*Client)(nil)
