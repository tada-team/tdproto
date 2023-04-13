package api_paths

func init() {
	for key, pathType := range AllPaths {
		for i := range pathType {
			AllPaths[key][i].Tags = append(AllPaths[key][i].Tags, key)
		}
	}
}
