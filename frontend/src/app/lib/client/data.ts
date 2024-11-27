import { User, Roles } from "../definitions";

export const passwordRegex = /^[a-zA-Z0-9!@#$%^&*()_+=\[\]{};':"\\|,.<>\/?~-]{8,30}$/;

export async function getUserStuff() {
    try {

    } catch(err) {
        
    }
}

export async function getUser(username: string, userType: string): Promise<User> {
    try {
        const userPath = resolveRouteFromRole(userType);
        const url = `http://localhost:8002/api/${userPath}/${username}`
        const resp = await fetch(url, {
            method: "GET"
        })

        if (resp.ok) {
            const user: User = await resp.json();
            return user;
        } else {
            throw new Error("Request to get user information was not successful")
        }
    } catch(err) {
       console.error("Error retrieving user information")
       throw new Error("Error retrieving user information")
    }
}

function resolveRouteFromRole(role: string) {
    switch (role) {
        case Roles.STUDENT:
            return "students"
        case Roles.TEACHER:
            return "teachers"
        case Roles.ADMIN:
            return "administrators"
        default:
            throw new Error("Unable to resolve route from role")
    }
}
