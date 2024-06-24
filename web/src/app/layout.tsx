"use client"
import React from "react";
import "@/src/app/global.css"
import bg from "@/public/images/grassland.png"
import {ApolloClient, ApolloProvider, InMemoryCache} from "@apollo/client";
import {useGlobalStore} from "@/src/stores/global-store";
import {ToDoListStoreProvider} from "@/src/providers/todo-list-store-provider";

const client = new ApolloClient({
    uri: 'http://127.0.0.1:9999/api/graphql',
    cache: new InMemoryCache(),
});

export default function RootLayout({
                                       children,
                                   }: Readonly<{
    children: React.ReactNode;
}>) {

    const theme = useGlobalStore((store) => store.theme)

    return (
        <html lang="en" data-theme={theme}>
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
                <ToDoListStoreProvider>
                    {/*<NavLinks/>*/}
                    {children}
                </ToDoListStoreProvider>
            </ApolloProvider>
        </main>
        </body>
        </html>
    );
}
