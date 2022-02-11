import react from "react";
import "./style.css";

export default function AddButton(props) {

    return <div className='add_button_container'>
        <button
            className='add_button'
            onClick={() => props.setShowAddItemModal(true)}
        >
            +
        </button>
    </div>
}