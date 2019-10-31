# Currently a work in progressgit
Codon2Protein = {
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
    # uses slicing to split the string into strings of 3 characters.
    codonList = [strand[i:i+3] for i in range(0, len(strand), 3)]
    #initialize an empty list 
    proteinList = []
    #for each item in the codonList....
    for codon in codonList:
        #...and it it translates to stop in the dictionart break -> return proteinList
        if Codon2Protein[codon] == "STOP" : break
        #...translate the value of the codon from the dictionary. then append it to the proteinList
        proteinList.append(Codon2Protein[codon])
    return proteinList




