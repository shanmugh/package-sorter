package main

import (
	"github.com/shanmugh/package-sorter/cmd"
	"go.uber.org/zap"
)

func main() {
	rootCmd, err := cmd.NewCommand()
	if err != nil {
		zap.L().Fatal("create root command", zap.Error(err))
	}

	if err := rootCmd.Execute(); err != nil {
		zap.L().Fatal("execute root command", zap.Error(err))
	}
}
