"use client"
import React, {useState} from "react";


export default function AuthLayout({children}: { children: React.ReactNode }) {
    const [input, setInput] = useState("")

    return <div>
        <h1>Auth layout</h1>
        <input value={input} onChange={(e) => setInput(e.target.value)}></input>
        {children}
    </div>
}