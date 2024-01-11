package formas

import (
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaRecebida != areaEsperada {
			t.Fatal("A área é recebida é diferente da esperada")
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}

		areaEsperada := float64(3.14 * 100)
		areaRecebida := circ.Area()

		if areaRecebida != areaEsperada {
			t.Fatal("A área é recebida é diferente da esperada")
		}
	})
}
