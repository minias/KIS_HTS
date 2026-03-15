// internal/config/watcher.go
package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// WatchConfig enable hot reload
func WatchConfig(onChange func()) {

	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {

		log.Println("config changed:", e.Name)

		onChange()
	})
}
