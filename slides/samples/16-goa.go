package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("cellar", func() {
	Description("The wine cellar API")
	Host("cellar.goa.design")
	Scheme("https")
	BasePath("/api")
})

var _ = Resource("Bottle", func() {
	Action("show", func() {
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer, "Bottle ID")
		})
		Response(OK, Bottle)
	})
})
