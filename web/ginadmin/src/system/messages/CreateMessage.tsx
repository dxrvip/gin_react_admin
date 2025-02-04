import {
    Create,
    SimpleForm,
    TextInput,
} from 'react-admin';
import { RichTextInput } from 'ra-input-rich-text';

function CreateMessage() {
    return (
        <Create title="Create a Message">
            <SimpleForm>
                <TextInput source="title" label="标题" />
                <RichTextInput source="content" label="消息内容" style={{height: '500px'}} />
                  
            </SimpleForm>

        </Create>
    );
}

export default CreateMessage;