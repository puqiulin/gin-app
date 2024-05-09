import React from "react";
import {NavLinks} from "@/src/app/ui/nav-links";

export const metadata = {
    title: {
        absolute: "Absolute title",
        default: "Default title",
        template: "Template title | %s",
    },
    description: "oneday"
}

export default function RootLayout({
                                       children,
                                   }: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang="en">
        <body>
        <main>
            <NavLinks/>
            {children}
        </main>
        </body>
        </html>
    );
}
