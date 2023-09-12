package controllers

type DataMap map[string]any

func errorResponse(err error) DataMap {
	return DataMap{"error": err.Error()}
}
