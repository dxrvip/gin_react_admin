import { Edit, maxLength, minLength, required, SimpleForm, TextInput } from "react-admin";

function BrandEdit() {
    return (
        <Edit>
            <SimpleForm>
                <TextInput source="name" label="品牌名称" validate={[required(), minLength(3), maxLength(255)]} />
                <TextInput source="logo" label="Logo URL" />
                <TextInput source="description" label="品牌描述" />
            </SimpleForm>
        </Edit>
    );
}

export default BrandEdit;