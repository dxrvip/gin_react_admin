import { Edit, maxLength, minLength, required, SimpleForm, TextInput, useRecordContext } from "react-admin";

function BrandEdit() {
    
    // const record = useRecordContext();
    // console.log(record)
    return (
        <Edit title={`编辑品牌`} redirect="list">
            <SimpleForm>
                <TextInput source="name" label="品牌名称" validate={[required(), minLength(3), maxLength(255)]} />
                <TextInput source="logo" label="Logo URL" />
                <TextInput source="description" label="品牌描述" />
            </SimpleForm>
        </Edit>
    );
}

export default BrandEdit;