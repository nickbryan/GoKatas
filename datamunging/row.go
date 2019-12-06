package datamunging

// Row represents a single row in the data.
type Row struct {
	Id       string
	Min, Max float64
}

// Spread calculates the difference between Min and Max to give the Spread.
func (r *Row) Spread() float64 {
	return r.Max - r.Min
}

// Rows represents a set of data.
type Rows []*Row

// MinSpread returns the Row with the minimum Spread value in the set.
func (rs Rows) MinSpread() *Row {
	var msr *Row

	for _, r := range rs {
		if msr == nil || r.Spread() <= msr.Spread() {
			msr = r
		}
	}

	if msr == nil {
		return &Row{}
	}

	return msr
}
