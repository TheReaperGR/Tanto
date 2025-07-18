package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/TheReaperGR/Tanto/pkg/config"
    "github.com/TheReaperGR/Tanto/pkg/proxmox"
    "github.com/TheReaperGR/Tanto/pkg/clab"
)

func main() {
    var cfgPath string

    root := &cobra.Command{
        Use:   "tanto",
        Short: "Manage Proxmox bridges/VMs + Containerlab labs",
    }
    root.PersistentFlags().StringVarP(&cfgPath, "config", "c", "tanto.yml", "path to YAML config")

    upCmd := &cobra.Command{
        Use:   "up [topo-file]",
        Short: "Create infra then deploy clab topology",
        Args:  cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            topo := args[0]
            cfg, err := config.Load(cfgPath)
            if err != nil { return err }

            pmx := proxmox.NewProxmoxClient(cfg.Proxmox)
            // 1. create bridges
            // 2. ensure VMs + attach NICs

            cl := clab.NewClabClient(cfg.Containerlab)
            labID, err := cl.Deploy(topo)
            if err != nil { return err }
            fmt.Println("Deployed lab", labID)
            return nil
        },
    }
    root.AddCommand(upCmd)

    destroyCmd := &cobra.Command{
        Use:   "destroy [lab-id]",
        Short: "Tear down a clab lab",
        Args:  cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            cfg, _ := config.Load(cfgPath)
            cl := clab.NewClabClient(cfg.Containerlab)
            return cl.Destroy(args[0])
        },
    }
    root.AddCommand(destroyCmd)

    if err := root.Execute(); err != nil {
        os.Exit(1)
    }
}
