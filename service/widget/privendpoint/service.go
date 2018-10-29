package privendpoint

import (
	"context"
	"log"

	api "github.com/jimmiebtlr/goav2-example/gen/api"
	widget "github.com/jimmiebtlr/goav2-example/gen/widget"
)

type widgetSvc struct {
	logger *log.Logger
}

func (s *widgetSvc) List(ctx context.Context) (res []widget.Widget, err error) {
	res = []widget.Widget{
		{ID: "1"},
	}
	s.logger.Print("widget.list")
	return res, nil
}

func (s *widgetSvc) Info(ctx context.Context) (res *api.APIInfo, err error) {
	res = &api.APIInfo{
		ID: "1",
	}
	s.logger.Print("api.info")
	return res, nil
}
