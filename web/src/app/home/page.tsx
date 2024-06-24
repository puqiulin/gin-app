"use client"

import {useCounterStore} from "@/src/stores/store";

export default function Home() {
    const count = useCounterStore((state) => state.count)
    const add = useCounterStore((state) => state.add)
    const addAsync = useCounterStore((state) => state.addAsync)
    const red = useCounterStore((state) => state.red)

    return <div className="info-card" onClick={() => {
    }}>
        <div className="flex gap-2 justify-center items-center">
            <div className='text-5xl'>😎</div>
            <div className="flex flex-col gap-2 items-start">
                <div className='flex gap-1 place-items-baseline'>
                    <div className="text-5xl cursor-pointer" onClick={red}>-</div>
                    <div className="text-5xl font-bold">{count}</div>
                    <div className="font-bold">GB</div>
                    <div className="text-5xl cursor-pointer" onClick={add}>+</div>
                    <div className="text-5xl cursor-pointer" onClick={addAsync}>p+</div>
                </div>
                <div> information</div>
            </div>
        </div>
    </div>
}
