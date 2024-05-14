import React from "react";
import image1 from "@/public/images/wallhaven-0jxxgm_3840x3072.png"
import image2 from "@/public/images/wallhaven-jxoxjm_3840x3072.png"
import image3 from "@/public/images/wallhaven-3k62g3_3840x3072.png"
import image4 from "@/public/images/wallhaven-47xvr9_3840x3072.png"
import image5 from "@/public/images/wallhaven-nzmeog_3840x3072.png"
import Image from 'next/image'

export default function Tailwind() {
    return <div className="h-screen grid grid-cols-4">
        <div className="flex bg-amber-100 md:col-span-1 md:justify-end">
            <div className="text-right mr-1">
                <ul>
                    <li>
                        <h1 className="text-gray-600 text-4xl font-bold border-red-300 font-poetsen">tailwind</h1>
                    </li>
                    <li>
                        <a className="border-r-4 border-red-300">home</a>
                    </li>
                    <li>
                        <a>about</a>
                    </li>
                </ul>
            </div>
        </div>

        <div className="bg-amber-600 col-span-3">
            <div className="flex justify-center md:justify-end">
                <div className="m-2 btn border-2 ">login</div>
                <div className="m-2 btn border-2 ">register</div>
            </div>

            <div
                className="h-56 w-56
                  sm:bg-cyan-200 md:bg-cyan-400 lg:bg-cyan-600
                  flex justify-center items-center
                  text-xs
                  sm:text-xs md:text-lg lg:text-3xl
                  ">
                response div
            </div>

            <div className="m-2 grid gap-2 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
                <div className="card">
                    <Image src={image1} alt={"one day"} className="h-36 w-full object-cover"></Image>
                    <div className="m-2">
                        <span className="font-bold">one day</span>
                        <span className="block text-gray-600 text-xs">i will make it</span>
                    </div>
                    <div
                        className="badge">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                             stroke="currentColor" className="w-5 h-5 inline-block">
                            <path stroke-linecap="round" stroke-linejoin="round"
                                  d="M14.857 17.082a23.848 23.848 0 0 0 5.454-1.31A8.967 8.967 0 0 1 18 9.75V9A6 6 0 0 0 6 9v.75a8.967 8.967 0 0 1-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 0 1-5.714 0m5.714 0a3 3 0 1 1-5.714 0"/>
                        </svg>

                        puqiuqlin
                    </div>
                </div>

                <div className="card">
                    <Image src={image2} alt={"one day"} className="h-36 w-full object-cover"></Image>
                    <div className="m-2">
                        <span className="font-bold">one day</span>
                        <span className="block text-gray-600 text-xs">i will make it</span>
                    </div>
                    <div
                        className="badge">
                        puqiuqlin
                    </div>
                </div>

                <div className="card">
                    <Image src={image3} alt={"one day"} className="h-36 w-full object-cover"></Image>
                    <div className="m-2">
                        <span className="font-bold">one day</span>
                        <span className="block text-gray-600 text-xs">i will make it</span>
                    </div>
                    <div
                        className="badge">
                        puqiuqlin
                    </div>
                </div>

                <div className="card">
                    <Image src={image4} alt={"one day"} className="h-36 w-full object-cover"></Image>
                    <div className="m-2">
                        <span className="font-bold">one day</span>
                        <span className="block text-gray-600 text-xs">i will make it</span>
                    </div>
                    <div
                        className="badge">
                        puqiuqlin
                    </div>
                </div>

                <div className="card sm:col-span-1 md:col-span-2 lg:col-span-3">
                    <Image src={image5} alt={"one day"} className="h-36 w-full object-cover"></Image>
                    <div className="m-2">
                        <span className="font-bold">one day</span>
                        <span className="block text-gray-600 text-xs">i will make it</span>
                    </div>
                    <div
                        className="badge">
                        puqiuqlin
                    </div>
                </div>
            </div>

            <div className="flex justify-center m-2">
                <div className="btn bg-emerald-500">load more</div>
            </div>

        </div>

    </div>
}