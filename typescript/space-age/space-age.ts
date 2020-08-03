enum OrbitalPeriod {
    Earth = 1,
    Mercury = 0.2408467,
    Venus = 0.61519726,
    Mars = 1.8808158,
    Jupiter = 11.862615,
    Saturn = 29.447498,
    Uranus = 84.016846,
    Neptune = 164.79132,


}

const round = (input: number): number => {
    return Math.round(input * 100) / 100
}

class SpaceAge {
    public seconds: number
    private readonly earthYear: number
    private readonly earthYearInSecond: number

    constructor(seconds: number) {
        this.seconds = seconds
        this.earthYearInSecond = 31557600
        this.earthYear = seconds / this.earthYearInSecond
    }

    private planetAge(planetPeriod: OrbitalPeriod): number {
        return round(this.earthYear / planetPeriod)
    }

    onEarth(): number {
        return this.planetAge(OrbitalPeriod.Earth)
    }

    onMercury(): number {
        return this.planetAge(OrbitalPeriod.Mercury)
    }

    onVenus() {
        return this.planetAge(OrbitalPeriod.Venus)
    }

    onMars() {
        return this.planetAge(OrbitalPeriod.Mars)
    }

    onJupiter() {
        return this.planetAge(OrbitalPeriod.Jupiter)
    }

    onSaturn() {
        return this.planetAge(OrbitalPeriod.Saturn)
    }

    onUranus() {
        return this.planetAge(OrbitalPeriod.Uranus)

    }

    onNeptune() {
        return this.planetAge(OrbitalPeriod.Neptune)
    }


}

export default SpaceAge;