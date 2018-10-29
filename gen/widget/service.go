// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// widget service
//
// Command:
// $ goa gen github.com/jimmiebtlr/goav2-example/design

package widget

import (
	"context"

	widgetviews "github.com/jimmiebtlr/goav2-example/gen/widget/views"
)

// Widgets microservice
type Service interface {
	// Show info of the service
	List(context.Context) (res WidgetCollection, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "widget"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"list"}

// WidgetCollection is the result type of the widget service list method.
type WidgetCollection []*Widget

// Info about service a
type Widget struct {
	// ID of the service
	ID string
}

// NewWidgetCollection initializes result type WidgetCollection from viewed
// result type WidgetCollection.
func NewWidgetCollection(vres widgetviews.WidgetCollection) WidgetCollection {
	var res WidgetCollection
	switch vres.View {
	case "default", "":
		res = newWidgetCollection(vres.Projected)
	}
	return res
}

// NewViewedWidgetCollection initializes viewed result type WidgetCollection
// from result type WidgetCollection using the given view.
func NewViewedWidgetCollection(res WidgetCollection, view string) widgetviews.WidgetCollection {
	var vres widgetviews.WidgetCollection
	switch view {
	case "default", "":
		p := newWidgetCollectionView(res)
		vres = widgetviews.WidgetCollection{p, "default"}
	}
	return vres
}

// newWidgetCollection converts projected type WidgetCollection to service type
// WidgetCollection.
func newWidgetCollection(vres widgetviews.WidgetCollectionView) WidgetCollection {
	res := make(WidgetCollection, len(vres))
	for i, n := range vres {
		res[i] = newWidget(n)
	}
	return res
}

// newWidgetCollectionView projects result type WidgetCollection into projected
// type WidgetCollectionView using the "default" view.
func newWidgetCollectionView(res WidgetCollection) widgetviews.WidgetCollectionView {
	vres := make(widgetviews.WidgetCollectionView, len(res))
	for i, n := range res {
		vres[i] = newWidgetView(n)
	}
	return vres
}

// newWidget converts projected type Widget to service type Widget.
func newWidget(vres *widgetviews.WidgetView) *Widget {
	res := &Widget{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	return res
}

// newWidgetView projects result type Widget into projected type WidgetView
// using the "default" view.
func newWidgetView(res *Widget) *widgetviews.WidgetView {
	vres := &widgetviews.WidgetView{
		ID: &res.ID,
	}
	return vres
}
