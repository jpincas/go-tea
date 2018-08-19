package gotea

func (existingMap MessageMap) MergeMap(thisMap MessageMap) {
	for k, v := range thisMap {
		existingMap[k] = v
	}
}
