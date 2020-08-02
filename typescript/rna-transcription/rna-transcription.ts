type DnaNucleotide = "G" | "C" | "T" | "A"
type RnaNucleotide = "C" | "G" | "A" | "U"
const NucleotidesMap: Record<DnaNucleotide, RnaNucleotide> = {
    "G": "C",
    "C": "G",
    "T": "A",
    "A": "U"
}

class Transcriptor {
    toRna = (dnaStrand: string): string => {
        return dnaStrand.replace(/./g, this.convertNucleotide)
    }

    convertNucleotide = (nucleotide: string): RnaNucleotide => {
        if (this.isDnaNucleotide(nucleotide)) {
            return NucleotidesMap[nucleotide]
        } else {
            throw new Error("Invalid input DNA.")
        }
    }

    isDnaNucleotide = (nucleotide: string): nucleotide is DnaNucleotide => {
        return nucleotide in NucleotidesMap
    }
}

export default Transcriptor