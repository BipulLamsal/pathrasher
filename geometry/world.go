package geometry

type World struct {
	Objects []Hittable
}

func (w *World) Add(object Hittable) {
	w.Objects = append(w.Objects, object)
}

func (w *World) Hit(ray *Ray, tMin, tMax float64, rec *HitRecord) bool {
	tempRec := HitRecord{} // Temporary storage for each hit test
	hitAnything := false   // Did we hit ANY object?
	closestSoFar := tMax   // Start with maximum search distance

	for _, object := range w.Objects {
		// Check if ray hits this object between tMin and closestSoFar
		if object.Hit(ray, tMin, closestSoFar, &tempRec) {
			hitAnything = true       // Yes, we hit something!
			closestSoFar = tempRec.T // Update search range to this closer hit
			*rec = tempRec           // Save this hit as the current closest
		}
	}

	return hitAnything // Return true if we hit anything
}
