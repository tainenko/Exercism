const MinuteInADay = 1440

class Clock {
    private minute: number
    constructor(hour: number, minute: number = 0) {
        this.minute = 60 * hour + minute
    }

    public toString(): string {
        this.correct()
        const hour = String(Math.floor(this.minute / 60)).padStart(2, '0')
        const minute = String(this.minute % 60).padStart(2, '0')
        return `${hour}:${minute}`
    }

    public plus(minute: number): this {
        this.minute += minute
        return this
    }

    public minus(minute: number): this {
        this.minute -= minute
        return this
    }

    public equals(clock: Clock): boolean {
        return this.toString() === clock.toString()
    }

    private correct(): void {
        if (this.minute >= 0 && this.minute < MinuteInADay) {
            return
        }
        this.minute %= MinuteInADay
        if (this.minute < 0) {
            this.minute += MinuteInADay
        }
    }
}

export default Clock;