import {Button, Form, Input, Modal} from 'antd';
import NiceModal, {useModal} from '@ebay/nice-modal-react';
import {CreateConnection} from "../../../wailsjs/go/service/ConnectionService";
import {biz} from "../../../wailsjs/go/models";

export default NiceModal.create(({name}) => {
    // Use a hook to manage the modal state
    const modal = useModal();

    const [form] = Form.useForm();

    const handleSave = () => {
        const connection = new biz.Connection();
        connection.name = form.getFieldValue("connectionName");
        connection.host = form.getFieldValue("host");
        connection.port = form.getFieldValue("port");
        connection.username = form.getFieldValue("username");
        connection.password = form.getFieldValue("password");
        CreateConnection(connection).then(modal.hide)
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
                form={form}
                labelCol={{span: 8}}
                wrapperCol={{span: 16}}
            >

                <Form.Item
                    label="Connection Name"
                    name="connectionName"
                >
                    <Input placeholder="Connection Name"/>
                </Form.Item>

                <Form.Item
                    label="Host"
                    name="host"
                    rules={[{required: false}]}
                    initialValue="127.0.0.1"
                >
                    <Input/>
                </Form.Item>

                <Form.Item
                    label="Port"
                    name="port"
                    rules={[{required: false}]}
                    initialValue="6379"
                >
                    <Input/>
                </Form.Item>

                <Form.Item
                    label="Username"
                    name="username"
                    rules={[{required: false}]}
                >
                    <Input placeholder="(Optional) Redis Auth Username"/>
                </Form.Item>

                <Form.Item
                    label="Password"
                    name="password"
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
