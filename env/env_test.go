package env

import (
	"os"
	"testing"

	. "github.com/franela/goblin"
)

func TestGet(t *testing.T) {
	g := Goblin(t)
	g.Describe("Variable#Get test", func() {
		g.It("Should parse the correct data", func() {
			mockURL := "http://mock.test"
			os.Setenv("CB_PRO_URL", mockURL)
			g.Assert(CoinbaseProURL.Get()).Equal(mockURL)
		})
	})
}

func TestName(t *testing.T) {
	g := Goblin(t)
	g.Describe("Variable#Get test", func() {
		g.It("Should parse the correct data", func() {
			g.Assert(CoinbaseProURL.Name()).Equal("CB_PRO_URL")
		})
	})
}

func TestLoad(t *testing.T) {
	g := Goblin(t)
	g.Describe("Load test", func() {
		g.It("Should load environment data from any file", func() {
			Load(".test.env")
			g.Assert(CoinbaseProSecret.Get()).Equal("some-beautiful-secret-just-between-us")
		})
	})
}
