package care

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/spf13/pflag"
	"k8s.io/klog/v2"

	"github.com/labring/sealos/pkg/constants"
	"github.com/labring/sealos/pkg/utils/hosts"
)

const (
	appName   = "lvscare"
	routeMode = "route"
	linkMode  = "link"
)

type options struct {
	VirtualServer string
	RealServer    []string
	scheduler     string
	IfaceName     string
	Logger        string
	Mode          string
	RunOnce       bool
	CleanAndExit  bool
	Interval      time.Duration
	TargetIP      net.IP
	MasqueradeBit int
}

func (o *options) RegisterFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.VirtualServer, "vs", "", "virtual server address, for example 169.254.0.1:6443")
	fs.StringSliceVar(&o.RealServer, "rs", []string{}, "real server address like 192.168.0.2:6443")
	fs.StringVar(&o.scheduler, "scheduler", "rr", "proxier scheduler")
	fs.StringVarP(&o.IfaceName, "iface", "i", appName, "name of dummy interface to created, same behavior as kube-proxy")
	fs.StringVar(&o.Logger, "logger", "INFO", "logger level: DEBG/INFO")
	fs.StringVar(&o.Mode, "mode", routeMode, fmt.Sprintf("proxy mode: %s/%s", routeMode, linkMode))
	fs.BoolVar(&o.RunOnce, "run-once", false, "create proxy rules and exit")
	fs.BoolVarP(&o.CleanAndExit, "clean", "C", false, "clean existing rules and then exit")
	fs.DurationVar(&o.Interval, "interval", 5*time.Second, "health check interval")
	fs.IPVar(&o.TargetIP, "ip", nil, "target ip as route gateway, use with route mode")
	fs.IntVar(&o.MasqueradeBit, "masqueradebit", 0, "IPTables masquerade bit")

	// set klog flag
	if v := os.Getenv("ENABLE_KLOG_FLAGS"); len(v) > 0 {
		if b, _ := strconv.ParseBool(v); b {
			klog.InitFlags(nil)
			fs.AddGoFlagSet(flag.CommandLine)
		}
	}
}

func (o *options) RequiredFlags() []string {
	return []string{"vs"}
}

func (o *options) ValidateAndSetDefaults() error {
	if len(o.RealServer) == 0 && !o.CleanAndExit {
		return errors.New(`required flag(s) "rs" not set`)
	}
	switch o.scheduler {
	case "rr", "lc", "dh", "sh", "wrr", "wlc":
	default:
		return fmt.Errorf(`invalid flag "scheduler=%s"`, o.scheduler)
	}
	if o.TargetIP == nil && o.Mode == routeMode {
		hf := &hosts.HostFile{Path: constants.DefaultHostsPath}
		if ip, ok := hf.HasDomain(constants.DefaultLvscareDomain); ok {
			o.TargetIP = net.ParseIP(ip)
		}
	}
	return nil
}
