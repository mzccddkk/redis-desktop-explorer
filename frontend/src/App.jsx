import {useEffect, useState} from 'react';
import './App.css';
import {ListConnection} from "../wailsjs/go/service/ConnectionService";
import NiceModal from "@ebay/nice-modal-react";
import MyAntdModal from './components/modal/Connection';

function App() {
    // const [resultText, setResultText] = useState("Please enter your name below 👇");
    // const [name, setName] = useState('');
    // const updateName = (e) => setName(e.target.value);
    // const updateResultText = (result) => setResultText(result);
    //
    // function greet() {
    //     Greet(name).then(updateResultText);
    // }

    const showAntdModal = () => {
        // Show a modal with arguments passed to the component as props
        NiceModal.show(MyAntdModal, { name: 'Nate' })
    };

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
                <button onClick={showAntdModal}>Connect to Redis</button>
                {/*<Connection></Connection>*/}
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
