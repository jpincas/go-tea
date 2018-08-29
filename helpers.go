package gotea

func mergeMaps(startMap MessageMap, msgMaps ...MessageMap) MessageMap {
	for _, thisMap := range msgMaps {
		for k, v := range thisMap {
			startMap[k] = v
		}
	}

	return startMap
}
