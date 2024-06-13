"use client"
import React from "react";
import Login from "./@login/page"

export default function DashboardLayout({children, login, user, revenue, notifications}: {
    children: React.ReactNode,
    login: React.ReactNode,
    user: React.ReactNode,
    revenue: React.ReactNode,
    notifications: React.ReactNode
}) {
    const isLogin = true

    if (!isLogin) {
        return <Login/>
    }

    return <div>
        <div>{children}</div>
        <div className="grid grid-rows-2 grid-flow-col gap-2">
            <div className="row-span-2 col-span-1">{user}</div>
            <div className="row-span-1 col-span-1 ">{revenue}</div>
            <div className="row-span-1 col-span-1 ">{notifications}</div>
        </div>
    </div>
}