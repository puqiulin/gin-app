"use client"
import Card from "@/src/app/ui/card";
import Loading from "./loading";
import {Suspense} from "react";

export default function UserAnalytics() {

    return <div>
        <Suspense fallback={<Loading/>}>
            <Card>UserAnalytics</Card>
        </Suspense>
    </div>
}
