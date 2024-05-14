"use client"
import React, {useState} from "react";
import ServerComponent from "@/src/app/ui/server-component";

export default function ClientComponent({children}: { children: React.ReactNode }) {
    const [name, setName] = useState("wangjie")

    return <>
        <h1>Client Component</h1>

        {/*can not direct use server component like this*/}
        {/*<ServerComponent></ServerComponent>*/}

        {/*pass it as a prop to the client component*/}
        {children}
    </>
}