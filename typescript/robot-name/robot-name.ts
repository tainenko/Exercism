const usedNames = new Set();

class RobotName {
    public name: string;

    constructor() {
        this.name = this.getRandomRobotName();
    }

    resetName(): void {
        this.name = this.getRandomRobotName()
    }

    private getRandomLetter(): string {
        let text = ""
        const possible = "ABCDEFGHIKLMNOPQRSTUVWXYZ"
        for (let i = 0; i < 2; i++) {
            text += possible.charAt(Math.floor(Math.random() * possible.length));
        }
        return text

    }

    private getRandomDigits(): string {
        return Math.floor(Math.random() * 1000).toString().padStart(3, '0')
    }

    private getRandomRobotName(): string {
        let name = this.getRandomLetter() + this.getRandomDigits()
        while (usedNames.has(name)) {
            name = this.getRandomLetter() + this.getRandomDigits()
        }
        usedNames.add(name)
        return name
    }
}

export default RobotName;