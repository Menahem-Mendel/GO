package lenconv

// MToF ...
func MToF(m Meter) Foot { return Foot(m * 3.281) }

// FToM ...
func FToM(f Foot) Meter { return Meter(f / 3.281) }
