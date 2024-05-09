import Link from "next/link";
import {Metadata} from "next";

type Props = {
    params: {
        blogId: string
    }
}
export const generateMetadata = ({params}: Props): Metadata => {
    return {
        title: `Blog ${params.blogId}`
    }
}

export default function BlogDetails({params}: { params: { blogId: string } }) {
    return <div>
        <h1>Blog {params.blogId}</h1>
        <ul>
            <li>
                <Link href={"/blog/1/comment/1"}>comment 1</Link>
            </li>
            <li>
                <Link href={"/blog/2/comment/2"}>comment 2</Link>
            </li>
        </ul>
    </div>
}
