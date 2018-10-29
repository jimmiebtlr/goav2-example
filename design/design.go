package design

import . "goa.design/goa/http/design"
import . "goa.design/goa/http/dsl"
import _ "goa.design/goa/codegen/generator"

var _ = API("servicea", func() {
	Title("The service a")
})

var Info = ResultType("application/api.info", func() {
	Description("Info about service a")
	Attributes(func() {
		Attribute("id", String, "ID of the service")
		Attribute("service_name", String, "Service's name")
		Attribute("version", String, "Service's version")
	})

	Required("id")
})

var FooBar = ResultType("application/foobar", func() {
	Description("Info about service a")
	Attributes(func() {
		Attribute("id", String, "ID of the service")
	})

	Required("id")
})

var Widget = ResultType("application/widget", func() {
	Description("Info about service a")
	Attributes(func() {
		Attribute("id", String, "ID of the service")
	})

	Required("id")
})

var _ = Service("api", func() {
	Description("API of service")

	HTTP(func() {
		Path("/api")
	})

	Method("info", func() {
		Description("Show info of the service")
		Payload(Empty)
		Result(Info)
		HTTP(func() {
			GET("/info")
			Response(StatusOK)
		})
	})
})

var _ = Service("foobar", func() {
	Description("Foobar microservice")

	HTTP(func() {
		Path("/foobars")
	})

	Method("list", func() {
		Description("Show info of the service")
		Payload(Empty)
		Result(CollectionOf(FooBar))
		HTTP(func() {
			GET("")
			Response(StatusOK)
		})
	})
})

var _ = Service("widget", func() {
	Description("Widgets microservice")

	HTTP(func() {
		Path("/widgets")
	})

	Method("list", func() {
		Description("Show info of the service")
		Payload(Empty)
		Result(CollectionOf(Widget))
		HTTP(func() {
			GET("")
			Response(StatusOK)
		})
	})
})
