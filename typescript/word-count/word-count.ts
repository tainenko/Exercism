class Words {
    count(input: string): Map<string, number> {
        let words = input.trim().toLowerCase().split(/\s+/);
        return words.reduce((map, index) => map.set(index, map.get(index) + 1 || 1), new Map())
    }
}

export default Words