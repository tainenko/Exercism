class Bob {

    hey(speech: string): string {
        speech = speech.trim()
        if (this.isEmpty(speech)) {
            return 'Fine. Be that way!'
        }
        if (this.isYellQuestion(speech)) {
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

    isEmpty(speech: string): boolean {
        return !speech.length
    }

    isQuestion(speech: string): boolean {
        return speech.endsWith("?")
    }

    isYelling(speech: string): boolean {
        return (/([A-Z])/.test(speech) && !/([a-z])/.test(speech))
    }

    isYellQuestion(speech: string): boolean {
        return this.isQuestion(speech) && this.isYelling(speech)
    }

}

export default Bob