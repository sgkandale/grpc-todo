import React from 'react';
import Item from './item';
import Header from './header';
import './style.css';
import AddButton from './addButton';

export default function View() {

    return <div className='page_view'>
        <Header />
        <main className='main_view'>
            <div className='items_container'>
                <Item />
                <Item />
            </div>
        </main>
        <AddButton />
    </div>
}