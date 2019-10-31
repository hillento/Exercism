# Currently a work in progressgit
Codons = {
    "AUG" :  "Methionine",
    "UUU" : "Phenylalanine",
    "UUC" : "Phenylalanine",
    "UUA" : "Leucine",
    "UUG" : "Leucine",
    "UCU" : "Serine",
    "UCC" : "Serine",
    "UCA" : "Serine",
    "UCG" : "Serine",
    "UAU" : "Tyrosine",
    "UAC" : "Tyrosine",
    "UGU" : "Cysteine",
    "UGC" : "Cysteine",
    "UGG" : "Tryptophan",
    "UAA" : "STOP",
    "UAG" : "STOP",
    "UGA" : "STOP"

}

def proteins(strand):
    codonList = [[strand[i:i+3] for i in range(0, len(strand), 3)]]
    proteinList = []
    for codon in codonList:
        if codon in Codons:
            proteinList.append(Codons[codon])
    return proteinList

