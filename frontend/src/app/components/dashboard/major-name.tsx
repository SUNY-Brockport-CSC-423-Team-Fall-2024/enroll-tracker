import { getStudentMajor } from "@/app/lib/server/data"

export async function MajorName() {
    const major = await getStudentMajor();
    return (
        <>
            {major && <p>{major.name}</p>}
            {!major && <p>No major.</p>}
        </>
    )
}
