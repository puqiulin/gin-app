import {notFound} from "next/navigation";

export default function review({params}: { params: { blogId: string, commentId: string } }) {
    if (parseInt(params.commentId) > 1000) {
        notFound()
    }
    return <h1>Comment {params.commentId} for Blog {params.blogId}</h1>
}
