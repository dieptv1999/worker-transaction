package responses

import "worker-transaction/models"

type ElasticLogResponse struct {
	Hits struct {
		Hits []struct {
			Source ElasticLogSource `json:"_source"`
			// Source Logging `json:"_source"`
		} `json:"hits"`
		Total struct {
			Value int64 `json:"value"`
		} `json:"total"`
	} `json:"hits"`
}

type ElasticLogSource struct {
	Data models.Logging `json:"Data"`
}
