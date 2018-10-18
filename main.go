package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "JSONify"
	app.Usage = "Easily Convert from a JSON string to pretty printed JSON"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "file-path,file,f",
			Usage:  "Path to a file whence JSON string can be read",
			EnvVar: "FILE_PATH",
		},
	}

	app.Action = jsonify

	app.Run(os.Args)
}

func jsonify(c *cli.Context) {
	filePath := c.GlobalString("file-path")

	fileData, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	var rawJson interface{}
	json.Unmarshal([]byte(fileData), &rawJson)

	fmt.Println(rawJson)
}
