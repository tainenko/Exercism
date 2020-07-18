export default class Gigasecond {
    private currentDate: Date

    constructor(date: Date) {
        this.currentDate = date
    }

    date(): Date {
        return new Date(this.currentDate.valueOf() + 1000000000000)
    }

}