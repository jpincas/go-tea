package gotea

func (existingMap MessageMap) MergeMap(thisMap MessageMap) MessageMap {
	for k, v := range thisMap {
		existingMap[k] = v
	}

	return existingMap
}
