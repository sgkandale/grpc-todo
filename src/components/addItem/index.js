import React from 'react'
import './style.css'

export default function AddItem(props) {

    const getModalDisplay = () => {
        if (props.showAddItemModal) {
            return 'add_item_modal_display_block'
        } else {
            return 'add_item_modal_display_none'
        }
    }

    const handleSubmit = (event) => {
        event.preventDefault()
    }

    return <div
        className={'add_item_modal ' + getModalDisplay()}
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
                    <input type="text" name="title" placeholder="Title" />
                    <input type="text" name="description" placeholder="Description" />
                    <button
                        className='add_item_button'
                        onClick={handleSubmit}
                    >
                        Add Item
                    </button>
                </form>
            </div>
        </div>
    </div>
}