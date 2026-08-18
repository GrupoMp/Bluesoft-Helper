package main

import (
	"strings"
	"testing"
)

// Cobre a leitura da coluna DESCRIÇÃO e o limite de 70 caracteres.
func TestLerPlanilhaFichaDescricao(t *testing.T) {
	longa := strings.Repeat("A", 71) // acima de qualquer um dos dois limites (60/70)
	rows := make([][]string, 7)
	rows = append(rows,
		[]string{"EAN 13", "DESCRIÇÃO", "1"},
		[]string{"111", "LIVRO X EB", "1"},
		[]string{"222", longa, "1"},
		[]string{"333", "", "1"},
	)

	produtos, avisos, err := lerPlanilhaFicha(rows)
	if err != nil {
		t.Fatal(err)
	}
	if len(produtos) != 3 {
		t.Fatalf("esperado 3 produtos, veio %d", len(produtos))
	}
	if produtos[0].Descricao != "LIVRO X EB" {
		t.Errorf("descrição não lida: %q", produtos[0].Descricao)
	}
	if produtos[1].Descricao != longa || produtos[2].Descricao != "" {
		t.Errorf("descrições inesperadas: %q / %q", produtos[1].Descricao, produtos[2].Descricao)
	}
	texto := strings.Join(avisos, " | ")
	if !strings.Contains(texto, "1 linhas sem DESCRIÇÃO") || !strings.Contains(texto, "acima de 60") {
		t.Errorf("avisos faltando: %s", texto)
	}
}
