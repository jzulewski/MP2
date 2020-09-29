package structs

type Client struct {
	IP       string
	PORT     string
	Username string
}

type Message struct {
	To      string
	From    string
	Content string
}
