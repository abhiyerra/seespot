// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var (
	config struct {
		HealthPort string
		HealthPath string

		AppHealth string

		CleanupTask string

		Terminated bool
	}
)

func healthHandler() {
	http.HandleFunc(config.HealthPath, func(w http.ResponseWriter, r *http.Request) {
		if config.Terminated {
			http.Error(w, "Machine Terminating", 503)
			return
		}

		resp, err := http.Get(config.AppHealth)
		if err != nil {
			fmt.Println(err)
		}

		if resp.StatusCode != 200 {
			http.Error(w, "App down", 404)
		}

		w.Write([]byte("OK"))
	})

	http.ListenAndServe(config.HealthPort, nil)
}

func terminationRunner() {
	cmd := exec.Command("sh", "-c", config.CleanupTask)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func terminated() bool {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/spot/termination-time")
	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode == 200 {
		return true
	}

	return false
}

func watchForTermination() {
	for {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("timed out")
			if config.Terminated = terminated(); config.Terminated {
				return
			}
		}
	}
}

func main() {
	flag.StringVar(&config.HealthPort, "health-port", ":8686", "Default health port to use with Load Balancers")
	flag.StringVar(&config.HealthPath, "health-path", "/health", "Default health path the Load Balancer hits")
	flag.StringVar(&config.AppHealth, "app-health", "http://127.0.0.1:8080/health", "Application health check")
	flag.StringVar(&config.CleanupTask, "cleanup-task", "", "Script to run upon termination")
	flag.Parse()

	config.Terminated = false

	go healthHandler()
	watchForTermination()
	terminationRunner()
}
