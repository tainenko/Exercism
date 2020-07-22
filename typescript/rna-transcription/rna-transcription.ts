const dna2Rna = {
    "G": "C",
    "C": "G",
    "T": "A",
    "A": "U"
}
type DNA = keyof typeof dna2Rna

class Transcriptor {
    toRna(dna: string): string {
        let res = ""
        for (let i = 0; i < dna.length; i++) {
            let rna = dna2Rna[dna[i] as DNA] || this.raiseInvalidInputError()
            res += rna
        }
        return res
    }

    raiseInvalidInputError() {
        throw new Error("Invalid input DNA.")
    }
}

export default Transcriptor