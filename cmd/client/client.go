package client

import (
	"fmt"
	"log"

	"qShell/cmd/common"
	"qShell/grpc/client"

	"github.com/spf13/cobra"
)

var (
	host        string
	port        string
	username    string
	password    string
	timeout     int
	scriptName  string
	scriptParas []string
	scriptFile  string
	ClientRoot  = &cobra.Command{
		Use:   "client",
		Short: "client control",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	pingCmd = &cobra.Command{
		Use:   "ping",
		Short: "Ping Server",
		Run: func(cmd *cobra.Command, args []string) {
			address := host + ":" + port
			log.Println("Connecting: " + address)

			c := new(client.ClientBasic)
			c.ServerAddress = address
			c.Timeout = timeout

			result, err := c.ServerCheck(username, password)
			if err != nil {
				common.GenMessage("ping", common.Error, err.Error())
			} else {
				common.GenMessage("ping", common.Succeed, result)
			}
		},
	}
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all scripts",
		Run: func(cmd *cobra.Command, args []string) {
			address := host + ":" + port
			log.Println("Connecting: " + address)

			c := new(client.ClientBasic)
			c.ServerAddress = address
			c.Timeout = timeout

			scripts, err := c.ListScripts(username, password)
			if err != nil {
				common.GenMessage("list", common.Error, err.Error())
			} else {
				if len(scripts) != 0 {
					common.GenMessage("list", common.Succeed, "your scripts")
					for _, script := range scripts {
						fmt.Println(script)
					}
				} else {
					common.GenMessage("list", common.Succeed, "no script")
				}
			}
		},
	}
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run a scripts",
		Run: func(cmd *cobra.Command, args []string) {
			address := host + ":" + port
			log.Println("Connecting: " + address)

			c := new(client.ClientBasic)
			c.ServerAddress = address
			c.Timeout = timeout

			results, err := c.RunScript(username, password, scriptName, scriptParas)
			if err != nil {
				common.GenMessage("run", common.Error, err.Error())
			} else {
				fmt.Println(results)
			}
		},
	}
	uploadCmd = &cobra.Command{
		Use:   "upload",
		Short: "Upload file",
		Run: func(cmd *cobra.Command, args []string) {
			address := host + ":" + port
			log.Println("Connecting: " + address)

			c := new(client.ClientBasic)
			c.ServerAddress = address
			c.Timeout = timeout

			err := c.Upload(username, password, scriptFile)
			if err != nil {
				common.GenMessage("upload", common.Error, err.Error())
			} else {
				common.GenMessage("upload", common.Error, "received")
			}
		},
	}
)

func init() {
	ClientRoot.PersistentFlags().StringVarP(&host, "host", "", common.DefaultHost, "Listen host")
	ClientRoot.PersistentFlags().StringVarP(&port, "port", "", common.DefaultPort, "Listen port")
	ClientRoot.PersistentFlags().StringVarP(&username, "username", "u", "", "Username")
	ClientRoot.PersistentFlags().StringVarP(&password, "password", "p", "", "Password")
	ClientRoot.PersistentFlags().IntVarP(&timeout, "timeout", "", 15, "Timeout")

	runCmd.Flags().StringVarP(&scriptName, "name", "", "", "Script name")
	runCmd.Flags().StringSliceVarP(&scriptParas, "paras", "", nil, "Script paras (para1,para2,para3)")
	runCmd.MarkFlagRequired("name")
	runCmd.MarkFlagRequired("paras")

	uploadCmd.Flags().StringVarP(&scriptFile, "file", "f", "", "Script File")

	ClientRoot.AddCommand(pingCmd)
	ClientRoot.AddCommand(listCmd)
	ClientRoot.AddCommand(runCmd)
	ClientRoot.AddCommand(uploadCmd)
}
