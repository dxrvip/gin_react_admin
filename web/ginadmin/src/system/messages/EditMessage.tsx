import {
    Edit,
    SimpleForm,
    TextInput,
} from "react-admin"
import { RichTextInput } from 'ra-input-rich-text';

function EditMessage() {
    return (
        <Edit>
            <SimpleForm>
                <TextInput source="id" label="id" disabled />

                <TextInput source="title" label="标题" />
                <RichTextInput source="content" label="消息内容" />


            </SimpleForm>
        </Edit>
    );
}

export default EditMessage;