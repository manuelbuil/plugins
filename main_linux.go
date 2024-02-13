//go:build linux
// +build linux

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/docker/pkg/reexec"

	"github.com/containernetworking/cni/pkg/version"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
	hostlocal "github.com/containernetworking/plugins/plugins/ipam/host-local"
	"github.com/containernetworking/plugins/plugins/ipam/static"
	"github.com/containernetworking/plugins/plugins/main/bridge"
	hostdevice "github.com/containernetworking/plugins/plugins/main/host-device"
	"github.com/containernetworking/plugins/plugins/main/ipvlan"
	"github.com/containernetworking/plugins/plugins/main/loopback"
	"github.com/containernetworking/plugins/plugins/main/macvlan"
	"github.com/containernetworking/plugins/plugins/main/vlan"
	"github.com/containernetworking/plugins/plugins/meta/bandwidth"
	"github.com/containernetworking/plugins/plugins/meta/firewall"
	"github.com/containernetworking/plugins/plugins/meta/flannel"
	"github.com/containernetworking/plugins/plugins/meta/multus"
	"github.com/containernetworking/plugins/plugins/meta/portmap"
)

func main() {
	os.Args[0] = filepath.Base(os.Args[0])
	reexec.Register("bandwidth", bandwidth.Main)
	reexec.Register("bridge", bridge.Main)
	reexec.Register("host-device", hostdevice.Main)
	reexec.Register("ipvlan", ipvlan.Main)
	reexec.Register("firewall", firewall.Main)
	reexec.Register("flannel", flannel.Main)
	reexec.Register("host-local", hostlocal.Main)
	reexec.Register("loopback", loopback.Main)
	reexec.Register("macvlan", macvlan.Main)
	reexec.Register("multus", multus.Main)
	reexec.Register("portmap", portmap.Main)
	reexec.Register("static", static.Main)
	reexec.Register("vlan", vlan.Main)
	if !reexec.Init() {
		_, _ = fmt.Fprintln(os.Stderr, bv.BuildString("plugins"))
		_, _ = fmt.Fprintf(os.Stderr, "CNI protocolo versions supported: %s\n", strings.Join(version.All.SupportedVersions(), ", "))
	}
}
