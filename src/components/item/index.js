import React from 'react';
import './style.css';

export default function Item() {

    return <div className='item_container'>
        <div className='item_title_container'>
            <span className='item_tick'>
                <input
                    type="checkbox"
                    name="item-name"
                    checked={true}
                    className='item_tick_input'
                />
            </span>
            <h4 className='item_title'>
                Hello
            </h4>
            <button className='item_delete_button'>
                X
            </button>
        </div>
        <div className='item_description_container'>
            Desc
        </div>
    </div>
}