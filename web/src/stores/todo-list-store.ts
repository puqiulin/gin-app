import {createStore} from 'zustand/vanilla'
import {createJSONStorage, persist} from "zustand/middleware";

export interface Task {
    title: string
    state: string
}

export type ToDoListState = {
    tasks: Task[]
    draggedTask: Task | null
}

export type ToDoListActions = {
    addTask: (task: Task) => void
    deleteTask: (title: string) => void
    setDraggedTask: (task: Task | null) => void
    moveTask: (task: Task | null, newState: string) => void
}

export type ToDoListStore = ToDoListState & ToDoListActions

export const initToDoListStore = (): ToDoListState => {
    return {
        tasks: [
            {title: "wangjie", state: "Todo"},
            {title: "wangjiea", state: "Todo"},
            {title: "puqiulin", state: "Doing"},
            {title: "wangting", state: "Done"},
        ],
        draggedTask: null
    }
}

export const defaultInitState: ToDoListState = {
    tasks: [
        {title: "wangjie", state: "Todo"},
        {title: "wangjiea", state: "Todo"},
        {title: "puqiulin", state: "Doing"},
        {title: "wangting", state: "Done"},
    ],
    draggedTask: null
}

export const createToDoListStore = (
    initState: ToDoListState = defaultInitState,
) => {
    return createStore<ToDoListStore>()(persist((set) => ({
            ...initState,
            addTask: (task) => set((store) => ({
                tasks: [...store.tasks, task]
            })),
            deleteTask: (title: string) => set((store) => ({
                tasks: store.tasks.filter(t => t.title !== title)
            })),
            setDraggedTask: (task) => set((store) => ({
                draggedTask: task
            })),
            moveTask: (task, newState) => set((store) => ({
                tasks: task ? store.tasks.map(t => (t.title === task.title ? {
                    ...t,
                    state: newState
                } : t)) : store.tasks
            })),
        }),
        {
            name: "todo-list-store",
            storage: createJSONStorage(() => localStorage)
        }
    ))
}

