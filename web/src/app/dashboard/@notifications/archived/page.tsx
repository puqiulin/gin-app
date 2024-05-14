import Card from "@/src/app/ui/card";
import Link from "next/link";

export default function Notifications() {
    return <Card>
        Archived
        <Link href={"/dashboard"}>Default</Link>
    </Card>
}
