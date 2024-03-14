import {Button, Form, Input, Modal} from 'antd';
import NiceModal, {useModal} from '@ebay/nice-modal-react';

export default NiceModal.create(({name}) => {
    // Use a hook to manage the modal state
    const modal = useModal();

    const handleSave = () => {
        modal.hide()
    };

    const handleTest = () => {
        modal.hide()
    };

    return (
        <Modal
            title="New Connection"
            open={modal.visible}
            afterClose={() => modal.remove()}
            footer={null}
            keyboard={false}
            mask={false}
            maskClosable={false}
            closable={false}
        >

            <Form
                labelCol={{span: 8}}
                // wrapperCol={{span: 16}}
            >

                <Form.Item
                    label="Connection Name"
                    rules={[{required: false}]}
                >
                    <Input placeholder="Connection Name"/>
                </Form.Item>

                <Form.Item
                    label="Host"
                    rules={[{required: false}]}
                >
                    <Input defaultValue="127.0.0.1"/>
                </Form.Item>

                <Form.Item
                    label="Port"
                    rules={[{required: false}]}
                >
                    <Input defaultValue="6379"/>
                </Form.Item>

                <Form.Item
                    label="Username"
                    rules={[{required: false}]}
                >
                    <Input placeholder="(Optional) Redis Auth Username"/>
                </Form.Item>

                <Form.Item
                    label="Password"
                    rules={[{required: false}]}
                >
                    <Input.Password placeholder="(Optional) Redis Auth Password"/>
                </Form.Item>

                <div style={{display: 'flex', justifyContent: 'space-between'}}>
                    <Button onClick={() => handleTest()}>Test Connection</Button>
                    <div>
                        <Button onClick={() => modal.hide()}>Cancel</Button>
                        <Button type="primary" style={{marginLeft: 16}} onClick={() => handleSave()}>Save</Button>
                    </div>
                </div>
            </Form>

        </Modal>
    );
});
