package media

func MediaNotas(values ...float64) (media float64) {
	var total float64
	for _, value := range values {
		total += value
	}

	media = total / float64(len(values))
	return media
}
