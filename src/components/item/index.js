import React from 'react';
import './style.css';
import { TodoServiceClient } from '../../proto/todo_grpc_web_pb';
import { Item as ItemRequest } from '../../proto/todo_pb';
import { TodoServiceAddress } from '../../proto/server';

export default function Item(props) {
    const item = props.item;

    const closeItem = () => {
        item.closed = true
        const client = new TodoServiceClient(TodoServiceAddress, null, null)
        const req = new ItemRequest();
        req.setId(item.id)

        client.closeItem(req, {}, (err, response) => {
            if (err) {
                console.log(err)
            }
        })
    }

    const deleteItem = () => {
        const client = new TodoServiceClient(TodoServiceAddress, null, null)
        const req = new ItemRequest();
        req.setId(item.id)

        client.deleteItem(req, {}, (err, response) => {
            if (err) {
                console.log(err)
            }
        })
    }

    return <div className='item_container'>
        <div className='item_title_container'>
            <span className='item_tick'>
                <input
                    type="checkbox"
                    name="item-name"
                    checked={item.closed}
                    className='item_tick_input'
                    onChange={closeItem}
                />
            </span>
            <h4 className='item_title'>
                {item.title}
            </h4>
            <button
                className='item_delete_button'
                onClick={deleteItem}
            >
                X
            </button>
        </div>
        <div className='item_description_container'>
            {item.description}
        </div>
    </div>
}