"use client"

import {useEffect, useState} from "react";
import {usePathname, useRouter, useSearchParams} from "next/navigation";

export default function GoogleLogin() {
    const clientID = "606833368399-hvkjir9r63q2vit2chfkvd16r1tu6h30.apps.googleusercontent.com"
    const redirectUri = "http://localhost:3001/google-login/call-back";
    const scope = "https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile";
    const googleAuthUrl = "https://accounts.google.com/o/oauth2/v2/auth"
    const options = {
        redirect_uri: redirectUri,
        client_id: clientID,
        scope: scope,
        prompt: "consent",
        include_granted_scopes: "true",
        state: "your-value",
        response_type: "token"
    }
    const params = new URLSearchParams(options)
    const authUrl = `${googleAuthUrl}?${params}`

    const [isLogin, setIsLogin] = useState(false)
    const [userData, setUserData] = useState(null)
    const router = useRouter();
    const searchParams = useSearchParams()
    const pathname = usePathname();

    const onClick = () => {
        window.location.href = authUrl
    }

    useEffect(() => {
        const access_token = searchParams.get("access_token")
        if (access_token) {
            alert(access_token)
        }
    }, []);


    return <div className="h-screen w-screen flex items-center justify-center">
        <button className="font-bold bg-white p-2 rounded hover:shadow-lg" onClick={onClick}>
            Login with Google
        </button>
    </div>
}
