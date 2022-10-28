package tailnet

import (
	"context"
	"os"
	"time"

	"tailscale.com/client/tailscale"
	"tailscale.com/tsnet"
)

func WaitForOnline(client *tailscale.LocalClient) error {
	for {
		status, err := client.Status(context.Background())
		if err != nil {
			return err
		}

		if status.Self.Online {
			return nil
		}

		time.Sleep(1 * time.Second)
	}
}

func NewServer(dir string, hostname string, authKey string) (*tsnet.Server, error) {
	s := tsnet.Server{
		Dir:       dir,
		Ephemeral: true,
		AuthKey:   authKey,
		Hostname:  hostname,
		Logf:      func(format string, args ...any) {},
	}

	if os.Getenv("TAILSCALE_DEBUG") != "" {
		s.Logf = nil
	}

	err := s.Start()
	if err != nil {
		return nil, err
	}

	return &s, nil
}
