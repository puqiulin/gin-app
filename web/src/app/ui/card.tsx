import React from "react";

export default function Card({children}: { children: React.ReactNode }) {
    return <div className="max-w-lg rounded-2xl overflow-hidden shadow-lg bg-cyan-300">
        <div className="m-10">{children}</div>
    </div>
}