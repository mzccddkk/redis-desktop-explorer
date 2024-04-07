import {useEffect, useState} from 'react';
import './App.css';
import {ListConnection} from "../wailsjs/go/service/ConnectionService";
import NiceModal from "@ebay/nice-modal-react";
import MyAntdModal from './components/modal/Connection';
import {DatabaseFilled, PlusCircleFilled, SettingFilled} from '@ant-design/icons';

function App() {
    const showAntdModal = () => {
        // Show a modal with arguments passed to the component as props
        NiceModal.show(MyAntdModal, {name: 'Nate'})
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
        <div id="app">
            <div className="side-bar">
                <div className="side-top" title="Connect">
                    <div className="text-letter " title="Connect">
                        <DatabaseFilled/>
                    </div>
                </div>
                <div className="side-bottom">
                    <div className="text-letter " title="Add">
                        <PlusCircleFilled/>
                    </div>
                    <div className="text-letter " title="Setting">
                        <SettingFilled/>
                    </div>
                </div>
            </div>

            <div className="content">
                <div className="ccc">
                    ccc
                </div>
                <div className="ddd">
                    ddd
                </div>
            </div>

            {/*<div id="Menu">*/}
            {/*    <button onClick={showAntdModal}>Connect to Redis</button>*/}
            {/*    /!*<Connection></Connection>*!/*/}
            {/*</div>*/}

            {/*<div id="Navigation">*/}
            {/*    {*/}
            {/*        connectionData ? (*/}
            {/*            <ul>*/}
            {/*                {*/}
            {/*                    connectionData.map(item => (*/}
            {/*                        <li key={item.name}>{item.name}</li>*/}
            {/*                    ))*/}
            {/*                }*/}
            {/*            </ul>*/}
            {/*        ) : (*/}
            {/*            <ul>*/}
            {/*            </ul>*/}
            {/*        )*/}
            {/*    }*/}
            {/*</div>*/}

            {/*<div id="Information">*/}

            {/*</div>*/}
        </div>
    )
}

export default App
