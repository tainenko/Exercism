class GradeSchool {
    private student2Grade = new Map<string, number>()

    studentRoster(): Map<string, string[]> {
        const roster = new Map<string, string[]>()
        this.student2Grade.forEach((value, key) => {
            roster.set(value.toString(), [...roster.get(value.toString()) || [], key])
        })
        roster.forEach((value) => {
            value.sort()
        })
        return roster

    }

    addStudent(student: string, grade: number) {
        this.student2Grade.set(student, grade)
    }

    studentsInGrade(grade: number): string[] {
        const students: string[] = []
        this.student2Grade.forEach((value, key) => {
            if (grade == value) {
                students.push(key)
            }
        })
        return students.sort()
    }

}

export default GradeSchool;