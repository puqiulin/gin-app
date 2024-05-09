import Link from "next/link";

export default function Blog() {
    return <div>
        <h1>Blog</h1>
        <ul>
            <li>
                <Link href={"/blog/1"}>blog1</Link>
            </li>
            <li>
                <Link href={"/blog/2"}>blog2</Link>
            </li>
            <li>
                <Link href={"/blog/3"}>blog3</Link>
            </li>
        </ul>
    </div>
}
