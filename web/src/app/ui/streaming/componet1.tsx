"use client"

export const Componet1 = async () => {
    await new Promise(resolve => setTimeout(resolve, 2000))
    return <>
        <h1>Componet1</h1>
    </>
}
