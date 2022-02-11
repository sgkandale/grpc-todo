import React, { useState } from 'react';
import Item from './item';
import Header from './header';
import './style.css';
import AddButton from './addButton';
import AddItem from './addItem';

export default function View() {

    const [showAddItemModal, setShowAddItemModal] = useState(false);

    return <div className='page_view'>
        <Header />
        <main className='main_view'>
            <div className='items_container'>
                <Item />
                <Item />
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