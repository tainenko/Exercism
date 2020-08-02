type DnaNucleotide = "G" | "C" | "T" | "A"
type RnaNucleotide = "C" | "G" | "A" | "U"
const NucleotidesMap: Record<DnaNucleotide, RnaNucleotide> = {
    "G": "C",
    "C": "G",
    "T": "A",
    "A": "U"
}

class Transcriptor {
    toRna = (dna: string): string => {
        return dna.replace(/./g, this.convertNucleotide)
    }

    convertNucleotide = (strand: string): RnaNucleotide => {
        if (this.isNucleotidesOfDna(strand)) {
            return NucleotidesMap[strand]
        } else {
            throw new Error("Invalid input DNA.")
        }
    }

    isNucleotidesOfDna = (strand: string): strand is DnaNucleotide => {
        return strand in NucleotidesMap
    }
}

export default Transcriptor