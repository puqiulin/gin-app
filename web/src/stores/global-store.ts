import {create} from "zustand";

type globalStore = {
    theme: string
    toggleTheme: () => void
}

export const useGlobalStore = create<globalStore>((set) => ({
    theme: "dark",
    toggleTheme: () => set(store => ({
        theme: store.theme === "light" ? "dark" : "light"
    }))
}))
