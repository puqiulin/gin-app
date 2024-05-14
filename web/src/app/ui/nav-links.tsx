'use client'

import {usePathname} from 'next/navigation'
import Link from 'next/link'

// import '@/src/app/ui/globals.css'

export function NavLinks() {
    const pathname = usePathname()
    const navLink = [
        {name: "Home", href: "/"},
        {name: "Dashboard", href: "/dashboard"},
        {name: "Login", href: "/login"},
        {name: "Register", href: "/register"},
        {name: "Forgot password", href: "/forgot-password"},
        {name: "Blog", href: "/blog"},
        {name: "Docs", href: "/docs"},
        {name: "Order", href: "/order"},
        {name: "Hello", href: "/hello"},
        {name: "Comments", href: "/comments"},
        {name: "Custom", href: "/custom"},
        {name: "Cache", href: "/cache"},
        {name: "Streaming", href: "/streaming"},
        {name: "Interleaving", href: "/interleaving"},
        {name: "Tailwind", href: "/tailwind"},
    ]

    return (
        <div className="flex flex-row gap-2">
            {navLink.map((link) => {
                const isActive = pathname.startsWith(link.href)
                return <Link href={link.href} key={link.name}>
                    <div className={isActive ? "text-teal-300" : ""}>{link.name}</div>
                </Link>
            })}
        </div>
    )
}