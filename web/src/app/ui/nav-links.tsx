'use client'

import {usePathname} from 'next/navigation'
import Link from 'next/link'
import '@/src/app/ui/global.css'

export function NavLinks() {
    const pathname = usePathname()
    const navLink = [
        {name: "Home", href: "/"},
        {name: "Login", href: "/login"},
        {name: "Register", href: "/login"},
        {name: "Forgot password", href: "/forgot-password"},
        {name: "Blog", href: "/blog"},
        {name: "Docs", href: "/docs"},
    ]

    return (
        <nav>
            <ul>
                {navLink.map((link) => {
                    const isActive = pathname.startsWith(link.href)
                    console.log(isActive)
                    return <li key={link.name}>
                        <Link href={link.href} className={isActive ? "font-bold mr-4" : "text-blue-500 m4-4"}>
                            {link.name}
                        </Link>
                    </li>
                })}
            </ul>
        </nav>
    )
}