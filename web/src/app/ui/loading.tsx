'use client'

import loadingCat from '@/public/lotties/loading-cat.json'
import Lottie from "lottie-react";

export default function Loading() {
    const defaultOptions = {
        loop: true,
        autoplay: true,
        animationData: loadingCat,
        rendererSettings: {
            preserveAspectRatio: 'xMidYMid slice',
        },
    };

    return <Lottie {...defaultOptions} className="lottie"/>
}