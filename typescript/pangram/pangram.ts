class Pangram {
    private readonly alphabetLength = 26
    private readonly sentence: string

    constructor(sentence: string) {
        this.sentence = sentence
    }

    isPangram(): boolean {
        const alphaSet = new Set(this.sentence.toLowerCase().replace(/[^a-z]/g, ""))
        return alphaSet.size == this.alphabetLength
    }
}

export default Pangram