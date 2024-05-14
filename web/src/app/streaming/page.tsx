"use client"
import {Componet1} from "@/src/app/ui/streaming/componet1";
import {Componet2} from "@/src/app/ui/streaming/componet2";
import {Suspense} from "react";

export default function Streaming() {
    return <>
        <h1>Streaming</h1>
        <Suspense fallback={"Loading Componet1..."}>
            <Componet1/>
        </Suspense>
        <Suspense fallback={"Loading Componet2..."}>
            <Componet2/>
        </Suspense>
    </>
}
