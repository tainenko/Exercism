const color2Resistor = {
    "black": 0,
    "brown": 1,
    "red": 2,
    "orange": 3,
    "yellow": 4,
    "green": 5,
    "blue": 6,
    "violet": 7,
    "grey": 8,
    "white": 9
}
type Color = keyof typeof color2Resistor

export class ResistorColor {
    private colors: Color[];

    constructor(colors: Color[]) {
        if (colors.length < 2) {
            throw new Error('At least two colors need to be present');
        }
        this.colors = colors;
    }

    value = (): number => {
        const [color0, color1] = this.colors
        return color2Resistor[color0] * 10 + color2Resistor[color1]
    };
}
