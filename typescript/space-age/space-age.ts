type Planet = "Earth" | "Mercury" | "Venus" | "Mars" | "Jupiter" | "Saturn" | "Uranus" | "Neptune"

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
    private readonly ageInEarthYear: number
    private readonly EARTH_YEAR_IN_SECOND = 31557600

    constructor(ageInSecond: number) {
        this.seconds = ageInSecond
        this.ageInEarthYear = ageInSecond / this.EARTH_YEAR_IN_SECOND
    }

    private planetAge(planet: Planet): number {
        return round(this.ageInEarthYear / OrbitalPeriod[planet])
    }

    onEarth(): number {
        return this.planetAge("Earth")
    }

    onMercury(): number {
        return this.planetAge("Mercury")
    }

    onVenus(): number {
        return this.planetAge("Venus")
    }

    onMars(): number {
        return this.planetAge("Mars")
    }

    onJupiter(): number {
        return this.planetAge("Jupiter")
    }

    onSaturn(): number {
        return this.planetAge("Saturn")
    }

    onUranus(): number {
        return this.planetAge("Uranus")
    }

    onNeptune(): number {
        return this.planetAge("Neptune")
    }
}

export default SpaceAge;