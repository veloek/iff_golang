food, err := hunt()
if err != nil {
	log.Errorf("Couldn't find any food: %s", err)
	return
}
