package helper

func HandleFilters(query map[string][]string) map[string]interface{} {
	query_map := make(map[string]interface{}, len(query))
	for k, v := range query {
		query_map[k] = v
	}
	return query_map
}
