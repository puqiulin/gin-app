"use client"

import {Column} from "@/src/app/ui/todolist/column";
import {useGlobalStore} from "@/src/stores/global-store";
import {useEffect} from "react";
import {ToDoListStoreProvider, useToDoListStore} from "@/src/providers/todo-list-store-provider";

export default function Zustand() {
    const {theme, toggleTheme} = useGlobalStore()
    const tasks = useToDoListStore(store => store.tasks)
    const draggedTask = useToDoListStore((store) => store.draggedTask)

    useEffect(() => {
        console.log(tasks)
    }, [tasks])

    return (
        <ToDoListStoreProvider>
            <div className="w-screen h-screen flex flex-col bg-base-100 p-10">
                <div className="flex gap-2 mb-3">
                    <input type="checkbox" className="toggle" onChange={toggleTheme}/>
                    <div>Dragged item: <span className="font-bold">{draggedTask?.title}</span></div>
                </div>
                <div className="grid grid-cols-3 gap-4">
                    <Column state={"Todo"}/>
                    <Column state={"Doing"}/>
                    <Column state={"Done"}/>
                </div>
            </div>
        </ToDoListStoreProvider>
    )
}
