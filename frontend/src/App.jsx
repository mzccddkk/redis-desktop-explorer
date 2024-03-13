import {useEffect, useState} from 'react';
import './App.css';
import {ListConnection} from "../wailsjs/go/service/ConnectionService.js";
import Connection from "./components/Connection.jsx";

function App() {
    // const [resultText, setResultText] = useState("Please enter your name below 👇");
    // const [name, setName] = useState('');
    // const updateName = (e) => setName(e.target.value);
    // const updateResultText = (result) => setResultText(result);
    //
    // function greet() {
    //     Greet(name).then(updateResultText);
    // }

    // Connection Data
    const [connectionData, setConnectionData] = useState(null);

    // Fetch Connection Data
    useEffect(() => {
        const fetchConnectionData = async () => {
            await ListConnection().then(setConnectionData);
        }
        fetchConnectionData().then(r => {
        })
    }, [])

    return (
        <div id="App">
            <div id="Menu">
                <Connection></Connection>
            </div>

            <div id="Navigation">
                {
                    connectionData ? (
                        <ul>
                            {
                                connectionData.map(item => (
                                    <li key={item.name}>{item.name}</li>
                                ))
                            }
                        </ul>
                    ) : (
                        <ul>
                        </ul>
                    )
                }
            </div>

            <div id="Information">

            </div>
        </div>
    )
}

export default App
