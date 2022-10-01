package ios

import (
	"fmt"
	"os"

	giDevice "github.com/electricbubble/gidevice"
	"github.com/spf13/cobra"

	"github.com/httprunner/httprunner/v4/hrp/internal/uixt"
)

var iosRootCmd = &cobra.Command{
	Use:              "ios",
	Short:            "simple utils for ios device management",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
}

func getDevice(udid string) (giDevice.Device, error) {
	devices, err := uixt.IOSDevices(udid)
	if err != nil {
		return nil, err
	}
	if len(devices) == 0 {
		fmt.Println("no ios device found")
		os.Exit(1)
	}
	if len(devices) > 1 {
		return nil, fmt.Errorf("multiple devices found, please specify udid")
	}
	return devices[0], nil
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(iosRootCmd)
}
