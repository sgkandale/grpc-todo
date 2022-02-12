import React, { useState } from 'react'
import './style.css'
import { TodoServiceClient } from '../../proto/todo_grpc_web_pb';
import { Item } from '../../proto/todo_pb';
import { TodoServiceAddress } from '../../proto/server';

const initValues = {
    title: "",
    description: "",
    error: "",
}

export default function AddItem(props) {
    const [values, setValues] = useState(initValues)
    const [state, setState] = useState('form')
    const changeValues = (event) => {
        const { name, value } = event.target
        setValues({ ...values, [name]: value })
    }

    const getModalDisplay = () => {
        if (props.showAddItemModal) {
            return 'visible'
        } else {
            return 'hidden'
        }
    }

    const handleSubmit = (event) => {
        event.preventDefault()
        setState('loading')
        const client = new TodoServiceClient(TodoServiceAddress, null, null)
        const req = new Item();
        req.setTitle(values.title)
        req.setDescription(values.description)

        client.createItem(req, {}, (err, response) => {
            if (err) {
                setValues({ ...values, error: err.message })
                setState('error')
            } else {
                setState('done')
                setValues(initValues)
                setTimeout(() => {
                    setState('form')
                    props.setShowAddItemModal(false)
                    props.loadItems()
                }, 1000)
            }
        })
    }

    return <div
        className={`add_item_modal ${getModalDisplay()}`}
        onCli
    >
        <div className='add_item_modal_content'>
            <div className='add_item_modal_header'>
                <h3 className='add_item_header'>
                    Add Item
                </h3>
                <button
                    className='close_add_item_modal'
                    onClick={() => props.setShowAddItemModal(false)}
                >
                    X
                </button>
            </div>
            <div className='add_item_modal_body'>
                <form
                    className='add_item_form'
                    onSubmit={handleSubmit}
                >
                    <input
                        type="text"
                        name="title"
                        placeholder="Title"
                        value={values.title}
                        onChange={changeValues}
                    />
                    <input
                        type="text"
                        name="description"
                        placeholder="Description"
                        value={values.description}
                        onChange={changeValues}
                    />
                    <button
                        className='add_item_button'
                        onClick={handleSubmit}
                        disabled={state === 'loading'}
                    >
                        {
                            state === 'done' ? 'Done' :
                                state === 'loading' ? 'Loading...' :
                                    'Add Item'
                        }
                    </button>
                </form>
                <p>
                    {state === 'error' ? values.error : ''}
                </p>
            </div>
        </div>
    </div>
}