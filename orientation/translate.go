package orientation

// TranslateFrom Returns a translated orientation for the specified azimuth amount
func (o *Orientation[T]) TranslateFrom(azimuth T) Orientation[T] {
	res := Orientation[T]{
		Heading: ScaleDegrees[T](o.Heading-azimuth, false),
	}

	res.Pitch = o.PitchAt(-1 * azimuth)
	res.Roll = o.RollAt(-1 * azimuth)

	return res
}
