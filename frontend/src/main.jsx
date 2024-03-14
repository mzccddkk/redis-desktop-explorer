import React from 'react'
import {createRoot} from 'react-dom/client'
import './style.css'
import App from './App'
import NiceModal from '@ebay/nice-modal-react';

const container = document.getElementById('root')

const root = createRoot(container)

root.render(
    <React.StrictMode>
        <NiceModal.Provider>
            <App/>
        </NiceModal.Provider>
    </React.StrictMode>
)
