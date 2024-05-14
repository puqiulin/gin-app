import Card from "@/src/app/ui/card";
import Link from "next/link";

export default function DefaultNotifications() {
    return <Card>
        Notifications
        <Link href={"/dashboard/archived"}>Archived</Link>
    </Card>
}
