type NucleotidesOfDna = "G" | "C" | "T" | "A"
type NucleotidesOfRna = "C" | "G" | "A" | "U"
const NucleotidesMap: Record<string, NucleotidesOfRna> = {
    "G": "C",
    "C": "G",
    "T": "A",
    "A": "U"
}

class Transcriptor {
    toRna = (dna: string): string => {
        return dna.split('').map(this.convertNucleotide).join('')
    }

    convertNucleotide = (strand: string): NucleotidesOfRna => {
        if (!this.isNucleotidesOfDna(strand)) {
            this.raiseInvalidInputError()
        }
        return NucleotidesMap[strand]
    }

    isNucleotidesOfDna = (strand: string): strand is NucleotidesOfDna => {
        return NucleotidesMap[strand] !== undefined
    }

    raiseInvalidInputError = (): void => {
        throw new Error("Invalid input DNA.")
    }
}

export default Transcriptor