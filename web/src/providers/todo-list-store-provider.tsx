'use client'

import {createContext, type ReactNode, useContext, useRef} from 'react'
import {useStore} from 'zustand'

import {createToDoListStore, initToDoListStore, type ToDoListStore} from '@/src/stores/todo-list-store'

export type ToDoListStoreApi = ReturnType<typeof createToDoListStore>

export const ToDoListStoreContext = createContext<ToDoListStoreApi | undefined>(
    undefined,
)

export interface ToDoListStoreProviderProps {
    children: ReactNode
}

export const ToDoListStoreProvider = ({
                                          children,
                                      }: ToDoListStoreProviderProps) => {
    const storeRef = useRef<ToDoListStoreApi>()
    if (!storeRef.current) {
        storeRef.current = createToDoListStore(initToDoListStore())
    }

    return (
        <ToDoListStoreContext.Provider value={storeRef.current}>
            {children}
        </ToDoListStoreContext.Provider>
    )
}

export const useToDoListStore = <T, >(
    selector: (store: ToDoListStore) => T,
): T => {
    const toDoListStoreContext = useContext(ToDoListStoreContext)

    if (!toDoListStoreContext) {
        throw new Error(`useToDoListStore must be used within ToDoListStoreProvider`)
    }

    return useStore(toDoListStoreContext, selector)
}