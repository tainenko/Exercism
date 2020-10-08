const usedNames = new Set();

class RobotName {
    public name: string;

    constructor() {
        this.name = this.getRobotName();
    }

    resetName(): void {
        this.name = this.getRobotName()
    }

    private getRobotName(): string {
        let name: string
        do {
            name = this.getRandomPrefixLetters() + this.getRandom3Digits()
        } while (usedNames.has(name))
        usedNames.add(name)
        return name
    }

    private getRandomPrefixLetters(): string {
        let upperLetter = ""
        const possible = "ABCDEFGHIKLMNOPQRSTUVWXYZ"
        for (let i = 0; i < 2; i++) {
            upperLetter += possible.charAt(Math.floor(Math.random() * possible.length));
        }
        return upperLetter

    }

    private getRandom3Digits(): string {
        return Math.floor(Math.random() * 1000).toString().padStart(3, '0')
    }
}

export default RobotName;