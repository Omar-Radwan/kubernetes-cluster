import { REACT_APP_API_URL, REACT_APP_API_PORT } from "./constants"
export function fetchCount(path, setMessage, increment) {
    console.log(process.env.NODE_ENV)
    fetch(`http://${REACT_APP_API_URL}:${REACT_APP_API_PORT}/${path}?${increment}`, {
        method: 'GET',
    }).then((res) => res.json())
        .then((json) => {
            setMessage(json.Value)
        })
}