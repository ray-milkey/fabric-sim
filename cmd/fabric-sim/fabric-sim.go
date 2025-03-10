// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/onosproject/fabric-sim/pkg/manager"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var log = logging.GetLogger()

// The main entry point
func main() {
	if err := getRootCommand().Execute(); err != nil {
		println(err)
		os.Exit(1)
	}
}

func getRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fabric-sim",
		Short: "fabric-sim",
		RunE:  runRootCommand,
	}
	cmd.Flags().String("caPath", "", "path to CA certificate")
	cmd.Flags().String("keyPath", "", "path to client private key")
	cmd.Flags().String("certPath", "", "path to client certificate")
	cmd.Flags().String("topoDescription", "", "topology description YAML file")
	return cmd
}

func runRootCommand(cmd *cobra.Command, args []string) error {
	caPath, _ := cmd.Flags().GetString("caPath")
	keyPath, _ := cmd.Flags().GetString("keyPath")
	certPath, _ := cmd.Flags().GetString("certPath")
	topoDescription, _ := cmd.Flags().GetString("topoDescription")

	log.Infow("Starting fabric-sim",
		"CAPath", caPath,
		"KeyPath", keyPath,
		"CertPath", certPath,
		"topoDescription", topoDescription,
	)

	cfg := manager.Config{
		CAPath:          caPath,
		KeyPath:         keyPath,
		CertPath:        certPath,
		TopoDescription: topoDescription,
		GRPCPort:        5150,
	}

	mgr := manager.NewManager(cfg)

	mgr.Run()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	mgr.Close()
	return nil
}
