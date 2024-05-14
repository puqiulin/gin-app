import {NextRequest, NextResponse} from "next/server.js";

export function middleware(request: NextRequest) {
    const response = NextResponse.next()

    const theme = request.cookies.get("theme")
    if (!theme) {
        response.cookies.set("theme", "dark")
    }

    response.headers.set("custom-header", "custom-value")

    return response
}