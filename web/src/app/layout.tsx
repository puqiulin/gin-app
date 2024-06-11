"use client"
import React from "react";
import "@/src/app/global.css"
import bg from "@/public/images/grassland.png"
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
    const bag2 = "https://picsum.photos/200/300"

    return (
        <html lang="en">
        <meta name="referrer" content="no-referrer-when-downgrade"/>
        <body className="
        {/*bg-gradient-to-br*/}
        {/*from-sky-400*/}
        {/*to-indigo-500*/}
        h-screen
        bg-no-repeat
        bg-cover
        bg-center
        backdrop-blur-md
        "
              style={{backgroundImage: `url(${bg.src})`}}
        >
        <main>
            <ApolloProvider client={client}>
                {/*<NavLinks/>*/}
                {children}
            </ApolloProvider>
        </main>
        </body>
        </html>
    );
}
