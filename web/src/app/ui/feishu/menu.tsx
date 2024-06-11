"use client"
import React from "react";

export default function Menu({title, icon}: { title: string, icon: React.ReactNode }) {
    return <div className="
    flex
    flex-col
    items-center
    p-2
    hover:bg-blue-400 rounded-2xl cursor-pointer
    transition esae-out duration-300
    ">
        <div className="text-3xl">{icon}</div>
        <div className="">{title}</div>
    </div>
}