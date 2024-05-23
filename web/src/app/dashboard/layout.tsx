import React from "react";
import Login from "@/src/app/dashboard/@login/page";

export default function DashboardLayout({children, login, user, revenue, notifications}: {
    children: React.ReactNode,
    user: React.ReactNode,
    login: React.ReactNode,
    revenue: React.ReactNode,
    notifications: React.ReactNode
}) {
    const isLogin = true

    return isLogin ? <div>
        <div>{children}</div>
        <div className="grid grid-rows-2 grid-flow-col gap-2">
            <div className="row-span-2 col-span-1">{user}</div>
            <div className="row-span-1 col-span-1 ">{revenue}</div>
            <div className="row-span-1 col-span-1 ">{notifications}</div>
        </div>
    </div> : <Login/>
}