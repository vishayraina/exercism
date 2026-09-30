package proteintranslation

import (
    "errors"
)

var m = map[string]string{
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

var ErrStop = errors.New("ErrStop")
var ErrInvalidBase = errors.New("ErrInvalidBase")

func FromRNA(rna string) ([]string, error) {
    res := []string{}
	for i := 0; i < len(rna); i += 3 {
        end := i+3
        if end > len(rna) {
            end = len(rna)
        }
        codon := rna[i:end]
        acid, err := FromCodon(codon)
        if err == ErrStop {
            return res, nil
        } else if err == ErrInvalidBase {
            return []string{}, ErrInvalidBase
        } else {
            res = append(res, acid)
        }
    }
    return res, nil
}

func FromCodon(codon string) (string, error) {
	acid, exists := m[codon]
    if !exists {
        return "", ErrInvalidBase
    } else if acid == "STOP" {
        return "", ErrStop
    }
    return acid, nil
}
