import {comments} from "@/src/app/lib/data"
import {redirect} from "next/navigation";
import {NextRequest} from "next/server.js";

export async function GET(request: NextRequest, {params}: { params: { id: string } }) {
    if (parseInt(params.id) > comments.length) {
        redirect("/comments")
    }
    const comment = comments.find(c => c.id === parseInt(params.id))
    return Response.json(comment)
}

export async function PATCH(request: Request, {params}: { params: { id: string } }) {
    const newComment = await request.json()
    const index = comments.findIndex(c => c.id === parseInt(params.id))
    comments[index].comment = newComment.text
    return Response.json(comments[index])
}

export async function DELETE(request: Request, {params}: { params: { id: string } }) {
    const index = comments.findIndex(c => c.id === parseInt(params.id))
    const deleteComment = comments[index]
    comments.splice(index, 1)
    return Response.json(deleteComment)
}