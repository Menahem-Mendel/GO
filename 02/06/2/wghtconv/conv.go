package wghtconv

// PToK ...
func PToK(p Pound) Kilogram { return Kilogram(p / 2.205) }

// KToP ...
func KToP(k Kilogram) Pound { return Pound(k * 2.205) }
