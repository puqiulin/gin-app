import {comments} from "@/src/app/lib/data"
import {NextRequest} from "next/server.js";
import {cookies} from "next/headers";

//get all comments
export async function GET(request: NextRequest) {
    const requestHeader = request.headers
    console.log(requestHeader)
    console.log(request.cookies.get("xxx"))
    console.log(cookies().get("xxx"))

    return new Response("<h1>Hello world</h1>", {
        headers: {
            "Content-Type": "text/html"
        }
    })
}
