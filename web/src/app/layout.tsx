"use client"
import React from "react";
import {NavLinks} from "@/src/app/ui/nav-links";
import "@/src/app/global.css"
import {ApolloClient, ApolloProvider, InMemoryCache} from "@apollo/client";

const client = new ApolloClient({
    uri: 'http://127.0.0.1:9999/api/graphql',
    cache: new InMemoryCache(),
});

export default function RootLayout({
                                       children,
                                   }: Readonly<{
    children: React.ReactNode;
}>) {


    return (
        <html lang="en">
        <body className="
        bg-gradient-to-br from-green-400 via-blue-500 to-purple-600 h-64 w-full">
        <main>
            <ApolloProvider client={client}>
                <NavLinks/>
                {children}
            </ApolloProvider>
        </main>
        </body>
        </html>
    );
}
