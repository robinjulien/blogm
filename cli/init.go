package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/robinjulien/rblog/entities"
)

// write string into filename, creating the file
func writeFile(fileName string, content []byte) {
	f, err := os.Create(fileName)
	check(err)
	_, err = f.Write(content)
	check(err)
	err = f.Close()
	check(err)
}

// initInstance creates all the directories and file that a rblog instance need
func initInstance() {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		file, err := os.Create("config.json")
		check(err)

		// Default config
		cfg := entities.Config{
			Host:                        "localhost",
			Port:                        "8080",
			BlogName:                    "Blog Name",
			BlogLogoURL:                 "//via.placeholder.com/60",
			HomePageTitle:               "Default Homepage Title",
			MaxPostsOnListPage:          10,
			PageTitleSuffix:             " - rblog",
			ListPostsPageTitle:          "Posts List",
			NoPostMessage:               "No post here...",
			InvalidListPostsPageMessage: "Invalid page",
			MenuLinks: []entities.Link{
				{
					Text:  "Home",
					Dest:  "/",
					Title: "Home",
				},
				{
					Text:  "Recent Posts",
					Dest:  "/posts",
					Title: "Recent Posts",
				},
			},
		}

		// Write json into config file
		jsoncfg, err := json.MarshalIndent(cfg, "", "	")
		check(err)
		_, err = file.Write(jsoncfg)
		check(err)
		err = file.Close()
		check(err)

		// create all directories
		err = os.Mkdir("posts", os.ModePerm)
		check(err)
		err = os.Mkdir("cdn", os.ModePerm)
		check(err)
		err = os.Mkdir("pages", os.ModePerm)
		check(err)
		err = os.Mkdir("assets", os.ModePerm)
		check(err)
		err = os.Mkdir("assets/public", os.ModePerm)
		check(err)

		// Here we're going to create the assets file. Note that the files are in the binary,
		// But we let the user the possibility to modify the pages. We don't want the user to have to rebuild the binary everytime he makes a change in the source code.
		// Later, maybe we can let the user the possibility to choose wether he wants the files to be loaded from the assets folder, or directly from the memory. TODO

		// Write the content of each file into the right folder, no way to do that automatically, so we have to tell manually which file goes where, and what is the corresponding const var
		writeFile("./assets/home.tpl", HOME_TPL)
		writeFile("./assets/header.tpl", HEADER_TPL)
		writeFile("./assets/footer.tpl", FOOTER_TPL)
		writeFile("./assets/menu.tpl", MENU_TPL)
		writeFile("./assets/posts.tpl", POSTS_TPL)
		writeFile("./assets/view_post.tpl", VIEW_POST_TPL)
		writeFile("./assets/view_page.tpl", VIEW_PAGE_TPL)
		writeFile("./assets/error_404.tpl", ERROR_404_TPL)

		writeFile("./assets/public/style.css", STYLE_CSS)
		writeFile("./assets/public/markdown.css", MARKDOWN_CSS)
		writeFile("./assets/public/highlight.js", HIGHLIGHT_JS)
		writeFile("./assets/public/highlightjs.css", HIGHLIGHTJS_CSS)

		writeFile("./pages/home.md", HOME_MD)

	} else {
		// config file already exists, there is a chance that a website is already init here
		fmt.Println("A rblog instance is already init here.")
	}
}
