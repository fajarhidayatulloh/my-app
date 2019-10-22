// Copyright © 2018 Trio Purnomo <trio.purnomo@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/my-app/infrastructures"
	"github.com/my-app/routes"
	"github.com/spf13/cobra"
)

// serveHTTPCmd represents the serveHttp command
var serveHTTPCmd = &cobra.Command{
	Use:   "serveHttp",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
			and usage of using your command. For example:

			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		route := new(routes.Route)
		router := route.Init()

		consul := new(infrastructures.ServiceDiscovery)
		consul.Register("service-client", "127.0.0.1", port())

		var gracefulStop = make(chan os.Signal)
		signal.Notify(gracefulStop, syscall.SIGTERM)
		signal.Notify(gracefulStop, syscall.SIGINT)

		go func() {
			sig := <-gracefulStop
			fmt.Printf("caught sig: %+v", sig)
			fmt.Println("Wait for 2 second to finish processing")
			time.Sleep(2 * time.Second)
			consul.DeRegister("service-client")
			os.Exit(0)
		}()

		if err := http.ListenAndServe(port(), router); err != nil {
			consul.DeRegister("service-client")
			log.Fatal("Unable to start service")
		}
	},
}

func init() {
	rootCmd.AddCommand(serveHTTPCmd)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8508"
	}

	return ":" + port
}

func hostname() string {
	hostname, _ := os.Hostname()
	return hostname
}
