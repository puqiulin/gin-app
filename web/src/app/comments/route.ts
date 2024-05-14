import {comments} from "@/src/app/lib/data"
import {NextRequest} from "next/server.js";

//get all comments
export async function GET(request: NextRequest) {
    const requestHeader = request.headers
    console.log(requestHeader)

    const query = request.nextUrl.searchParams.get("query")
    const comment = comments.filter(c => c.comment.includes(query ? query : ""))
    return Response.json(comment)
}

//add new comment
export async function POST(request: Request) {
    const comment = await request.json()
    const newComment = {
        id: comments.length + 1,
        comment: comment.text
    }
    comments.push(newComment)
    return new Response(JSON.stringify(newComment), {
        headers: {
            "Content-Type": "application/json",
        },
        status: 201
    })
}