package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/jpillora/opts"
	"github.com/jpillora/requestlog"
	"github.com/jpillora/serve/handler"
)

var VERSION string = "0.0.0"

type Config struct {
	Host           string `help:"Host interface"`
	Port           int    `help:"Listening port"`
	Open           bool   `help:"On server startup, open the root in the default browser (uses the 'open' command)"`
	handler.Config `type:"embedded"`
}

func main() {

	//defaults
	c := Config{
		Host: "0.0.0.0",
		Port: 3000,
		Config: handler.Config{
			Directory: ".",
		},
	}

	//parse
	opts.New(&c).
		Name("serve").
		Version(VERSION).
		Repo("github.com/jpillora/serve").
		Parse()

	//ready!
	h, err := handler.New(c.Config)
	if err != nil {
		log.Fatal(err)
	}

	port := strconv.Itoa(c.Port)

	if c.Open {
		go func() {
			time.Sleep(500 * time.Millisecond)
			cmd := exec.Command("open", "http://localhost:"+port)
			cmd.Run()
		}()
	}

	fmt.Printf("%sserving %s%s %son port %s%d%s\n",
		requestlog.DefaultOptions.Colors.Grey,
		requestlog.DefaultOptions.Colors.Cyan, c.Config.Directory,
		requestlog.DefaultOptions.Colors.Grey,
		requestlog.DefaultOptions.Colors.Cyan, c.Port,
		requestlog.DefaultOptions.Colors.Reset,
	)

	log.Fatal(http.ListenAndServe(":"+port, h))
}
