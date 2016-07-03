import "cellar/app"

// Show retrieves the bottle with the given id.
// Look Ma' no Bind!
func (b *BottleController) Show(ctx *app.ShowBottleContext) error {
	bottle, ok := db.GetBottle(ctx.BottleID)
	if !ok {
		return ctx.NotFound()
	}
	return ctx.OK(bottle)
}

// OMIT
