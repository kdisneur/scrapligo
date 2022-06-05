package options

import (
	"github.com/scrapli/scrapligo/channel"
	"github.com/scrapli/scrapligo/driver/network"
	"github.com/scrapli/scrapligo/transport"
	"github.com/scrapli/scrapligo/util"
)

// WithAuthUsername provides the username to use for authentication to a device.
func WithAuthUsername(s string) util.Option {
	return func(o interface{}) error {
		a, ok := o.(*transport.Args)

		if ok {
			a.User = s

			return nil
		}

		return util.ErrIgnoredOption
	}
}

// WithAuthPassword provides the password to use for authentication to a device.
func WithAuthPassword(s string) util.Option {
	return func(o interface{}) error {
		a, ok := o.(*transport.Args)

		if ok {
			a.Password = s

			return nil
		}

		return util.ErrIgnoredOption
	}
}

// WithAuthSecondary provides the "secondary" password to use for authentication to a device -- this
// is usually the "enable" password.
func WithAuthSecondary(s string) util.Option {
	return func(o interface{}) error {
		d, ok := o.(*network.Driver)

		if ok {
			d.AuthSecondary = s

			return nil
		}

		return util.ErrIgnoredOption
	}
}

// WithAuthPassphrase provides the ssh key passphrase to use during authentication to a device.
func WithAuthPassphrase(s string) util.Option {
	return func(o interface{}) error {
		a, ok := o.(*transport.SSHArgs)

		if ok {
			a.PrivateKeyPassPhrase = s

			return nil
		}

		return util.ErrIgnoredOption
	}
}

// WithAuthBypass allows for skipping of "in channel" authentication. This "in channel"
// authentication occurs when connecting with the System or Telnet transports and happens right
// after the initial transport connection is opened up. This option allows for skipping this, you
// may want to do this if connecting to a terminal server or using the System transport with
// a patched open binary (ex: docker/kubectl exec).
func WithAuthBypass() util.Option {
	return func(o interface{}) error {
		c, ok := o.(*channel.Channel)

		if ok {
			c.AuthBypass = true

			return nil
		}

		return util.ErrIgnoredOption
	}
}
