package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"Tirta", "Prana", "Bhakti"})
	_ = writer.Write([]string{"Bayu", "Prayudha"})
	_ = writer.Write([]string{"Putri", "Salwaa", "Nanda", "Persada"})

	writer.Flush()

}
