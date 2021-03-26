package protein

import (
	"errors"
)

/*
Codon                 | Protein
:---                  | :---
AUG                   | Methionine
UUU, UUC              | Phenylalanine
UUA, UUG              | Leucine
UCU, UCC, UCA, UCG    | Serine
UAU, UAC              | Tyrosine
UGU, UGC              | Cysteine
UGG                   | Tryptophan
UAA, UAG, UGA         | STOP
*/

//ErrStop is used when a stop codon is detected
var ErrStop error = errors.New("Stop")

//ErrInvalidBase is used when a unrecognizable codon is detected
var ErrInvalidBase error = errors.New("Invalid base")

var codToProt = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

//FromCodon translates a codon to a protien
func FromCodon(codon string) (string, error) {
	protein, ok := codToProt[codon]

	if !ok {
		return "", ErrInvalidBase
	}
	if protein == "STOP" {
		return "", ErrStop
	}
	return protein, nil
}

//FromRNA translates the long form rna string into a slice of proteins
func FromRNA(rna string) ([]string, error) {
	codons := len(rna) / 3
	var proteins []string
	for i := 0; i < codons; i++ {
		protein, err := FromCodon(rna[i*3 : i*3+3])
		if err == ErrStop {
			break
		}
		if err == ErrInvalidBase {
			return proteins, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}
