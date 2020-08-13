package cli

import (
	"encoding/json"
	"fmt"
	"os"
)

// write string into filename, creating the file
func writeFile(fileName, content string) {
	f, err := os.Create(fileName)
	check(err)
	_, err = f.WriteString(content)
	check(err)
	err = f.Close()
	check(err)
}

// initInstance creates all the directories and file that a blogm instance need
func initInstance() {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		file, err := os.Create("config.json")
		check(err)

		// Default config
		cfg := Config{
			Host:     "localhost",
			Port:     "8080",
			BlogName: "default blog name",
		}

		jsoncfg, err := json.MarshalIndent(cfg, "", "	")
		check(err)
		_, err = file.Write(jsoncfg)
		check(err)
		err = file.Close()
		check(err)

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

		writeFile("./assets/home.tpl", HOME_TPL)
		writeFile("./assets/public/style.css", STYLE_CSS)
		writeFile("./assets/header.tpl", HEADER_TPL)
		writeFile("./assets/footer.tpl", FOOTER_TPL)
		writeFile("./assets/menu.tpl", MENU_TPL)
		writeFile("./assets/posts.tpl", POSTS_TPL)
		writeFile("./assets/view_post.tpl", VIEW_POST_TPL)
		writeFile("./assets/view_page.tpl", VIEW_PAGE_TPL)

	} else {
		// config file already exists, there is a chance that a website is already init here
		fmt.Println("A blogm instance is already init here.")
	}
}
