package entities

// Link is a link in the nav menu
type Link struct {
	Text  string `json:"text"`
	Dest  string `json:"dest"`
	Title string `json:"title"`
}

// Config represents the config located in the config.json file
type Config struct {
	Host               string `json:"host"`
	Port               string `json:"port"`
	BlogName           string `json:"blogName"`
	BlogLogoURL        string `json:"blogLogoURL"`
	HomePageTitle      string `json:"homePageTitle"`
	MaxPostsOnListPage int    `json:"maxPostsOnListPage"`
	PageTitleSuffix    string `json:"pageTitleSuffix"`
	ListPostsPageTitle string `json:"listPostsPageTitle"`
	NoPostMessage      string `json:"noPostMessage"`
	MenuLinks          []Link `json:"menuLinks"`
}
