import Image from "next/image";
import head from "@/public/images/head.jpg"
import Menu from "@/src/app/ui/feishu/menu";
import {AiOutlineMessage, AiOutlineSetting} from "react-icons/ai"

export default function Root() {


    return <div className="
    h-screen
    bg-black/50
    backdrop-blur-sm
    text-white
    p-2
    ">
        <div className="grid grid-cols-10 gap-2 h-full">
            <div className="flex flex-col justify-between">
                <div className="flex flex-col items-center gap-6">
                    <Image className="rounded-full w-14" src={head} alt={"head"}></Image>
                    <Menu title={"Message"} icon={<AiOutlineMessage/>}/>
                    <Menu title={"Message"} icon={<AiOutlineMessage/>}/>
                    <Menu title={"Message"} icon={<AiOutlineMessage/>}/>
                    <Menu title={"Message"} icon={<AiOutlineMessage/>}/>
                    <Menu title={"Message"} icon={<AiOutlineMessage/>}/>
                </div>
                <div>
                    <Menu title={"Message"} icon={<AiOutlineSetting/>}/>
                </div>
            </div>
            <div className="content col-span-3">
                content1
            </div>
            <div className="content col-span-6">
                content2
            </div>
        </div>
    </div>
}
