package poker_test

import (
	"strings"
	"testing"

	poker "github.com/krls08/go-with-tests/23_APP/27_Command_line_package_structure"
)

func TestCLI(t *testing.T) {

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		aPlayerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(aPlayerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, aPlayerStore, "Chris")

	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		aPlayerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(aPlayerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, aPlayerStore, "Cleo")

	})
}
