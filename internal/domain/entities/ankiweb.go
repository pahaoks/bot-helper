package entities

type AnkiRequest struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
	Params  struct {
		Note struct {
			DeckName  string `json:"deckName"`
			ModelName string `json:"modelName"`
			Fields    struct {
				Front string `json:"Front"`
				Back  string `json:"Back"`
			} `json:"fields"`
			Options struct {
				AllowDuplicate        bool   `json:"allowDuplicate"`
				DuplicateScope        string `json:"duplicateScope"`
				DuplicateScopeOptions struct {
					DeckName       string `json:"deckName"`
					CheckChildren  bool   `json:"checkChildren"`
					CheckAllModels bool   `json:"checkAllModels"`
				} `json:"duplicateScopeOptions"`
			} `json:"options"`
			Tags  []string `json:"tags"`
			Audio []struct {
				URL      string   `json:"url"`
				Filename string   `json:"filename"`
				SkipHash string   `json:"skipHash"`
				Fields   []string `json:"fields"`
			} `json:"audio"`
			Video []struct {
				URL      string   `json:"url"`
				Filename string   `json:"filename"`
				SkipHash string   `json:"skipHash"`
				Fields   []string `json:"fields"`
			} `json:"video"`
			Picture []struct {
				URL      string   `json:"url"`
				Filename string   `json:"filename"`
				SkipHash string   `json:"skipHash"`
				Fields   []string `json:"fields"`
			} `json:"picture"`
		} `json:"note"`
	} `json:"params"`
}
