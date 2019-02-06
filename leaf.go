package leaf

import (
	"os"
	"os/signal"

	"github.com/BenjaminDuchadeuil/leaf/cluster"
	"github.com/BenjaminDuchadeuil/leaf/conf"
	"github.com/BenjaminDuchadeuil/leaf/console"
	"github.com/BenjaminDuchadeuil/leaf/log"
	"github.com/BenjaminDuchadeuil/leaf/module"
)

func Run(mods ...module.Module) {
	// logger
	if conf.LogLevel != "" {
		logger, err := log.New(conf.LogLevel, conf.LogPath, conf.LogFlag)
		if err != nil {
			panic(err)
		}
		log.Export(logger)
		defer logger.Close()
	}

	log.Release("Serveur lancé")

	// module
	for i := 0; i < len(mods); i++ {
		module.Register(mods[i])
	}
	module.Init()

	// cluster
	cluster.Init()

	// console
	console.Init()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Release("Fermeture du serveur (signal: %v)", sig)
	console.Destroy()
	cluster.Destroy()
	module.Destroy()
}
