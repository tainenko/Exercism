class Transcriptor {
    toRna(dna: string): string {
        return dna.split('').map(this.convertNucleotide).join('')
    }

    convertNucleotide(strand: string): string {
        switch (strand) {
            case "G":
                return "C"
            case "C":
                return "G"
            case "T":
                return "A"
            case "A":
                return "U"
            default:
                throw new Error("Invalid input DNA.")

        }


    }
}

export default Transcriptor