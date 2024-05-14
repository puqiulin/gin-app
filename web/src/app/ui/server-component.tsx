import fs from "fs"

export default function ServerComponent() {
    fs.readFileSync("src/app/global.css", "utf-8")
    return <>
        <h1>Server Component</h1>
    </>
}