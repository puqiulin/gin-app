"use client"

export const Componet2 = async () => {
    await new Promise(resolve => setTimeout(resolve, 3000))

    return <>
        <h1>Componet2</h1>
    </>
}
