import { useState, useEffect } from "react";
import { fetchCount } from "../utils/utils"

function ButtonWithLabel(props) {
    const [message, setMessage] = useState("-1")
    useEffect(() => {
        fetchCount(props.path, setMessage, "")
    }, []);
    var resource = (props.path === "" ? "root" : props.path)
    resource = resource.charAt(0).toUpperCase() + resource.slice(1)
    return (
        <div >
            <button onClick={() => {
                fetchCount(props.path, setMessage, "increment")
            }} >{resource} Press Me</button>
            <label> {message}</label>
        </div>
    );
}

export default ButtonWithLabel;
