import {create} from "zustand";

type CounterStore = {
    count: number
    add: () => void;
    addAsync: () => Promise<void>;
    red: () => void;
}

export const useCounterStore = create<CounterStore>((set: any) => ({
    count: 0,
    add: () => {
        set((state: any) => ({
            count: state.count + 1
        }))
    },
    addAsync: async () => {
        await new Promise((resolve, reject) => setTimeout(resolve, 1000))
        set((state: any) => ({
            count: state.count + 1
        }))
    },
    red: () => {
        set((state: any) => ({
            count: state.count - 1
        }))
    }
}))