"use client"
import Card from "@/src/app/ui/card";
import {Suspense, useEffect, useState} from "react";
import Loading from "@/src/app/dashboard/@user/loading";

export default function UserAnalytics() {

    return <Suspense fallback={<Loading/>}><Card>UserAnalytics</Card></Suspense>
}
