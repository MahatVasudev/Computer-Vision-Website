package response_success

func DataMessage(data any) map[string]any {
	table := map[string]any{"data": data}
	return table
}
