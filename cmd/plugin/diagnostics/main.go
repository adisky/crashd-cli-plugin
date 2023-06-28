package main

import (
	"embed"
	"os"

	"github.com/vmware-tanzu/crashd-cli-plugin/pkg"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/log"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/plugin"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/plugin/buildinfo"
)

var descriptor = plugin.PluginDescriptor{
	Name:        "diagnostics",
	Description: "crashd plugin",
	Target:      types.TargetGlobal, // <<<FIXME! set the Target of the plugin to one of {TargetGlobal,TargetK8s,TargetTMC}
	Version:     buildinfo.Version,
	BuildSHA:    buildinfo.SHA,
	Group:       plugin.RunCmdGroup, // set group
}

var (
	//go:embed scripts
	scriptFS embed.FS
)

func main() {
	p, err := plugin.NewPlugin(&descriptor)
	if err != nil {
		log.Fatal(err, "")
	}
	p.AddCommands(pkg.CollectCmd(scriptFS))
	if err := p.Execute(); err != nil {
		os.Exit(1)
	}
}
