package server

import (
	"log"

	"qShell/cmd/common"
	"qShell/grpc/server"

	"github.com/spf13/cobra"
)

var (
	host       string
	port       string
	username   string
	password   string
	ServerRoot = &cobra.Command{
		Use:   "server",
		Short: "server control",
		Run: func(cmd *cobra.Command, args []string) {
			address := host + ":" + port
			log.Println("Listening: " + address)

			s := new(server.ServerBasic)
			s.ListenAddress = address
			s.Username = username
			s.Password = password
			err := s.Run()
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	ServerRoot.Flags().StringVarP(&host, "host", "", common.DefaultHost, "Listen host")
	ServerRoot.Flags().StringVarP(&port, "port", "", common.DefaultPort, "Listen port")
	ServerRoot.Flags().StringVarP(&username, "username", "u", "", "Username")
	ServerRoot.Flags().StringVarP(&password, "password", "p", "", "Password")
}
