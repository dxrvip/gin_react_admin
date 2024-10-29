import {
    BooleanInput,
    Edit,
    NumberInput,
    SimpleForm,
    TextInput
}from "react-admin"

function EditRole() {
    return ( 
        <Edit title="编辑角色" redirect="list">
             <SimpleForm>
                <TextInput source="id" label="Id" disabled />
                <TextInput source='name' label="名称" />
                <TextInput source="key" label="标识符" />
                <NumberInput source="sort" label="排序" defaultValue={0} />
                <BooleanInput source="active" label="是否启用" defaultValue={true} />
            </SimpleForm>
        </Edit>
     );
}

export default EditRole;