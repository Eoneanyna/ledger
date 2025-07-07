package conf

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	_, err := InitConfig()
	require.NoError(t, err)

	fmt.Printf("ProjectName: %s\n", Conf.ProjectName)
	fmt.Printf("Server Port: %d\n", Conf.Server.Port)
	fmt.Printf("Database IP: %s\n", Conf.Database.IP)
}
