class HandShake {
    private decimal: number
    private handshakes: string[] = []
    private messageMap: Map<number, ()=>void> = new Map([
        [1, ():void=>{this.handshakes.push("wink")}],
        [2, ():void=>{this.handshakes.push("double blink")}],
        [3, ():void=>{this.handshakes.push("close your eyes")}],
        [4, ():void=>{this.handshakes.push("jump")}],
        [5, ():void=>{this.handshakes.reverse()}]
    ]);

    constructor(decimal: number) {
        this.decimal = decimal
    }

    public commands(): string[] {
        const binaryArray = (this.decimal >>> 0).toString(2).split("").reverse();
        binaryArray.forEach((item, index) => {
            if (item === '1') {
                this.messageMap.get(index+1)?.call(this)
            }
        })
        return this.handshakes;
    }
}

export default HandShake;