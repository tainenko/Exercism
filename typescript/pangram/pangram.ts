class Pangram {
    private input: string

    constructor(input: string) {
        this.input = input
    }

    isPangram(): boolean {
        const res = new Set()
        this.input.toLowerCase().replace(/./g, c => {
            if (c >= 'a' && c <= 'z') {
                res.add(c)
            }
            return c
        })
        return res.size == 26
    }
}

export default Pangram