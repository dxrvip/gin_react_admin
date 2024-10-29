import {
    Create,
    SimpleForm,
    TextInput,
    NumberInput,
    BooleanInput,
} from 'react-admin'

function CreateRole() {
    return ( 
        <Create redirect="list" title="添加角色">
            <SimpleForm>
                <TextInput source='name' label="名称" />
                <TextInput source="key" label="标识符" />
                <NumberInput source="sort" label="排序" defaultValue={0} />
                <BooleanInput source="active" label="是否启用" defaultValue={true} />
            </SimpleForm>

        </Create>
     );
}

export default CreateRole;