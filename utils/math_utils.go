package utils

import "math"

// floating point comparison delta https://stackoverflow.com/a/47969546/3679900
const FLOAT32_DELTA float32 = 1e-9

// fn needs to be named (starting) with uppercase to be 'import-able'
// default value maxDeltaOpt in Go fns https://stackoverflow.com/a/23650312/3679900
func AreEqual(f1, f2 float32, maxDeltaOpt ...float32) bool {
	maxDelta := FLOAT32_DELTA
	if len(maxDeltaOpt) > 0 {
		maxDelta = maxDeltaOpt[0]
	}

	return float32(math.Abs(float64(f1-f2))) <= maxDelta
}
