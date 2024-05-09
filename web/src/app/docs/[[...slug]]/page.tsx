import Link from "next/link";

export default function Docs({params}: { params: { slug: string[] } }) {
    if (params.slug?.length === 2) {
        return <h1>Feature {params.slug[0]}, concept {params.slug[1]}</h1>
    } else if (params.slug?.length === 1) {
        return <h1>Feature {params.slug[0]}</h1>
    }

    return <div>
        <h1>Docs</h1>
        <ul>
            <li>
                <Link href={"/docs/feature1"}>feature1</Link>
                <ul>
                    <li>
                        <Link href={"/docs/feature1/concept1"}>concept1</Link>
                    </li>
                    <li>
                        <Link href={"/docs/feature1/concept2"}>concept2</Link>
                    </li>
                </ul>
            </li>
            <li>
                <Link href={"/docs/feature2"}>feature2</Link>
            </li>
            <li>
                <Link href={"/docs/feature3"}>feature3</Link>
            </li>
        </ul>
    </div>

}
