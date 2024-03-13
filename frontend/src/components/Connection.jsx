import React, {useState} from 'react';
import '../styles/connection.css'

const LoginForm = ({onClose}) => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        console.log('Username:', username);
        console.log('Password:', password);
        onClose();
    };

    return (
        <div className="modal-content">
            <form>
                <div>
                    <label>Username:</label>
                    <input type="text" value={username} onChange={(e) => setUsername(e.target.value)}/>
                </div>
                <div>
                    <label>Password:</label>
                    <input type="password" value={password} onChange={(e) => setPassword(e.target.value)}/>
                </div>
                <button>Test</button>
                <button onClick={onClose}>Cancel</button>
                <button onClick={handleSubmit}>Save</button>
            </form>
        </div>
    );
};

const Connection = () => {
    const [showModal, setShowModal] = useState(false);

    const handleOpenModal = () => {
        setShowModal(true);
    };

    const handleCloseModal = () => {
        setShowModal(false);
    };

    return (
        <div>
            <button onClick={handleOpenModal}>Connect to Redis</button>
            {showModal && <div className="modal">
                <div className="modal-background" onClick={handleCloseModal}></div>
                <LoginForm onClose={handleCloseModal}/>
            </div>}
        </div>
    );
};

export default Connection;
