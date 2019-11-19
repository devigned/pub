package args_test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/devigned/pub/cmd/args"
)

func TestBind(t *testing.T) {
	cases := []struct {
		Name     string
		BindFunc func(command *cobra.Command, binder *string) error
		FlagName string
	}{
		{
			Name:     "Publisher",
			BindFunc: args.BindPublisher,
			FlagName: "publisher",
		},
		{
			Name:     "Offer",
			BindFunc: args.BindOffer,
			FlagName: "offer",
		},
	}

	arg := struct {
		binder string
	}{}

	for _, tc := range cases {
		c := tc
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()

			cmd := new(cobra.Command)
			assert.NoError(t, c.BindFunc(cmd, &arg.binder))
			f := cmd.Flag(c.FlagName)
			assert.Equal(t, []string{"true"}, f.Annotations[cobra.BashCompOneRequiredFlag])
		})
	}

}
