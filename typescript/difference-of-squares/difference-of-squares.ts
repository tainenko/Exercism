export default class Squares {
    private readonly target: number
    squareOfSum: number
    sumOfSquares: number
    difference: number

    constructor(target: number) {
        this.target = target
        this.squareOfSum = this._squareOfSum()
        this.sumOfSquares = this._sumOfSquares()
        this.difference = Math.abs(this.sumOfSquares - this.squareOfSum)
    }

    private _squareOfSum(): number {
        let res = 0
        for (let i = 0; i <= this.target; i++) {
            res += i
        }
        return res * res
    }

    private _sumOfSquares(): number {
        let res = 0
        for (let i = 0; i <= this.target; i++) {
            res += i * i
        }
        return res

    }
}