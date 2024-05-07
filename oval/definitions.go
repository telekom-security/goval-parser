package oval

func (d *Definitions) init() {
	d.lookupTable = make(map[string]int)

	for i, def := range d.Definitions {
		d.lookupTable[def.ID] = i
	}
}

// Lookup searches for a definition by its id and returns its index. If the id cannot be found return -1
func (d *Definitions) Lookup(ref string) int {
	d.once.Do(d.init)
	if i, ok := d.lookupTable[ref]; ok {
		return i
	} else {
		return -1
	}
}
