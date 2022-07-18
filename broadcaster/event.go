package broadcaster

type Event struct {
	Sender  string `json:"sender"`
	Id      string `json:"id"`
	Content string `json:"content"`
}
