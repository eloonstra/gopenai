package response

type Model struct {
	Id         string `json:"id"`
	Object     string `json:"object"`
	OwnedBy    string `json:"owned_by"`
	Permission []struct {
		AllowCreateEngine  bool    `json:"allow_create_engine"`
		AllowFineTuning    bool    `json:"allow_fine_tuning"`
		AllowLogProbs      bool    `json:"allow_log_probs"`
		AllowSampling      bool    `json:"allow_sampling"`
		AllowSearchIndices bool    `json:"allow_search_indices"`
		AllowView          bool    `json:"allow_view"`
		Created            float64 `json:"created"`
		Group              string  `json:"group"`
		Id                 string  `json:"id"`
		Object             string  `json:"object"`
		IsBlocking         bool    `json:"is_blocking"`
		Organization       string  `json:"organization"`
	} `json:"permission"`
}

type Models struct {
	Data   []Model `json:"data"`
	Object string  `json:"object"`
}
