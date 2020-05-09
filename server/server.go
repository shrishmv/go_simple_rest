package server

import (
	"fmt"
	"go_simple_rest/config"
	"sync"
)

var onceRest sync.Once

// Init function to initialize the service
func Init() {

	// Server should be initialized only once
	onceRest.Do(func() {
		conf := config.GetConfig()
		fmt.Println("Initializing Rest server")
		r := NewRouter()
		if err := r.Start(conf.GetString("general.server_port")); err != nil {
			fmt.Println("Unable to bring service up")
		}
	})
}
