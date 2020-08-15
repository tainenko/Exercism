class Bob {

    hey(speech: string): string {
        if (this.isNothing(speech)) {
            return 'Fine. Be that way!'
        }
        if (this.isQuestion(speech) && this.isYelling(speech)) {
            return 'Calm down, I know what I\'m doing!'
        }
        if (this.isYelling(speech)) {
            return 'Whoa, chill out!'
        }
        if (this.isQuestion(speech)) {
            return 'Sure.'
        }
        return 'Whatever.'
    }

    private isQuestion(speech: string): boolean {
        return speech.trimEnd().endsWith("?")
    }

    private isYelling(speech: string): boolean {
        return /[A-Z]/.test(speech) && speech === speech.toUpperCase()
    }

    private isNothing(speech: string): boolean {
        return speech.length === 0 || /^\s*$/.test(speech)
    }
}

export default Bob