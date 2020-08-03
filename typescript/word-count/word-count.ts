class Words {
    count(input: string): Map<string, number> {
        let words = input.trim().toLowerCase().split(/\s+/);
        return words.reduce((map, word) => map.set(word, map.get(word) + 1 || 1), new Map())
    }
}

export default Words