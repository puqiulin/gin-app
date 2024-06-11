"use client"

import {useEffect, useState} from "react";

export default function CallBack() {
    const [loading, setLoading] = useState(true)
    const [userData, setUserData] = useState<any>()

    useEffect(() => {
        const hash = window.location.hash;
        const result = hash.split('&').reduce((r: any, i: string) => {
            const parts = i.split('=');
            r[parts[0]] = decodeURIComponent(parts[1]);
            return r
        }, {});
        if (result.access_token) {
            fetchUserInfo(result.access_token).then(() => {
                // alert("Login success!")
            }).catch(e => {
                alert(`Login error:${JSON.stringify(e)}`)
            });
        }
        setLoading(false)
    }, []);

    const fetchUserInfo = async (access_token: string) => {
        const response = await fetch('https://www.googleapis.com/oauth2/v2/userinfo', {
            headers: {
                Authorization: `Bearer ${access_token}`,
            },
        });
        const userInfo = await response.json();
        console.log(JSON.stringify(userInfo))
        setUserData(userInfo)
    }

    return <div className='w-screen h-screen flex justify-center items-center'>
        {loading ? <div>Loading...</div> :
            <div>
                <img src={userData?.picture} alt={'asd'}></img>
                <div>{userData?.email}</div>
                <div>{userData?.name}</div>
            </div>
        }
    </div>;
}
