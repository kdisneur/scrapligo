package options

import (
	"io"
	"regexp"
	"time"

	"github.com/scrapli/scrapligo/channel"
	"github.com/scrapli/scrapligo/util"
)

// WithPromptPattern allows for providing a custom regex pattern to use for the channel
// PromptPattern.
func WithPromptPattern(p *regexp.Regexp) util.Option {
	return func(o interface{}) error {
		c, ok := o.(*channel.Channel)

		if ok {
			c.PromptPattern = p

			return nil
		}

		return util.ErrIgnoredOption
	}
}

// WithUsernamePattern allows for patching the "in channel" authentication username pattern -- this
// is only relevant when using the Telnet transport.
func WithUsernamePattern(p *regexp.Regexp) util.Option {
	return func(o interface{}) error {
		c, ok := o.(*channel.Channel)

		if ok {
			c.UsernamePattern = p

			return nil
		}

		return util.ErrIgnoredOption
	}
}

// WithPasswordPattern allows for patching the "in channel" authentication password prompt pattern,
// this is only relevant for Telnet and System transports.
func WithPasswordPattern(p *regexp.Regexp) util.Option {
	return func(o interface{}) error {
		c, ok := o.(*channel.Channel)

		if ok {
			c.PasswordPattern = p

			return nil
		}

		return util.ErrIgnoredOption
	}
}

// WithPassphrasePattern allows for patching the "in channel" authentication SSH key passphrase
// pattern.
func WithPassphrasePattern(p *regexp.Regexp) util.Option {
	return func(o interface{}) error {
		c, ok := o.(*channel.Channel)

		if ok {
			c.PassphrasePattern = p

			return nil
		}

		return util.ErrIgnoredOption
	}
}

// WithReturnChar allows for patching the channel ReturnChar value -- *typically* you should not
// need to use this option.
func WithReturnChar(s string) util.Option {
	return func(o interface{}) error {
		c, ok := o.(*channel.Channel)

		if ok {
			c.ReturnChar = []byte(s)

			return nil
		}

		return util.ErrIgnoredOption
	}
}

// WithTimeoutOps allows for modifying the channel TimeoutOps value -- this is the value that gets
// set as the TimeoutOps for the Channel at driver creation. The TimeoutOps value is what governs
// the "operation" timeouts for Channel read operations.
func WithTimeoutOps(t time.Duration) util.Option {
	return func(o interface{}) error {
		c, ok := o.(*channel.Channel)

		if ok {
			c.TimeoutOps = t

			return nil
		}

		return util.ErrIgnoredOption
	}
}

// WithReadDelay sets the ReadDelay for the channel read loop. This value is the sleep time between
// dequeue'ing data from the read queue that the transport fills. This value defaults to 5ms.
func WithReadDelay(t time.Duration) util.Option {
	return func(o interface{}) error {
		c, ok := o.(*channel.Channel)

		if ok {
			c.ReadDelay = t

			return nil
		}

		return util.ErrIgnoredOption
	}
}

// WithChannelLog accepts an io.Writer that can be used to write all Channel read data out to.
func WithChannelLog(w io.Writer) util.Option {
	return func(o interface{}) error {
		c, ok := o.(*channel.Channel)

		if ok {
			c.ChannelLog = w

			return nil
		}

		return util.ErrIgnoredOption
	}
}
