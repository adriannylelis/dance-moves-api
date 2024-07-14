package movement

type Movement struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Difficulty  int    `json:"difficulty"`
	Category    string `json:"category"`
	ImageURL    string `json:"imageUrl"`
	VideoURL    string `json:"videoUrl"`
}
