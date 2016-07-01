var _ = Resource("pet", func() {
	Action("getPetById", func() {
		Description("Returns a single pet")

		Routing(GET("/:petId"))

		Params(func() {
			Param("petId", Integer, "ID of pet to return")
			Required("petId")
		})

		Response(OK, PetMedia, func() {
			Description("successful operation")
		})
		Response(BadRequest, func() {
			Description("Invalid ID supplied")
		})
		Response(NotFound, func() {
			Description("Pet not found")
		})
	})
})
