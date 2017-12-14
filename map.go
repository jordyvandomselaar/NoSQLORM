package main

func GetFieldsFromMap(input map[string]string, fields []string) map[string]string {
	out := make(map[string]string)

	for _, field := range fields {
		out[field] = input[field]
	}

	return out
}
