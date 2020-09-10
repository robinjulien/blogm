package entities

// Config represents the config located in the config.json file
type Config struct {
	Host               string `json:"host"`
	Port               string `json:"port"`
	BlogName           string `json:"blogName"`
	HomePageTitle      string `json:"homePageTitle"`
	MaxPostsOnListPage int    `json:"maxPostsOnListPage"`
	PageTitleSuffix    string `json:"pageTitleSuffix"`
	ListPostsPageTitle string `json:"listPostsPageTitle"`
}
