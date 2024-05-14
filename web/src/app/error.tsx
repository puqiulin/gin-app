"use client"
export default function Error({error, reset}: { error: Error, reset: () => void }) {
    return <>
        <h1>Error: {error.message}</h1>
        <button onClick={reset}>try again</button>
    </>
}
