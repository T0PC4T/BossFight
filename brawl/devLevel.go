package brawl

// newDevelopmentLevel Creates a new level in this case its a development level
func newDevelopmentLevel() *level {
	l := &level{name: "DevLevel", allElements: make([]element, 0, 50)}
	for tx, col := range l.mapLayout {
		for ty := range col {
			if ty == 20 {
				l.mapLayout[tx][ty] = newTile(tx, ty)
			}
		}
	}
	return l
}
