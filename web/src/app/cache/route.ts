// force update the get cache
export const dynamic = "force-dynamic"

//get all comments
export async function GET() {
    return Response.json({
        time: new Date().toLocaleTimeString()
    })
}
