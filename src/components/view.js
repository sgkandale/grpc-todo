import React, { useState, useEffect } from 'react';
import Item from './item';
import Header from './header';
import './style.css';
import AddButton from './addButton';
import AddItem from './addItem';
import { TodoServiceClient } from '../proto/todo_grpc_web_pb';
import { GetItemsRequest } from '../proto/todo_pb';
import { TodoServiceAddress } from '../proto/server';

export default function View() {

    const [showAddItemModal, setShowAddItemModal] = useState(false);
    const [state, setState] = useState('loading')
    const [error, setError] = useState('')
    const [itemsList, setItemsList] = useState([])

    const loadItems = () => {
        const client = new TodoServiceClient(TodoServiceAddress, null, null)
        const req = new GetItemsRequest();

        client.getItems(req, {}, (err, response) => {
            if (err) {
                setError(err.message)
                setState('error')
            } else {
                const res = response.toObject()
                setState('done')
                setItemsList(res.itemsList)
            }
        })
    }

    useEffect(() => {
        loadItems()
    }, [])

    const renderItems = () => {
        if (state === 'loading') {
            return <p>Loading...</p>
        } else if (state === 'error') {
            return <p>Error : {error}</p>
        } else {
            return itemsList.map((item, index) => {
                return <Item key={index} item={item} />
            })
        }
    }

    return <div className='page_view'>
        <Header />
        <main className='main_view'>
            <div className='items_container'>
                {renderItems()}
            </div>
        </main>
        <AddItem
            showAddItemModal={showAddItemModal}
            setShowAddItemModal={setShowAddItemModal}
        />
        <AddButton
            setShowAddItemModal={setShowAddItemModal}
        />
    </div>
}