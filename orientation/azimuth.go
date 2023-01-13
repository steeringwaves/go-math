package orientation

// AzimuthAt Returns the local azimuth angle at the target world azimuth in degrees
func (o *Orientation[T]) AzimuthAt(worldAzimuth T) T {
	return ScaleDegrees[T](worldAzimuth-o.Heading, false)
}

// WorldAzimuthAt Returns the world azimuth angle at the local azimuth in degrees
func (o *Orientation[T]) WorldAzimuthAt(azimuth T) T {
	return ScaleDegrees[T](azimuth+o.Heading, false)
}
