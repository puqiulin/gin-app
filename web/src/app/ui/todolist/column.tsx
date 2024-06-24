import {TaskC} from "@/src/app/ui/todolist/task";
import {useState} from "react";
import classNames from "classnames";
import {useToDoListStore} from "@/src/providers/todo-list-store-provider";

export function Column({state}: { state: string }) {
    const tasks = useToDoListStore(store => store.tasks.filter(task => task.state === state))
    const addTask = useToDoListStore(store => store.addTask)
    const draggedTask = useToDoListStore(store => store.draggedTask)
    const setDraggedTask = useToDoListStore(store => store.setDraggedTask)
    const moveTask = useToDoListStore(store => store.moveTask)

    const [title, setTitle] = useState<string>("")
    const [open, setOpen] = useState<boolean>(false)
    const [onDrag, setOnDrag] = useState<boolean>(false)

    const columnClass = classNames({
        "bg-base-300 rounded-lg p-4 border": true,
        //for same border after add border-2
        "border-solid border-5 transparent": true,
        "border-dashed border-sky-500": onDrag,
    })

    const modalClass = classNames({
        "modal": true,
        "modal-open": open
    })

    return <div
        className={columnClass}
        onDragOver={e => {
            setOnDrag(true)
            e.preventDefault()
        }}
        onDragLeave={e => {
            setOnDrag(false)
        }}
        onDrop={e => {
            setOnDrag(false)
            moveTask(draggedTask, state)
            setDraggedTask(null)
        }}
    >
        <div className="flex items-center justify-between mb-3">
            <div className='flex items-center gap-2'>
                <div className="font-bold text-3xl text-primary">{state}</div>
                <div className='text-2xl font-bold text-gray-500'>{tasks.length}</div>
            </div>
            <button
                className="btn btn-primary"
                onClick={() => {
                    setOpen(true)
                }}>
                +
            </button>
        </div>
        <div className="flex flex-col gap-2 ">
            {tasks.map(task => (
                <TaskC key={task.title} task={task}/>
            ))}
        </div>

        <dialog id="add_task" className={modalClass}>
            <div className="modal-box flex flex-col gap-2">
                <div>Add <span className='font-bold'>{state} </span>task</div>
                <div className="flex justify-between">
                    <input type="text" placeholder="title" className="input w-full max-w-xs" onChange={e => {
                        setTitle(e.target.value)
                    }}/>
                    <button className="btn" onClick={() => {
                        addTask({title, state})
                        setOpen(false)
                    }}>👌
                    </button>
                </div>
            </div>

            <form method="dialog" className="modal-backdrop">
                <button>close</button>
            </form>
        </dialog>
    </div>
}