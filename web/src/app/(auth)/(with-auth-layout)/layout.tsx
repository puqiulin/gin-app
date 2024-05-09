import React from "react";


export default function AuthLayout({children}: { children: React.ReactNode }) {
    return <div>
        <h1>Auth layout</h1>
        {children}
    </div>
}