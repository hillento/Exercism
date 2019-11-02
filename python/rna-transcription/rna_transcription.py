translate = {
    "G": "C",
    "C": "G",
    "T": "A",
    "A": "U"
}

def to_rna(dna_strand):
    lst = [x for x in dna_strand]
    lst2 = []
    rna_strand = ""
    for dna in lst:
        lst2.append(translate[dna])
    return rna_strand.join(lst2)
