import loadingCat from "@/public/lotties/loading-cat.json";

export default function Loading() {

    const defaultOptions = {
        loop: true,
        autoplay: true,
        animationData: loadingCat,
        rendererSettings: {
            preserveAspectRatio: 'xMidYMid slice',
        },
    };

    return <div>
        {/*<Lottie {...defaultOptions} />*/}
        Loading...
    </div>
}